// Copyright 2019 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openflow

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"antrea.io/libOpenflow/openflow15"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/tools/cache"

	"antrea.io/antrea/v2/pkg/agent/config"
	"antrea.io/antrea/v2/pkg/agent/openflow/cookie"
	"antrea.io/antrea/v2/pkg/agent/types"
	"antrea.io/antrea/v2/pkg/apis/controlplane/v1beta2"
	binding "antrea.io/antrea/v2/pkg/ovs/openflow"
)

// errSkipNetworkPolicyMetricFlow means the OpenFlow dump line is not a valid
// NetworkPolicy metric flow (for example a stale routing flow after a pipeline table ID shift).
var errSkipNetworkPolicyMetricFlow = errors.New("not a valid network policy metric flow")

// errSkipMulticastMetricFlow means the OpenFlow dump line is not a valid multicast metric flow.
var errSkipMulticastMetricFlow = errors.New("not a valid multicast metric flow")

var (
	MatchDstIP          = types.NewMatchKey(binding.ProtocolIP, types.IPAddr, "nw_dst")
	MatchSrcIP          = types.NewMatchKey(binding.ProtocolIP, types.IPAddr, "nw_src")
	MatchDstIPNet       = types.NewMatchKey(binding.ProtocolIP, types.IPNetAddr, "nw_dst")
	MatchSrcIPNet       = types.NewMatchKey(binding.ProtocolIP, types.IPNetAddr, "nw_src")
	MatchCTDstIP        = types.NewMatchKey(binding.ProtocolIP, types.IPAddr, "ct_nw_dst")
	MatchCTSrcIP        = types.NewMatchKey(binding.ProtocolIP, types.IPAddr, "ct_nw_src")
	MatchCTDstIPNet     = types.NewMatchKey(binding.ProtocolIP, types.IPNetAddr, "ct_nw_dst")
	MatchCTSrcIPNet     = types.NewMatchKey(binding.ProtocolIP, types.IPNetAddr, "ct_nw_src")
	MatchDstIPv6        = types.NewMatchKey(binding.ProtocolIPv6, types.IPAddr, "ipv6_dst")
	MatchSrcIPv6        = types.NewMatchKey(binding.ProtocolIPv6, types.IPAddr, "ipv6_src")
	MatchDstIPNetv6     = types.NewMatchKey(binding.ProtocolIPv6, types.IPNetAddr, "ipv6_dst")
	MatchSrcIPNetv6     = types.NewMatchKey(binding.ProtocolIPv6, types.IPNetAddr, "ipv6_src")
	MatchCTDstIPv6      = types.NewMatchKey(binding.ProtocolIPv6, types.IPAddr, "ct_ipv6_dst")
	MatchCTSrcIPv6      = types.NewMatchKey(binding.ProtocolIPv6, types.IPAddr, "ct_ipv6_src")
	MatchCTDstIPNetv6   = types.NewMatchKey(binding.ProtocolIPv6, types.IPNetAddr, "ct_ipv6_dst")
	MatchCTSrcIPNetv6   = types.NewMatchKey(binding.ProtocolIPv6, types.IPNetAddr, "ct_ipv6_src")
	MatchDstOFPort      = types.NewMatchKey(binding.ProtocolIP, types.OFPortAddr, "reg1[0..31]")
	MatchSrcOFPort      = types.NewMatchKey(binding.ProtocolIP, types.OFPortAddr, "in_port")
	MatchTCPDstPort     = types.NewMatchKey(binding.ProtocolTCP, types.L4PortAddr, "tp_dst")
	MatchTCPv6DstPort   = types.NewMatchKey(binding.ProtocolTCPv6, types.L4PortAddr, "tp_dst")
	MatchUDPDstPort     = types.NewMatchKey(binding.ProtocolUDP, types.L4PortAddr, "tp_dst")
	MatchUDPv6DstPort   = types.NewMatchKey(binding.ProtocolUDPv6, types.L4PortAddr, "tp_dst")
	MatchSCTPDstPort    = types.NewMatchKey(binding.ProtocolSCTP, types.L4PortAddr, "tp_dst")
	MatchSCTPv6DstPort  = types.NewMatchKey(binding.ProtocolSCTPv6, types.L4PortAddr, "tp_dst")
	MatchSCTPSrcPort    = types.NewMatchKey(binding.ProtocolSCTP, types.L4PortAddr, "tp_src")
	MatchSCTPv6SrcPort  = types.NewMatchKey(binding.ProtocolSCTPv6, types.L4PortAddr, "tp_src")
	MatchTCPSrcPort     = types.NewMatchKey(binding.ProtocolTCP, types.L4PortAddr, "tp_src")
	MatchTCPv6SrcPort   = types.NewMatchKey(binding.ProtocolTCPv6, types.L4PortAddr, "tp_src")
	MatchUDPSrcPort     = types.NewMatchKey(binding.ProtocolUDP, types.L4PortAddr, "tp_src")
	MatchUDPv6SrcPort   = types.NewMatchKey(binding.ProtocolUDPv6, types.L4PortAddr, "tp_src")
	MatchICMPType       = types.NewMatchKey(binding.ProtocolICMP, types.ICMPAddr, "icmp_type")
	MatchICMPCode       = types.NewMatchKey(binding.ProtocolICMP, types.ICMPAddr, "icmp_code")
	MatchICMPv6Type     = types.NewMatchKey(binding.ProtocolICMPv6, types.ICMPAddr, "icmpv6_type")
	MatchICMPv6Code     = types.NewMatchKey(binding.ProtocolICMPv6, types.ICMPAddr, "icmpv6_code")
	MatchServiceGroupID = types.NewMatchKey(binding.ProtocolIP, types.ServiceGroupIDAddr, "reg7[0..31]")
	MatchIGMPProtocol   = types.NewMatchKey(binding.ProtocolIGMP, types.IGMPAddr, "igmp")
	MatchLabelID        = types.NewMatchKey(binding.ProtocolIP, types.LabelIDAddr, "tun_id")
	MatchTCPFlags       = types.NewMatchKey(binding.ProtocolTCP, types.TCPFlagsAddr, "tcp_flags")
	MatchTCPv6Flags     = types.NewMatchKey(binding.ProtocolTCPv6, types.TCPFlagsAddr, "tcp_flags")
	// MatchCTState should be used with ct_state condition as matchValue.
	// MatchValue example: `+rpl+trk`.
	MatchCTState = types.NewMatchKey(binding.ProtocolIP, types.CTStateAddr, "ct_state")
	Unsupported  = types.NewMatchKey(binding.ProtocolIP, types.UnSupported, "unknown")

	// metricFlowIdentifier is used to identify metric flows in metric table.
	// There could be other flows like default flow and Traceflow flows in the table. Only metric flows are supposed to
	// have normal priority.
	metricFlowIdentifier = fmt.Sprintf("priority=%d,", priorityNormal)

	protocolTCP = v1beta2.ProtocolTCP
	dnsPort     = int32(53)
)

