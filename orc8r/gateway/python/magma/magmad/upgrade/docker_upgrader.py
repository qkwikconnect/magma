"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import asyncio
import json
import logging
import pathlib
import subprocess

from magma.common.service import MagmaService
from magma.configuration.service_configs import load_service_config
from magma.magmad.upgrade.magma_upgrader import compare_package_versions
from magma.magmad.upgrade.upgrader import UpgraderFactory
from magma.magmad.upgrade.upgrader2 import ImageNameT, run_command, \
    UpgradeIntent, Upgrader2, VersionInfo, VersionT


MAGMA_GITHUB_PATH = "/tmp/magma_upgrade"
MAGMA_GITHUB_URL = "https://github.com/facebookincubator/magma.git"


class DockerUpgrader(Upgrader2):
    """
    Downloads and installs images
    """

    def perform_upgrade_if_necessary(self, target_version: str) -> None:
        """
        Target version comes from tier configuration currently.
        """

        if self.upgrade_task and not self.upgrade_task.done():
            logging.info("Not starting another upgrade, upgrade in progress")
            return
        self.upgrade_task = self.loop.create_task(self.do_docker_upgrade())

    def version_to_image_name(self, version: VersionT) -> ImageNameT:
        """
        Returns the image tag from the version string.
        (i.e) 0.3.68-1541626353-d1c29db1 -> d1c29db1
        """
        parts = version.split("-")
        if len(parts) != 3:
            raise ValueError("Unknown version format: %s" % version)
        return ImageNameT("%s" % parts[2])

    async def get_upgrade_intent(self) -> UpgradeIntent:
        """
        Returns the desired version for the gateway.
        We don't support downgrading, and so checks are made to update
        only if the target version is higher than the current version.
        """
        tgt_version = self.service.mconfig.package_version
        curr_version = self.service.version
        if (tgt_version == "0.0.0-0"
                or compare_package_versions(curr_version, tgt_version) <= 0):
            tgt_version = curr_version
        return UpgradeIntent(stable=VersionT(tgt_version), canary=VersionT(""))

    async def get_versions(self) -> VersionInfo:
        """ Returns the current version by parsing the currently running
        docker image tag
        """
        magmad_stdout = subprocess.check_output(["docker", "inspect", "magmad"])
        magmad_stdoutstr = str(magmad_stdout, 'utf-8').strip()[1:-1]
        magmad_inspect_json = json.loads(magmad_stdoutstr)
        current_version = str.split(magmad_inspect_json["Config"]["Image"], ":")[-1]
        return VersionInfo(
            current_version=current_version,
            available_versions=set(),
        )

    async def prepare_upgrade(
        self, version: VersionT, path_to_image: pathlib.Path
    ) -> None:
        """Install the new docker-compose file"""

        gw_module = self.service.config["upgrader_factory"].get("gateway_module")

        # Copy updated docker-compose
        await run_command("cp {}/magma/{}/gateway/docker/docker-compose.yml "
                          "/var/opt/magma/docker".format(MAGMA_GITHUB_PATH,
                                                         gw_module),
                          shell=True, check=True)

    async def upgrade(self, version: VersionT, path_to_image: pathlib.Path) -> None:
        """Upgrade is a no-op as an external process (e.g. cron) must
        trigger it
        """

    async def do_docker_upgrade(self) -> None:
        """
           Do a single loop of the upgrade process, using the given upgrader.
        """
        try:
            await self._do_docker_upgrade()
        except Exception as exp:  # pylint: disable=broad-except
            logging.error("Upgrade loop failed! Error: %s", exp)

    async def _do_docker_upgrade(self) -> None:
        upgrade_intent, version_info = await asyncio.gather(
            self.get_upgrade_intent(), self.get_versions()
        )

        current_version = version_info.current_version or object()
        to_upgrade_to = upgrade_intent.version_to_force_upgrade(version_info)

        if to_upgrade_to:
            logging.info(
                "There is work to be done:\n"
                "  stable: %s\n"
                "  current: %s\n"
                "  to_upgrade: %s",
                upgrade_intent.stable,
                current_version,
                to_upgrade_to,
            )

            assert to_upgrade_to != version_info.current_version
            logging.warning(
                "Version %r is out of date! Upgrading to %r",
                version_info.current_version,
                to_upgrade_to,
            )
            image_name = self.version_to_image_name(to_upgrade_to)
            await download_update(current_version, image_name)
            await self.prepare_upgrade(
                current_version, pathlib.Path(MAGMA_GITHUB_PATH, "magma"))


async def download_update(
    old_version: ImageNameT,
    new_version: ImageNameT,
) -> None:
    """
    Download the images for the given tag and clones the github repo.
    """
    await run_command("rm -rf {}".format(MAGMA_GITHUB_PATH), shell=True,
                      check=True)
    await run_command("mkdir -p {}".format(MAGMA_GITHUB_PATH), shell=True,
                      check=True)

    control_proxy_config = load_service_config('control_proxy')
    await run_command("cp {} /usr/local/share/ca-certificates/rootCA.crt".
                      format(control_proxy_config['rootca_cert']), shell=True,
                      check=True)
    await run_command("update-ca-certificates", shell=True, check=True)

    git_clone_cmd = "git -c http.proxy=https://{}:{} -C {} clone {}".format(
        control_proxy_config['bootstrap_address'],
        control_proxy_config['bootstrap_port'], MAGMA_GITHUB_PATH,
        MAGMA_GITHUB_URL)
    await run_command(git_clone_cmd, shell=True, check=True)
    git_checkout_cmd = "git -C {}/magma checkout {}".format(MAGMA_GITHUB_PATH,
                                                      new_version)
    await run_command(git_checkout_cmd, shell=True, check=True)
    docker_login_cmd = "docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD " \
                       "$DOCKER_REGISTRY"
    await run_command(docker_login_cmd, shell=True, check=True)
    docker_pull_cmd = "IMAGE_VERSION={} docker-compose --project-directory " \
                      "/var/opt/magma/docker -f " \
                      "/var/opt/magma/docker/docker-compose.yml pull -q".\
        format(new_version)
    await run_command(docker_pull_cmd, shell=True, check=True)

    # Update the image tag in the .env
    sed_args = "sed -i s/IMAGE_VERSION={}/IMAGE_VERSION={}/g " \
               "var/opt/magma/docker/.env".format(old_version, new_version)
    await run_command(sed_args, shell=True, check=True)


class DockerUpgraderFactory(UpgraderFactory):
    """ Returns an instance of the DockerUpgrader """

    def create_upgrader(
        self,
        magmad_service: MagmaService,
        loop: asyncio.AbstractEventLoop,
    ) -> DockerUpgrader:
        return DockerUpgrader(magmad_service)