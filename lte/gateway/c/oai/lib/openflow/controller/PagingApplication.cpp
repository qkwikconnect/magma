/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the Apache License, Version 2.0  (the "License"); you may not use this file
 * except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#include <netinet/ip.h>
#include <arpa/inet.h>
#include "OpenflowController.h"
#include "PagingApplication.h"
#include "RpcClient.h"

extern "C" {
#include "log.h"
#include "sgw_paging.h"
}

using namespace fluid_msg;

namespace openflow {

uint32_t prefix2mask(int prefix)
{
  if (prefix) {
    return htonl(~((1 << (32 - prefix)) - 1));
  } else {
    return htonl(0);
  }
}

void PagingApplication::event_callback(
  const ControllerEvent &ev,
  const OpenflowMessenger &messenger)
{
  if (ev.get_type() == EVENT_PACKET_IN) {
    OAILOG_DEBUG(LOG_GTPV1U, "Handling packet-in message in paging app\n");
    const PacketInEvent &pi = static_cast<const PacketInEvent &>(ev);
    of13::PacketIn ofpi;
    ofpi.unpack(const_cast<uint8_t *>(pi.get_data()));
    handle_paging_message(
      ev.get_connection(), static_cast<uint8_t *>(ofpi.data()), messenger);
  } else if (ev.get_type() == EVENT_SWITCH_UP) {
    install_default_flow(ev.get_connection(), messenger);
  }
}

void PagingApplication::handle_paging_message(
  fluid_base::OFConnection *ofconn,
  uint8_t *data,
  const OpenflowMessenger &messenger)
{
  // send paging request to MME
  struct ip *ip_header = (struct ip *) (data + ETH_HEADER_LENGTH);
  struct in_addr dest_ip;
  memcpy(&dest_ip, &ip_header->ip_dst, sizeof(struct in_addr));
  char *dest_ip_str = inet_ntoa(dest_ip);
  OAILOG_DEBUG(
    LOG_GTPV1U, "Initiating paging procedure for IP %s\n", dest_ip_str);
  sgw_send_paging_request(&dest_ip);

  /*
   * Clamp on this ip for configured amount of time
   * Priority is above default paging flow, but below gtp flow. This way when
   * paging succeeds, this flow will be ignored.
   * The clamping time is necessary to prevent packets from continually hitting
   * userspace, and as a retry time if paging fails
   */
  of13::FlowMod fm =
    messenger.create_default_flow_mod(0, of13::OFPFC_ADD, MID_PRIORITY + 1);
  fm.hard_timeout(CLAMPING_TIMEOUT);
  of13::EthType type_match(IP_ETH_TYPE);
  fm.add_oxm_field(type_match);

  of13::IPv4Dst ip_match(dest_ip.s_addr);
  fm.add_oxm_field(ip_match);

  // No actions mean packet is dropped
  messenger.send_of_msg(fm, ofconn);
  return;
}

void PagingApplication::install_default_flow(
  fluid_base::OFConnection *ofconn,
  const OpenflowMessenger &messenger)
{
  // Get assigned IP block from mobilityd
  struct in_addr netaddr;
  uint32_t prefix;
  int ret = get_assigned_ipv4_block(0, &netaddr, &prefix);

  // Convert to string for logging
  char ip_str[INET_ADDRSTRLEN];
  inet_ntop(AF_INET, &(netaddr.s_addr), ip_str, INET_ADDRSTRLEN);
  OAILOG_INFO(
    LOG_GTPV1U,
    "Setting default paging flow for UE IP block %s/%d\n",
    ip_str,
    prefix);

  of13::FlowMod fm =
    messenger.create_default_flow_mod(0, of13::OFPFC_ADD, MID_PRIORITY);
  // IP eth type
  of13::EthType type_match(IP_ETH_TYPE);
  fm.add_oxm_field(type_match);

  // Match on ip dest equalling assigned ue ip block
  of13::IPv4Dst ip_match(netaddr.s_addr, prefix2mask(prefix));
  fm.add_oxm_field(ip_match);

  // Output to controller
  of13::OutputAction act(of13::OFPP_CONTROLLER, of13::OFPCML_NO_BUFFER);
  of13::ApplyActions inst;
  inst.add_action(act);
  fm.add_instruction(inst);

  messenger.send_of_msg(fm, ofconn);
  OAILOG_DEBUG(LOG_GTPV1U, "Default paging flow added\n");
}

} // namespace openflow