type TCPFlags struct {
	Flag uint16
	Mask uint16
}

// IP address calculated from Pod's address.
type IPAddress net.IP

func (a *IPAddress) GetMatchKey(addrType types.AddressType) *types.MatchKey {
	_ = "STUB: not implemented"
	return nil
}

func (a *IPAddress) GetMatchValue() string { _ = "STUB: not implemented"; return "" }

func (a *IPAddress) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func NewIPAddress(addr net.IP) *IPAddress { _ = "STUB: not implemented"; return nil }

// IP block calculated from Pod's address.
type IPNetAddress net.IPNet

func (a *IPNetAddress) GetMatchKey(addrType types.AddressType) *types.MatchKey {
	_ = "STUB: not implemented"
	return nil
}

func (a *IPNetAddress) GetMatchValue() string { _ = "STUB: not implemented"; return "" }

func (a *IPNetAddress) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func NewIPNetAddress(addr net.IPNet) *IPNetAddress { _ = "STUB: not implemented"; return nil }

// OFPortAddress is the Openflow port of an interface.
type OFPortAddress int32

func (a *OFPortAddress) GetMatchKey(addrType types.AddressType) *types.MatchKey {
	_ = "STUB: not implemented"
	return nil
}

// in_port is used in egress rule to match packets sent from local Pod. Service traffic is not covered by this
// match, and source IP will be matched instead.

func (a *OFPortAddress) GetMatchValue() string { _ = "STUB: not implemented"; return "" }

func (a *OFPortAddress) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func NewOFPortAddress(addr int32) *OFPortAddress { _ = "STUB: not implemented"; return nil }

type ServiceGroupIDAddress binding.GroupIDType

func (a *ServiceGroupIDAddress) GetMatchKey(addrType types.AddressType) *types.MatchKey {
	_ = "STUB: not implemented"
	return nil
}

func (a *ServiceGroupIDAddress) GetMatchValue() string { _ = "STUB: not implemented"; return "" }

func (a *ServiceGroupIDAddress) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func NewServiceGroupIDAddress(groupID binding.GroupIDType) *ServiceGroupIDAddress {
	_ = "STUB: not implemented"
	return nil
}

// CT IP address calculated from Pod's address.
type CTIPAddress net.IP

func (a *CTIPAddress) GetMatchKey(addrType types.AddressType) *types.MatchKey {
	_ = "STUB: not implemented"
	return nil
}

func (a *CTIPAddress) GetMatchValue() string { _ = "STUB: not implemented"; return "" }

func (a *CTIPAddress) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func NewCTIPAddress(addr net.IP) *CTIPAddress { _ = "STUB: not implemented"; return nil }

// CT IP block calculated from Pod's address.
type CTIPNetAddress net.IPNet

func (a *CTIPNetAddress) GetMatchKey(addrType types.AddressType) *types.MatchKey {
	_ = "STUB: not implemented"
	return nil
}

func (a *CTIPNetAddress) GetMatchValue() string { _ = "STUB: not implemented"; return "" }

func (a *CTIPNetAddress) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func NewCTIPNetAddress(addr net.IPNet) *CTIPNetAddress { _ = "STUB: not implemented"; return nil }

type LabelIDAddress uint32

func (a *LabelIDAddress) GetMatchKey(addrType types.AddressType) *types.MatchKey {
	_ = "STUB: not implemented"
	return nil
}

func (a *LabelIDAddress) GetMatchValue() string { _ = "STUB: not implemented"; return "" }

func (a *LabelIDAddress) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func NewLabelIDAddress(labelID uint32) *LabelIDAddress { _ = "STUB: not implemented"; return nil }

// ConjunctionNotFound is an error response when the specified policyRuleConjunction is not found from the local cache.
type ConjunctionNotFound uint32

func (e *ConjunctionNotFound) Error() string { _ = "STUB: not implemented"; return "" }

func newConjunctionNotFound(conjunctionID uint32) *ConjunctionNotFound {
	_ = "STUB: not implemented"
	return nil
}

// conjunctiveMatch generates match conditions for conjunctive match flow entry, including source or destination
// IP address, ofport number of OVS interface, or Service port. When conjunctiveMatch is used to match IP
// address or ofport number, matchProtocol is "ip". When conjunctiveMatch is used to match Service
// port, matchProtocol is Service protocol. If Service protocol is not set, "tcp" is used by default.
type conjunctiveMatch struct {
	tableID    uint8
	priority   *uint16
	matchPairs []matchPair
}

type matchPair struct {
	matchKey   *types.MatchKey
	matchValue interface{}
}

func (m *matchPair) KeyString() string { _ = "STUB: not implemented"; return "" }

// Use the unique format "x.x.x.x/xx" for IP address and IP net, to avoid generating two different global map
// keys for IP and IP/mask. Use MatchDstIPNet/MatchSrcIPNet as match type to generate global cache key for both IP
// and IPNet. This is because OVS treats IP and IP/$maskLen as the same condition (maskLen=32 for an IPv4 address,
// and maskLen=128 for an IPv6 address). If Antrea has two different conjunctive match flow contexts, only one
// flow entry is installed on OVS, and the conjunctive actions in the first context wil be overwritten by those
// in the second one.

// To normalize the key, set full mask while a single port is provided.

// This case includes the matchValue is ICMPType or ICMPCode.

// The default cases include the matchValue is an ofport Number.

func (m *conjunctiveMatch) generateGlobalMapKey() string { _ = "STUB: not implemented"; return "" }

// changeType is generally used to describe the change type of a conjMatchFlowContext. It is also used in "flowChange"
// to describe the expected OpenFlow operation which needs to be applied on the OVS bridge, and used in "actionChange"
// to describe the policyRuleConjunction is expected to be added to or removed from conjMatchFlowContext's actions.
// The value of changeType could be creation, modification, and deletion.
type changeType int

const (
	insertion changeType = iota
	modification
	deletion
)

// flowChange stores the expected OpenFlow entry and flow operation type which need to be applied on the OVS bridge.
// The "flow" in flowChange should be nil if there is no change on the OpenFlow entry. A possible case is that a
// DENY-ALL rule is required by a policyRuleConjunction, the flowChange will update the in-memory cache, but will not
// change on OVS.
type flowChange struct {
	flow       *openflow15.FlowMod
	changeType changeType
}

// actionChange stores the changed action of the conjunctive match flow, and the change type.
// The "action" in actionChange is not nil.
type actionChange struct {
	action     *conjunctiveAction
	changeType changeType
}

// conjunctiveAction generates the policyRuleConjunction action in Openflow entry. The flow action is like
// policyRuleConjunction(conjID,clauseID/nClause) when it has been realized on the switch.
type conjunctiveAction struct {
	conjID   uint32
	clauseID uint8
	nClause  uint8
}

// conjMatchFlowContext generates conjunctive match flow entries for conjunctions share the same match conditions.
// One conjMatchFlowContext is responsible for one specific conjunctive match flow entry. As the match condition
// of the flow entry can be shared by different conjunctions, the realized Openflow entry might have multiple
// conjunctive actions. If the dropTable is not nil, conjMatchFlowContext also installs a drop flow in the dropTable.
type conjMatchFlowContext struct {
	// conjunctiveMatch describes the match condition of conjunctive match flow entry.
	*conjunctiveMatch
	// actions is a map from policyRuleConjunction ID to conjunctiveAction. It records all the conjunctive actions in
	// the conjunctive match flow. When the number of actions is reduced to 0, the conjMatchFlowContext.flow is
	// uninstalled from the switch.
	actions map[uint32]*conjunctiveAction
	// denyAllRules is a set to cache the "DENY-ALL" rules that is applied to the matching address in this context.
	denyAllRules         map[uint32]bool
	featureNetworkPolicy *featureNetworkPolicy
	// flow is the conjunctive match flow built from this context. flow needs to be updated if actions are changed.
	flow *openflow15.FlowMod
	// dropflow is the default drop flow built from this context to drop packets in the AppliedToGroup but not pass the
	// NetworkPolicy rule. dropFlow is installed on the switch as long as either actions or denyAllRules is not
	// empty, and uninstalled when both two are empty. When the dropFlow is uninstalled from the switch, the
	// conjMatchFlowContext is removed from the cache.
	dropFlow *openflow15.FlowMod
	// dropFlowEnableLogging describes the logging requirement of the dropFlow.
	dropFlowEnableLogging bool
}

// createOrUpdateConjunctiveMatchFlow creates or updates the conjunctive match flow with the latest actions. It returns
// the flowChange including the changed OpenFlow entry and the expected operation which need to be applied on the OVS bridge.
func (ctx *conjMatchFlowContext) createOrUpdateConjunctiveMatchFlow(actions []*conjunctiveAction) *flowChange {
	_ = "STUB: not implemented"
	// Check if flow is already installed. If not, create a new flow.
	return nil
}

// Check the number of valid conjunctiveActions, and return nil immediately if it is 0. It happens when the match
// condition is used only for matching AppliedToGroup, but no From or To is defined in the NetworkPolicy rule.

// Create the conjunctive match flow entry. The actions here should not be empty for either add or update case.
// The expected operation for a new Openflow entry should be "insertion".

// Modify the existing Openflow entry and reset the actions.

// The expected operation for an existing Openflow entry should be "modification".

// deleteAction deletes the specified policyRuleConjunction from conjunctiveMatchFlow's actions, and then returns the
// flowChange.
func (ctx *conjMatchFlowContext) deleteAction(conjID uint32) *flowChange {
	_ = "STUB: not implemented"
	// If the specified conjunctive action is the last one, delete the conjunctive match flow entry from the OVS bridge.
	// No need to check if the conjunction ID of the only conjunctive action is the specified ID or not, as it
	// has been checked in the caller.
	return nil
}

// Modify the Openflow entry and reset the other conjunctive actions.

// addAction adds the specified policyRuleConjunction into conjunctiveMatchFlow's actions, and then returns the flowChange.
func (ctx *conjMatchFlowContext) addAction(action *conjunctiveAction) *flowChange {
	_ = "STUB: not implemented"
	// Check if the conjunction exists in conjMatchFlowContext actions or not. If yes, return nil immediately.
	return nil
}

// Append current conjunctive action to the existing actions, and then calculate the conjunctive match flow changes.

func (ctx *conjMatchFlowContext) addDenyAllRule(ruleID uint32) { _ = "STUB: not implemented"; return }

func (ctx *conjMatchFlowContext) delDenyAllRule(ruleID uint32) {
	_ = "STUB: not implemented"
	// Delete the DENY-ALL rule if it is in context denyAllRules.
	return
}

// conjMatchFlowContextChange describes the changes of a conjMatchFlowContext. It is generated when a policyRuleConjunction
// is added, deleted, or the addresses in an existing policyRuleConjunction are changed. The changes are calculated first,
// and then applied on the OVS bridge using a single Bundle, and lastly the local cache is updated. The local cahce
// is updated only if conjMatchFlowContextChange is applied on the OVS bridge successfully.
type conjMatchFlowContextChange struct {
	// context is the changed conjMatchFlowContext, which needs to be updated after the OpenFlow entries are applied to
	// the OVS bridge. context is not nil.
	context *conjMatchFlowContext
	// ctxChangeType is the changed type of the conjMatchFlowContext. The possible values are "creation", "modification"
	// and "deletion". Add the context into the globalConjMatchFlowCache if the ctxChangeType is "insertion", and remove
	// from the globalConjMatchFlowCache if it is "deletion".
	ctxChangeType changeType
	// matchFlow is the changed conjunctive match flow which needs to be realized on the OVS bridge. It is used to update
	// conjMatchFlowContext.flow. matchFlow is set if the conjunctive match flow needs to be updated on the OVS bridge, or
	// a DENY-ALL rule change is required by the policyRuleConjunction. matchFlow is nil if the policyRuleConjunction
	// is already added/removed in the conjMatchFlowContext's actions or denyAllRules.
	matchFlow *flowChange
	// dropFlow is the changed drop flow which needs to be realized on the OVS bridge. It is used to update
	// conjMatchFlowContext.dropFlow. dropFlow is set when the default drop flow needs to be added or removed on the OVS
	// bridge, and it is nil in other cases.
	dropFlow *flowChange
	// clause is the policyRuleConjunction's clause having current conjMatchFlowContextChange. It is used to update the
	// mapping relations between the policyRuleConjunction and the conjMatchFlowContext. Update the clause.matches after
	// the conjMatchFlowContextChange is realized on the OVS bridge. clause is not nil.
	clause *clause
	// actChange is the changed conjunctive action. It is used to update the conjMatchFlowContext's actions. actChange
	// is not nil.
	actChange *actionChange
}

// updateContextStatus changes conjMatchFlowContext's status, including,
//  1. reset flow and dropFlow after the flow changes have been applied to the OVS bridge,
//  2. modify the actions with the changed action,
//  3. update the mapping of denyAllRules and corresponding policyRuleConjunction,
//  4. add the new conjMatchFlowContext into the globalConjMatchFlowCache, or remove the deleted conjMatchFlowContext
//     from the globalConjMatchFlowCache.
func (c *conjMatchFlowContextChange) updateContextStatus() { _ = "STUB: not implemented"; return }

// Update clause.matches with the conjMatchFlowContext, and update conjMatchFlowContext.actions with the changed
// conjunctive action.

// Update the match flow in the conjMatchFlowContext. There are two kinds of possible changes on the match flow:
// 1) A conjunctive match flow change required by the policyRuleConjunction.
// 2) A DENY-ALL rule required by the policyRuleConjunction.
// For 1), conjMatchFlowContext.Flow should be updated with the conjMatchFlowContextChange.matchFlow.flow.
// For 2), append or delete the conjunction ID from the conjMatchFlowContext's denyAllRules.

// Update conjMatchFlowContext.dropFlow.

// Update globalConjMatchFlowCache. Add the conjMatchFlowContext into the globalConjMatchFlowCache if the ctxChangeType
// is "insertion", or delete from the globalConjMatchFlowCache if the ctxChangeType is "deletion".

// policyRuleConjunction is responsible to build Openflow entries for Pods that are in a NetworkPolicy rule's AppliedToGroup.
// The Openflow entries include conjunction action flows, conjunctive match flows, and default drop flows in the dropTable.
// NetworkPolicyController will make sure only one goroutine operates on a policyRuleConjunction.
//  1. Conjunction action flows use policyRuleConjunction ID as match condition. policyRuleConjunction ID is the single
//     match condition for conjunction action flows to allow packets. If the NetworkPolicy rule has also configured excepts
//     in From or To, Openflow entries are installed only for diff IPBlocks between From/To and Excepts. These are added as
//     conjunctive match flows as described below.
//  2. Conjunctive match flows adds conjunctive actions in Openflow entry, and they are grouped by clauses. The match
//     condition in one clause is one of these three types: from address(for fromClause), or to address(for toClause), or
//     service ports(for serviceClause) configured in the NetworkPolicy rule. Each conjunctive match flow entry is
//     maintained by one specific conjMatchFlowContext which is stored in globalConjMatchFlowCache, and shared by clauses
//     if they have the same match conditions. clause adds or deletes conjunctive action to conjMatchFlowContext actions.
//     A clause is hit if the packet matches any conjunctive match flow that are grouped by this clause. Conjunction
//     action flow is hit only if all clauses in the policyRuleConjunction are hit.
//  3. Default drop flows are also maintained by conjMatchFlowContext. It is used to drop packets sent from or to the
//     AppliedToGroup but not pass the Network Policy rule.
type policyRuleConjunction struct {
	id            uint32
	fromClause    *clause
	toClause      *clause
	serviceClause *clause
	actionFlows   []*openflow15.FlowMod
	metricFlows   []*openflow15.FlowMod
	// NetworkPolicy reference information for debugging usage, its value can be nil
	// for conjunctions that are not built for a specific NetworkPolicy, e.g. DNS packetin Conjunction.
	npRef        *v1beta2.NetworkPolicyReference
	ruleName     string
	ruleTableID  uint8
	ruleLogLabel string
}

// clause groups conjunctive match flows. Matches in a clause represent source addresses(for fromClause), or destination
// addresses(for toClause) or service ports(for serviceClause) in a NetworkPolicy rule. When the new address or service
// port is added into the clause, it adds a new conjMatchFlowContext into globalConjMatchFlowCache (or finds the
// existing one from globalConjMatchFlowCache), and then update the key of the conjunctiveMatch into its own matches.
// When address is deleted from the clause, it deletes the conjunctive action from the conjMatchFlowContext,
// and then deletes the key of conjunctiveMatch from its own matches.
type clause struct {
	action *conjunctiveAction
	// matches is a map from the unique string generated from the conjunctiveMatch to conjMatchFlowContext. It is used
	// to cache conjunctive match conditions in the same clause.
	matches map[string]*conjMatchFlowContext
	// ruleTable is where to install conjunctive match flows.
	ruleTable binding.Table
	// dropTable is where to install Openflow entries to drop the packet sent to or from the AppliedToGroup but does not
	// satisfy any conjunctive match conditions. It should be nil, if the clause is used for matching service port.
	dropTable binding.Table
}

func (c *client) NewDNSPacketInConjunction(id uint32) error { _ = "STUB: not implemented"; return nil }

// Use ct_state=+trk+rpl as matching condition.
// CTState bit-state map:
// dnat | snat | trk | inv | rpl | rel | est | new

// Add CTState for UDP as well to make sure only solicited DNS responses are sent
// to userspace.

// Add the policyRuleConjunction into policyCache

func (c *client) AddAddressToDNSConjunction(id uint32, addrs []types.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) DeleteAddressFromDNSConjunction(id uint32, addrs []types.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *clause) addConjunctiveMatchFlow(featureNetworkPolicy *featureNetworkPolicy, match *conjunctiveMatch, enableLogging, isMCNPRule bool) *conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

// Get conjMatchFlowContext from globalConjMatchFlowCache. If it doesn't exist, create a new one and add into the cache.

// Generate the default drop flow if dropTable is not nil and the default drop flow is not set yet.

// Logging requirement of the rule has changed, modify default drop flow accordingly.

// Calculate the change on the conjMatchFlowContext.

// Append the conjunction to conjunctiveFlowContext's actions, and add the changed flow into the conjMatchFlowContextChange.

// Set the flowChange type as "insertion" but do not set flowChange.Flow. In this case, the policyRuleConjunction should
// be added into conjunctiveFlowContext's denyAllRules.

func generateAddressConjMatch(ruleTableID uint8, addr types.Address, addrType types.AddressType, priority *uint16) *conjunctiveMatch {
	_ = "STUB: not implemented"
	return nil
}

func generateServiceConjMatches(ruleTableID uint8, service v1beta2.Service, priority *uint16, ipProtocols []binding.Protocol) []*conjunctiveMatch {
	_ = "STUB: not implemented"
	return nil
}

func getServiceMatchPairs(service v1beta2.Service, ipProtocols []binding.Protocol) [][]matchPair {
	_ = "STUB: not implemented"
	return nil
}

// Since OVS only matches layer 3 IP address on the IGMP query packet, and doesn't
// identify the multicast group address set in the IGMP protocol, the flow entry
// processes all IGMP query packets by matching the destination IP address ( 224.0.0.1 )

// portsToBitRanges converts ports in Service to a list of BitRange.
func portsToBitRanges(port *intstr.IntOrString, endPort *int32) []types.BitRange {
	_ = "STUB: not implemented"
	return nil
}

// If `EndPort` is equal to `Port`, then treat it as single port case.

// Add several antrea range services based on a port range.

// Add single antrea service based on a single port.

// Match all ports with the given protocol type if `Port` and `EndPort` are not
// specified (value is 0).

// addAddrFlows translates the specified addresses to conjunctiveMatchFlows, and returns the corresponding changes on the
// conjunctiveMatchFlows.
func (c *clause) addAddrFlows(featureNetworkPolicy *featureNetworkPolicy, addrType types.AddressType, addresses []types.Address, priority *uint16, enableLogging, isMCNPRule bool) []*conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

// Calculate Openflow changes for the added addresses.

// addServiceFlows translates the specified Antrea Service to conjunctiveMatchFlow,
// and returns corresponding conjMatchFlowContextChange.
func (c *clause) addServiceFlows(featureNetworkPolicy *featureNetworkPolicy, services []v1beta2.Service, priority *uint16, enableLogging bool) []*conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

// deleteConjunctiveMatchFlow deletes the specific conjunctiveAction from existing flow.
func (c *clause) deleteConjunctiveMatchFlow(flowContextKey string) *conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

// Match is not located in clause cache. It happens if the conjMatchFlowContext is already deleted from clause local cache.

// Delete the conjunctive action if it is in context actions.

// Delete the DENY-ALL rule if it is in context denyAllRules.

// Uninstall default drop flow if the deleted conjunctiveAction is the last action or the rule is the last one in
// the denyAllRules.

// Remove the context from global cache if the match condition is not used by either DENEY-ALL or the conjunctive
// match flow.

// deleteAddrFlows deletes conjunctiveMatchFlow relevant to the specified addresses from local cache,
// and uninstalls Openflow entry.
func (c *clause) deleteAddrFlows(addrType types.AddressType, addresses []types.Address, priority *uint16) []*conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

// deleteAllMatches deletes all conjunctiveMatchFlow in the clause, and removes Openflow entry. deleteAllMatches
// is always invoked when NetworkPolicy rule is deleted.
func (c *clause) deleteAllMatches() []*conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

func (c *policyRuleConjunction) getAddressClause(addrType types.AddressType) *clause {
	_ = "STUB: not implemented"
	return nil
}

// InstallPolicyRuleFlows installs flows for a new NetworkPolicy rule. Rule should include all fields in the
// NetworkPolicy rule. Each ingress/egress policy rule installs Openflow entries on two tables, one for ruleTable and
// the other for dropTable. If a packet does not pass the ruleTable, it will be dropped by the dropTable.
// NetworkPolicyController will make sure only one goroutine operates on a PolicyRule and addresses in the rule.
// For a normal NetworkPolicy rule, these Openflow entries are installed: 1) 1 conjunction action flow; 2) multiple
// conjunctive match flows, the flow number depends on addresses in rule.From and rule.To, or if
// rule.FromExcepts/rule.ToExcepts are present, flow number is equal to diff of addresses between rule.From and
// rule.FromExcepts, and diff addresses between rule.To and rule.ToExcepts, and in addition number includes service ports
// in rule.Service; and 3) multiple default drop flows, the number is dependent on the addresses in rule.From for
// an egress rule, and addresses in rule.To for an ingress rule.
// For ALLOW-ALL rule, the Openflow entries installed on the switch are similar to a normal rule. The differences include,
// 1) rule.Service is nil; and 2) rule.To has only one address "0.0.0.0/0" for egress rule, and rule.From is "0.0.0.0/0"
// for ingress rule.
// For DENY-ALL rule, only the default drop flow is installed for the addresses in rule.From for egress rule, or
// addresses in rule.To for ingress rule. No conjunctive match flow or conjunction action except flows are installed.
// A DENY-ALL rule is configured with rule.ID, rule.Direction, and either rule.From(egress rule) or rule.To(ingress rule).
// Other fields in the rule should be nil.
// If there is an error in any clause's addAddrFlows or addServiceFlows, the conjunction action flow will never be hit.
// If the default drop flow is already installed before this error, all packets will be dropped by the default drop flow,
// Otherwise all packets will be allowed.
func (c *client) InstallPolicyRuleFlows(rule *types.PolicyRule) error {
	_ = "STUB: not implemented"
	return nil
}

// Add the policyRuleConjunction into policyCache

// calculateActionFlowChangesForRule calculates and updates the actionFlows for the conjunction corresponded to the ofPolicyRule.
func (f *featureNetworkPolicy) calculateActionFlowChangesForRule(rule *types.PolicyRule) *policyRuleConjunction {
	_ = "STUB: not implemented"
	return nil

	// Check if the policyRuleConjunction is added into cache or not. If yes, return nil.
}

// Conjunction action flows are installed only if the number of clauses in the conjunction is > 1. It should be a rule
// to drop all packets.  If the number is 1, no conjunctive match flows or conjunction action flows are installed,
// but the default drop flow is installed.

// Install action flows.

// calculateMatchFlowChangesForRule calculates the contextChanges for the policyRule, and updates the context status in case of batch install.
func (f *featureNetworkPolicy) calculateMatchFlowChangesForRule(conj *policyRuleConjunction, rule *types.PolicyRule) []*conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	// Calculate the conjMatchFlowContext changes. The changed Openflow entries are included in the conjMatchFlowContext change.
	return nil
}

// addRuleToConjunctiveMatch adds a rule's clauses to corresponding conjunctive match contexts.
// Unlike calculateMatchFlowChangesForRule, it updates the context status directly and doesn't calculate flow changes.
// It's used in initial batch install where we first add all rules then calculates flows change based on final state.
func (f *featureNetworkPolicy) addRuleToConjunctiveMatch(conj *policyRuleConjunction, rule *types.PolicyRule) {
	_ = "STUB: not implemented"
	return
}

// addActionToConjunctiveMatch adds a clause to corresponding conjunctive match context.
// It updates the context status directly and doesn't calculate the match flow, which is supposed to be calculated after
// all actions are added. It's used in initial batch install only.
func (f *featureNetworkPolicy) addActionToConjunctiveMatch(clause *clause, match *conjunctiveMatch, enableLogging, isMCNPRule bool) {
	_ = "STUB: not implemented"
	return
}

// Get conjMatchFlowContext from globalConjMatchFlowCache. If it doesn't exist, create a new one and add into the cache.

// Generate the default drop flow if dropTable is not nil.

// Add the conjunction to the conjunctiveFlowContext's actions.

// Add the conjunction ID to the conjunctiveFlowContext's denyAllRules.

// BatchInstallPolicyRuleFlows installs flows for NetworkPolicy rules in case of agent restart. It calculates and
// accumulates all Openflow entry updates required and installs all of them on OVS bridge in one bundle.
// It resets the global conjunctive match flow cache upon failure, and should NOT be used after any rule is installed
// via the InstallPolicyRuleFlows method. Otherwise the cache would be out of sync.
func (c *client) BatchInstallPolicyRuleFlows(ofPolicyRules []*types.PolicyRule) error {
	_ = "STUB: not implemented"
	return nil
}

// In theory there must be at least one action but InstallPolicyRuleFlows currently handles the 1 clause case
// and we do the same in addRuleToConjunctiveMatch. The check is added only for consistency. Later we should
// return error if clients install a rule with only 1 clause, and should remove the extra code for processing it.

// Send the changed Openflow entries to the OVS bridge.

// Reset the global conjunctive match flow cache since the OpenFlow bundle, which contains
// all the match flows to be installed, was not applied successfully.

// Update conjMatchFlowContexts as the expected status.

// Add the policyRuleConjunction into policyCache

// applyConjunctiveMatchFlows installs OpenFlow entries on the OVS bridge, and then updates the conjMatchFlowContext.
func (f *featureNetworkPolicy) applyConjunctiveMatchFlows(flowChanges []*conjMatchFlowContextChange) error {
	_ = "STUB: not implemented"
	// Send the OpenFlow entries to the OVS bridge.
	return nil
}

// Update conjunctiveMatchContext.

// sendConjunctiveFlows sends all the changed OpenFlow entries to the OVS bridge in a single Bundle.
func (f *featureNetworkPolicy) sendConjunctiveFlows(changes []*conjMatchFlowContextChange) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve the OpenFlow entries from the flowChanges.

// ActionFlowPriorities returns the OF priorities of the actionFlows in the policyRuleConjunction
func (c *policyRuleConjunction) ActionFlowPriorities() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *policyRuleConjunction) newClause(clauseID uint8, nClause uint8, ruleTable, dropTable binding.Table) *clause {
	_ = "STUB: not implemented"
	return nil
}

// calculateClauses configures the policyRuleConjunction's clauses according to the PolicyRule. The Openflow entries are
// not installed on the OVS bridge when calculating the clauses.
func (c *policyRuleConjunction) calculateClauses(rule *types.PolicyRule) (uint8, binding.Table, binding.Table) {
	_ = "STUB: not implemented"
	return 0, *new(binding.Table), *new(binding.Table)
}

// Calculate clause IDs and the total number of clauses.

// deny rule does not need to be created for ClusterNetworkPolicies

// calculateChangesForRuleCreation returns the conjMatchFlowContextChanges of the new policyRuleConjunction. It
// will calculate the expected conjMatchFlowContext status, and the changed Openflow entries.
func (c *policyRuleConjunction) calculateChangesForRuleCreation(featureNetworkPolicy *featureNetworkPolicy, rule *types.PolicyRule) []*conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

func containsLabelIdentityAddress(addresses []types.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// calculateChangesForRuleDeletion returns the conjMatchFlowContextChanges of the deleted policyRuleConjunction. It
// will calculate the expected conjMatchFlowContext status, and the changed Openflow entries.
func (c *policyRuleConjunction) calculateChangesForRuleDeletion() []*conjMatchFlowContextChange {
	_ = "STUB: not implemented"
	return nil
}

// getAllFlowKeys returns the match strings used in ovs-ofctl dump-flows command, including
// actions flows of policyRuleConjunction, as well as matching flows of all its clauses.
func (c *policyRuleConjunction) getAllFlowKeys() []string { _ = "STUB: not implemented"; return nil }

// Add flows in the order of action flows, conjunctive match flows, drop flows.

func (f *featureNetworkPolicy) getPolicyRuleConjunction(ruleID uint32) *policyRuleConjunction {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) GetPolicyInfoFromConjunction(ruleID uint32) (bool, *v1beta2.NetworkPolicyReference, string, string, string) {
	_ = "STUB: not implemented"
	return false, nil, "", "", ""
}

// UninstallPolicyRuleFlows removes the Openflow entry relevant to the specified NetworkPolicy rule.
// It also returns a slice of stale ofPriorities used by ClusterNetworkPolicies.
// UninstallPolicyRuleFlows will do nothing if no Openflow entry for the rule is installed.
func (c *client) UninstallPolicyRuleFlows(ruleID uint32) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete action flows from the OVS bridge.

// Get the conjMatchFlowContext changes.

// Send the changed OpenFlow entries to the OVS bridge and update the conjMatchFlowContext.

// getStalePriorities returns the ofPriorities that will be stale on the rule table where the
// policyRuleConjunction is installed, after the deletion of that policyRuleConjunction.
func (f *featureNetworkPolicy) getStalePriorities(conj *policyRuleConjunction) (staleOFPriorities []string) {
	_ = "STUB: not implemented"
	return nil
}

// Filter out all the policyRuleConjuctions created at the ofPriority across all CNP tables.

// There are other policyRuleConjuctions in the same table created with this
// ofPriority. The ofPriority is thus not stale and cannot be released.

func (f *featureNetworkPolicy) replayFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

// AddPolicyRuleAddress adds one or multiple addresses to the specified NetworkPolicy rule. If addrType is srcAddress, the
// addresses are added to PolicyRule.From, else to PolicyRule.To.
func (c *client) AddPolicyRuleAddress(ruleID uint32, addrType types.AddressType, addresses []types.Address, priority *uint16, enableLogging, isMCNPRule bool) error {
	_ = "STUB: not implemented"
	return nil
}

// If policyRuleConjunction doesn't exist in client's policyCache return not found error. It should not happen, since
// NetworkPolicyController will guarantee the policyRuleConjunction is created before this method is called. The check
// here is for safety.

// Check if the clause is nil or not. The clause is nil if the addrType is an unsupported type.

// DeletePolicyRuleAddress removes addresses from the specified NetworkPolicy rule. If addrType is srcAddress, the addresses
// are removed from PolicyRule.From, else from PolicyRule.To.
func (c *client) DeletePolicyRuleAddress(ruleID uint32, addrType types.AddressType, addresses []types.Address, priority *uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// If policyRuleConjunction doesn't exist in client's policyCache return not found error. It should not happen, since
// NetworkPolicyController will guarantee the policyRuleConjunction is created before this method is called. The check
//	here is for safety.

// Check if the clause is nil or not. The clause is nil if the addrType is an unsupported type.

// Remove policyRuleConjunction to actions of conjunctive match using specific address.

// Update the Openflow entries on the OVS bridge, and update local cache.

func (c *client) GetNetworkPolicyFlowKeys(npName, npNamespace string, npType v1beta2.NetworkPolicyType) []string {
	_ = "STUB: not implemented"
	return nil

	// Hold replayMutex write lock to protect flows from being modified by
	// NetworkPolicy updates and replayFlows. This is more for logic
	// cleanliness, as: for now flow updates do not impact the matching string
	// generation; NetworkPolicy updates do not change policyRuleConjunction.actionFlows;
	// and last for protection of clause flows, conjMatchFlowLock is good enough.
}

// If the NetworkPolicyReference in the policyRuleConjunction is nil then that entry in client's
// policyCache should be ignored because here we need to dump flows of NetworkPolicy.

// There can be duplicated flows added due to conjunctive matches
// shared by multiple policy rules (clauses).

// flowUpdates stores updates to the actionFlows and matchFlows in a policyRuleConjunction.
type flowUpdates struct {
	newActionFlows []*openflow15.FlowMod
	newPriority    uint16
}

// getMatchFlowUpdates calculates the update for conjuctiveMatchFlows in a policyRuleConjunction to be
// installed on a new priority.
func getMatchFlowUpdates(conj *policyRuleConjunction, newPriority uint16) (add, del []*openflow15.FlowMod) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processFlowUpdates identifies the update cases in flow adds and deletes.
// For conjunctiveMatchFlow updates, the following scenario is possible:
//
// A flow {priority=100,ip,reg1=0x1f action=conjunction(1,1/3)} need to be re-assigned priority=99.
// In this case, an addFlow of <priority=99,ip,reg1=0x1f> and delFlow <priority=100,ip,reg1=0x1f> will be issued.
// At the same time, another flow {priority=99,ip,reg1=0x1f action=conjunction(2,1/3)} exists and now needs to
// be re-assigned priority 98. This operation will issue a delFlow <priority=99,ip,reg1=0x1f>, which
// would essentially void the add flow for conj=1.
//
// In this case, we remove the conflicting delFlow and set addFlow as a modifyFlow.
func (f *featureNetworkPolicy) processFlowUpdates(addFlows, delFlows []*openflow15.FlowMod) (add, update, del []*openflow15.FlowMod) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// treat the addFlow as update

// remove the delFlow from the list

// reset list index as delFlows[i] is removed

// updateConjunctionActionFlows constructs a new policyRuleConjunction with actionFlows updated to be
// stored in the policyCache.
func (f *featureNetworkPolicy) updateConjunctionActionFlows(conj *policyRuleConjunction, updates flowUpdates) *policyRuleConjunction {
	_ = "STUB: not implemented"
	return nil
}

// updateConjunctionMatchFlows updates the conjuctiveMatchFlows in a policyRuleConjunction.
func (f *featureNetworkPolicy) updateConjunctionMatchFlows(conj *policyRuleConjunction, newPriority uint16) {
	_ = "STUB: not implemented"
	return
}

// update the globalConjMatchFlowCache so that the keys are updated

// calculateFlowUpdates calculates the flow updates required for the priority re-assignments specified in the input map.
func (f *featureNetworkPolicy) calculateFlowUpdates(updates map[uint16]uint16, table uint8) (addFlows, delFlows []*openflow15.FlowMod,
	conjFlowUpdates map[uint32]flowUpdates) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Only re-assign flow priorities for flows in the table specified.

// The OF flow was created at the priority which need to be re-installed
// at the NewPriority now

// Store the actionFlow update to the policyRuleConjunction and update all
// policyRuleConjunctions if flow installation is successful.

// ReassignFlowPriorities takes a list of priority updates, and update the actionFlows to replace
// the old priority with the desired one, for each priority update.
func (c *client) ReassignFlowPriorities(updates map[uint16]uint16, table uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func parseMulticastIngressPodFlow(flowMap map[string]string) (uint32, types.RuleMetric) {
	_ = "STUB: not implemented"
	return 0, *new(types.RuleMetric)
}

func parseMulticastEgressPodFlow(flowMap map[string]string) (string, types.RuleMetric) {
	_ = "STUB: not implemented"
	return "", *new(types.RuleMetric)
}

func parseMulticastMetricFlow(flowMap map[string]string) (uint32, types.RuleMetric, error) {
	_ = "STUB: not implemented"
	// example MulticastEgressMetric allow flow format:
	// table=MulticastEgressMetric, n_packets=11, n_bytes=1562, priority=200,reg0=0x400/0x400,reg3=0x4 actions=goto_table:MulticastEgressPodMetric
	// example MulticastEgressMetric drop flow format:
	// table=MulticastEgressMetric, n_packets=230, n_bytes=32660, priority=200,reg0=0x400/0x400,reg3=0x3 actions=drop
	// example MulticastIngressMetric allow flow format:
	// table=MulticastIngressMetric, n_packets=18, n_bytes=780, priority=200,reg0=0x400/0x400,reg3=0x3 actions=resubmit(,MulticastOutput)
	//
	// Valid flows always match APConjIDField (reg3). Stale non-multicast flows can share the same
	// table id and priority=200 after a pipeline shift; they lack a usable conjunction id and must
	// not be attributed as multicast policy metrics (cookie filtering also skips wrong categories).
	return 0, *new(types.RuleMetric), nil
}

func parseFlowMetric(flowMap map[string]string) types.RuleMetric {
	_ = "STUB: not implemented"
	return *new(types.RuleMetric)
}

func parseDropFlow(flowMap map[string]string) (uint32, types.RuleMetric, error) {
	_ = "STUB: not implemented"
	// Deny metric flows match reg0 (deny mark) and encode the conjunction id in reg3; see denyRuleMetricFlow.
	return 0, *new(types.RuleMetric), nil
}

func parseAllowFlow(flowMap map[string]string) (uint32, types.RuleMetric, error) {
	_ = "STUB: not implemented"
	// A valid allow metric flow must have ct_label, which encodes the rule ID.
	// Flows without ct_label are not metric flows — they may be stale flows from a
	// prior agent version that occupied the same table ID after a pipeline table shift
	// (e.g., L3 routing flows from an old pipeline that now share the EgressMetricTable
	// ID after a new table was inserted earlier in the pipeline). Skip them to avoid
	// misinterpreting non-metric flows as NetworkPolicy statistics.
	return 0, *new(types.RuleMetric), nil
}

// ct_state=+new

// ct_label is absent or malformed; this is not a valid allow metric flow.

// only 32 bits are valid.

func parseFlowToMap(flow string) map[string]string { _ = "STUB: not implemented"; return nil }

// Some substrings spilt by "," may have no "=", for instance, if "resubmit(,70)" is present.

// There is space not comma before actions, which causes troubles when parsing value, example:
// table=MulticastEgressRule, ...,conj_id=2 actions=load:0x2->NXM_NX_REG5[],load...

func parseMetricFlow(flowMap map[string]string) (uint32, types.RuleMetric, error) {
	_ = "STUB: not implemented"
	return 0,

		// example allow flow format:
		// table=101, n_packets=0, n_bytes=0, priority=200,ct_state=-new,ct_label=0x1/0xffffffff,ip actions=goto_table:105
		// example drop flow format:
		// table=101, n_packets=9, n_bytes=666, priority=200,reg0=0x100000/0x100000,reg3=0x5 actions=drop
		*new(types.RuleMetric), nil
}

// metricFlowsDumpFilter builds ovs-ofctl dump-flows FLOW match text for Antrea metric flows: exact
// cookie for the agent round and category. This excludes stale flows from prior rounds and flows
// from other feature categories (e.g. PodConnectivity, Traceflow) that may share the same table ID
// after a pipeline table shift. Note: "priority=N" is not a valid ovs-ofctl dump-flows keyword and
// must not be included; priority-based filtering is handled by parseMetricFlow and its callers.
func (c *client) metricFlowsDumpFilter(wantCategory cookie.Category) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *client) MulticastIngressPodMetrics() map[uint32]*types.RuleMetric {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) MulticastIngressPodMetricsByOFPort(ofPort int32) *types.RuleMetric {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) MulticastEgressPodMetrics() map[string]*types.RuleMetric {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) MulticastEgressPodMetricsByIP(ip net.IP) *types.RuleMetric {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) NetworkPolicyMetrics() map[uint32]*types.RuleMetric {
	_ = "STUB: not implemented"
	return nil
}

// Multicast ANP metric flows are installed by featureNetworkPolicy and use the same cookie
// category as unicast ANP metrics (NetworkPolicy), not cookie.Multicast (used by the multicast
// feature for IGMP / pod tables). Use NetworkPolicy in the dump filter or those flows are never matched.

// We have two flows for each allow rule. One matches 'ct_state=+new'
// and counts the number of first packets, which is also the number
// of sessions (this is the category why we have 2 flows). The other
// matches 'ct_state=-new' and is used to count all subsequent
// packets in the session. We need to merge metrics from these 2
// flows to get the correct number of total packets.

type featureNetworkPolicy struct {
	cookieAllocator       cookie.Allocator
	ipProtocols           []binding.Protocol
	bridge                binding.Bridge
	nodeType              config.NodeType
	l7NetworkPolicyConfig *config.L7NetworkPolicyConfig

	// globalConjMatchFlowCache is a global map for conjMatchFlowContext. The key is a string generated from the
	// conjMatchFlowContext.
	globalConjMatchFlowCache map[string]*conjMatchFlowContext
	conjMatchFlowLock        sync.Mutex // Lock for access globalConjMatchFlowCache
	// policyCache is a storage that supports listing policyRuleConjunction with different indexers.
	// It's guaranteed that one policyRuleConjunction is processed by at most one goroutine at any given time.
	policyCache cache.Indexer
	// egressTables map records all IDs of tables related to egress rules.
	egressTables map[uint8]struct{}

	cachedFlows *flowCategoryCache

	// loggingGroupCache is a storage for the logging groups, each one includes 2 buckets: one is to send the packet
	// to antrea-agent using the packetIn mechanism, the other is to force the packet to continue forwarding in the
	// OVS pipeline. The key is the next table used in the second bucket, and the value is the Openflow group.
	loggingGroupCache sync.Map
	groupAllocator    GroupAllocator

	ovsMetersAreSupported bool
	enableDenyTracking    bool
	enableAntreaPolicy    bool
	enableL7NetworkPolicy bool
	enableMulticast       bool
	proxyAll              bool
	ctZoneSrcField        *binding.RegField
	// deterministic represents whether to generate flows deterministically.
	// For example, if a flow has multiple actions, setting it to true can get consistent flow.
	// Enabling it may carry a performance impact. It's disabled by default and should only be used in testing.
	deterministic bool

	category cookie.Category
}

func (f *featureNetworkPolicy) getFeatureName() string { _ = "STUB: not implemented"; return "" }

func newFeatureNetworkPolicy(
	cookieAllocator cookie.Allocator,
	ipProtocols []binding.Protocol,
	bridge binding.Bridge,
	l7NetworkPolicyConfig *config.L7NetworkPolicyConfig,
	ovsMetersAreSupported,
	enableDenyTracking,
	enableAntreaPolicy bool,
	enableL7NetworkPolicy bool,
	enableMulticast bool,
	proxyAll bool,
	connectUplinkToBridge bool,
	nodeType config.NodeType,
	grpAllocator GroupAllocator) *featureNetworkPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureNetworkPolicy) initFlows() []*openflow15.FlowMod {
	_ = "STUB: not implemented"
	return nil
}

// skipPolicyRuleCheckFlows generates the flows to forward the packets in an established or related
// connections to the metric table in the same stage directly, so that these packets would skip the flows
// for NetworkPolicy rules.
func (f *featureNetworkPolicy) skipPolicyRuleCheckFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureNetworkPolicy) l7NPTrafficControlFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

// This generates the flow to output the packets returned from an application-aware engine return port to their
// original target ofPort. It has the highest priority to prevent these packets from being matched by other
// flows in this table that redirect packets to an application-aware engine target ofPort again.

// This generates the flow to output the packets marked with L7NPRedirectCTMark to an application-aware engine
// via the target ofPort. Note that, before outputting the packets, VLAN ID stored on field L7NPRuleVlanIDCTMarkField
// will be copied to VLAN ID register (OXM_OF_VLAN_VID) to set VLAN ID of the packets.

// This generates the flow to mark the packets from an application-aware engine via the return ofPort and forward
// the packets to stageConntrackState directly. Note that, for the packets which are originally destined for a
// tunnel port, value of NXM_NX_TUN_IPV4_DST needs to be loaded in L3ForwardingTable of stageRouting.

// This generates the flow to forward the returned packets (with FromL7NPReturnRegMark) to stageOutput directly
// after loading output port number to reg1 in L2ForwardingCalcTable.

// This generates the flow to restore the corresponding connection tracking (CT) state, excluding NAT, of
// packets and resubmit them back to the ConntrackTable. CtStateNotRestoredRegMark and CtStateRestoredRegMark
// are used to prevent packets from being resubmitted in a cyclic manner, ensuring that the packets are
// resubmitted only once.

// This generates the flow to match the reply packets returning from the application-aware engine. If the
// packets belong to a Service connection, DNAT is applied to restore the original source IP (Service IP) with
// a connection tracking (CT) action. If the packets don't belong to a Service connection, the CT action is
// harmless and will have no effect on those packets.
// The reason why the target table is L3ForwardingTable is that for the packets which are originally destined
// for a tunnel port, the value of NXM_NX_TUN_IPV4_DST needs to be loaded in L3ForwardingTable in stageRouting.
// This ensures that the correct tunnel destination IP is used for these packets, enabling proper forwarding
// through the tunnel.

// This generates the flow to match the request packets returning from the application-aware engine. A
// connection tracking (CT) action with NAT is not needed because the DNAT has been done before they are
// redirected to the application-aware engine. The reason why the target table is L3ForwardingTable is the
// same as above.

// This generates the flow to match the reply packets that should be redirected to the application-aware engine.
// A connection tracking (CT) action with NAT is not needed because the DNAT will be done after they are returned
// from the application-aware engine.

// This generates the flow to match request packets that are going to be redirected to the application-aware
// engine. A connection tracking (CT) action with NAT is needed because the DNAT should be done before they
// are redirected to the application-aware engine.

func (f *featureNetworkPolicy) loggingNPPacketFlowWithOperations(cookieID uint64, operations uint8) binding.Flow {
	_ = "STUB: not implemented"
	return *new(binding.Flow)
}

func (f *featureNetworkPolicy) initLoggingFlows() []binding.Flow {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureNetworkPolicy) initGroups() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

// Create OpenFlow group to both log Antrea-native policy events and resubmit the packet to nextTable.

// There are two buckets in this type All group which have generated two copies of the packet: 1) resubmit
// the first one back to the next table of which the original packet consumes the group, then it is able to
// continue with the previous forwarding logic; 2) set packetIn marks in the second one and resubmit it to
// OutputTable.

func (f *featureNetworkPolicy) replayMeters() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}

func (f *featureNetworkPolicy) getLoggingAndResubmitGroupID(nextTable uint8) binding.GroupIDType {
	_ = "STUB: not implemented"
	return *new(binding.GroupIDType)
}

func (f *featureNetworkPolicy) replayGroups() []binding.OFEntry {
	_ = "STUB: not implemented"
	return nil
}
