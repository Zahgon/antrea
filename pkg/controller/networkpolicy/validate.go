// Copyright 2020 Antrea Authors
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

package networkpolicy

import (
	"regexp"

	admv1 "k8s.io/api/admission/v1"
	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/network-policy-api/apis/v1alpha1"

	crdv1beta1 "antrea.io/antrea/v2/pkg/apis/crd/v1beta1"
)

// validator interface introduces the set of functions that must be implemented
// by any resource validator.
type validator interface {
	// createValidate is the interface which must be satisfied for resource
	// CREATE events.
	createValidate(curObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool)
	// updateValidate is the interface which must be satisfied for resource
	// UPDATE events.
	updateValidate(curObj, oldObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool)
	// deleteValidate is the interface which must be satisfied for resource
	// DELETE events.
	deleteValidate(oldObj interface{}, userInfo authenticationv1.UserInfo) (string, bool)
}

// resourceValidator maintains a reference of the NetworkPolicyController and
// provides a base struct for validating objects which implement the validator
// interface.
type resourceValidator struct {
	networkPolicyController *NetworkPolicyController
}

// antreaPolicyValidator implements the validator interface for Antrea-native
// policies.
type antreaPolicyValidator resourceValidator

// tierValidator implements the validator interface for Tier resources.
type tierValidator resourceValidator

// groupValidator implements the validator interface for the ClusterGroup resource.
type groupValidator resourceValidator

// adminPolicyValidator implements the validator interface for the AdminNetworkPolicy resource.
type adminPolicyValidator resourceValidator

var (
	// reservedTierPriorities stores the reserved priority range from 251, 252, 254 and 255.
	// The priority 250 is reserved for default Tier but not part of this set in order to be
	// able to create the Tier by Antrea. Same for priority 253 which is reserved for the
	// baseline tier.
	reservedTierPriorities = sets.New[int32](int32(251), int32(252), int32(254), int32(255))
	// reservedTierNames stores the set of Tier names which cannot be deleted
	// since they are created by Antrea.
	reservedTierNames = sets.New[string]("baseline", "application", "platform", "networkops", "securityops", "emergency")
	// allowedFQDNChars validates that the matchPattern field contains only valid DNS characters
	// and the wildcard '*' character.
	allowedFQDNChars = regexp.MustCompile("^[-0-9a-zA-Z.*]+$")
)

// RegisterAntreaPolicyValidator registers an Antrea-native policy validator
// to the resource registry. A new validator must be registered by calling
// this function before the Run phase of the APIServer.
func (v *NetworkPolicyValidator) RegisterAntreaPolicyValidator(a validator) {
	_ = "STUB: not implemented"
	return
}

// RegisterTierValidator registers a Tier validator to the resource registry.
// A new validator must be registered by calling this function before the Run
// phase of the APIServer.
func (v *NetworkPolicyValidator) RegisterTierValidator(t validator) {
	_ = "STUB: not implemented"
	return
}

// RegisterGroupValidator registers a Group validator to the resource registry.
// A new validator must be registered by calling this function before the Run
// phase of the APIServer.
func (v *NetworkPolicyValidator) RegisterGroupValidator(g validator) {
	_ = "STUB: not implemented"
	return
}

func (v *NetworkPolicyValidator) RegisterAdminNetworkPolicyValidator(a validator) {
	_ = "STUB: not implemented"
	return
}

// NetworkPolicyValidator maintains list of validator objects which validate
// the Antrea-native policy related resources.
type NetworkPolicyValidator struct {
	// antreaPolicyValidators maintains a list of validator objects which
	// implement the validator interface for Antrea-native policies.
	antreaPolicyValidators []validator
	// tierValidators maintains a list of validator objects which
	// implement the validator interface for Tier resources.
	tierValidators []validator
	// groupValidators maintains a list of validator objects which
	// implement the validator interface for ClusterGroup resources.
	groupValidators []validator
	// adminNPValidators maintains a list of validator objects which
	// implement the validator interface for ANP and BANP resources.
	adminNPValidators []validator
}

// NewNetworkPolicyValidator returns a new *NetworkPolicyValidator.
func NewNetworkPolicyValidator(networkPolicyController *NetworkPolicyController) *NetworkPolicyValidator {
	_ = "STUB: not implemented"
	// initialize the validator registry with the default validators that need to
	// be called.
	return nil
}

// apv is an instance of antreaPolicyValidator to validate Antrea-native
// policy events.

// tv is an instance of tierValidator to validate Tier resource events.

// gv is an instance of groupValidator to validate ClusterGroup
// resource events.

// Validate function validates a Group, ClusterGroup, Tier or Antrea Policy object
func (v *NetworkPolicyValidator) Validate(ar *admv1.AdmissionReview) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}

// Current serving versions of Tier are v1alpha1 and v1beta1. They have the same
// schema and the same validating logic, and we only store v1beta1 in the etcd. So
// we unmarshal both of them into a v1beta1 object to do validation.

// Current serving versions of ClusterGroup are v1alpha3 and v1beta1. They have
// the same schema and the same validating logic, and we only store v1beta1 in
// the etcd. So we unmarshal both of them into a v1beta1 object to do validation.

// Current serving versions of Group are v1alpha3 and v1beta1. They have the same
// schema and the same validating logic, and we only store v1beta1 in the etcd. So
// we unmarshal both of them into a v1beta1 object to do validation.

// validateAntreaPolicy validates the admission of Antrea NetworkPolicy CRDs
func (v *NetworkPolicyValidator) validateAntreaPolicy(curObj, oldObj interface{}, op admv1.Operation, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// Delete of Antrea Policies have no validation. This will be an
// empty for loop.

func (v *NetworkPolicyValidator) validateAdminNetworkPolicy(curObj, oldObj interface{}, op admv1.Operation, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

func (v *antreaPolicyValidator) checkLogLabel(specAppliedTo []crdv1beta1.AppliedTo, ingress, egress []crdv1beta1.Rule) []string {
	_ = "STUB: not implemented"
	return nil
}

// validatePort validates if ports is valid
func (v *antreaPolicyValidator) validatePort(ingress, egress []crdv1beta1.Rule) error {
	_ = "STUB: not implemented"
	return nil
}

// validateAntreaGroup validates the admission of a Group, ClusterGroup resource
func (v *NetworkPolicyValidator) validateAntreaGroup(curAG, oldAG interface{}, op admv1.Operation, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// validateTier validates the admission of a Tier resource
func (v *NetworkPolicyValidator) validateTier(curTier, oldTier *crdv1beta1.Tier, op admv1.Operation, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// Tier priority updates are not allowed

func (v *antreaPolicyValidator) tierExists(name string) bool {
	_ = "STUB: not implemented"
	return false
}

// GetAdmissionResponseForErr returns an object of type AdmissionResponse with
// the submitted error message.
func GetAdmissionResponseForErr(err error) *admv1.AdmissionResponse {
	_ = "STUB: not implemented"
	return nil
}

// createValidate validates the CREATE events of Antrea-native policies,
func (v *antreaPolicyValidator) createValidate(curObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false

	// validatePolicy validates the CREATE and UPDATE events of Antrea-native policies,
}

func (v *antreaPolicyValidator) validatePolicy(curObj interface{}) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// validateRuleName validates if the name of each rule is unique within a policy
func (v *antreaPolicyValidator) validateRuleName(ingress, egress []crdv1beta1.Rule) bool {
	_ = "STUB: not implemented"
	return false
}

func (v *antreaPolicyValidator) validateAppliedTo(ingress, egress []crdv1beta1.Rule, specAppliedTo []crdv1beta1.AppliedTo) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Ensure that AppliedTo is not set in both spec and rules.

// Ensure that all rules have AppliedTo set.

// validatePeers ensures that the NetworkPolicyPeer object set in rules are valid, i.e.
// currently it ensures that a Group cannot be set with other stand-alone selectors or IPBlock.
func (v *antreaPolicyValidator) validatePeers(ingress, egress []crdv1beta1.Rule) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// validateAppliedToServiceIngressPeer ensures that if a policy or an ingress rule
// is applied to Services, the ingress rule can only use ipBlock to select workloads.
func (v *antreaPolicyValidator) validateAppliedToServiceIngressPeer(specAppliedTo []crdv1beta1.AppliedTo, ingress []crdv1beta1.Rule) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// numFieldsSetInStruct returns the number of fields in use of the object.
func numFieldsSetInStruct(obj interface{}) int { _ = "STUB: not implemented"; return 0 }

// validateIPBlock validates the CIDR values in the IPBlock.
func validateIPBlock(ipb *crdv1beta1.IPBlock) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// checkSelectorsLabels validates labels used in all selectors passed in.
func checkSelectorsLabels(selectors ...*metav1.LabelSelector) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// validateTierForPolicy validates whether a referenced Tier exists.
func (v *antreaPolicyValidator) validateTierForPolicy(tier string) (string, bool) {
	_ = "STUB: not implemented"
	// "tier" must exist before referencing
	return "", false
}

// Empty Tier name corresponds to default Tier.

// validateTierForPassAction validates that rules with pass action are not created in the Baseline Tier.
func (v *antreaPolicyValidator) validateTierForPassAction(tier string, ingress, egress []crdv1beta1.Rule) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (v *antreaPolicyValidator) validateEgressMulticastAddress(egressRule []crdv1beta1.Rule) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func validateIGMPProtocol(protocol crdv1beta1.NetworkPolicyProtocol) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (v *antreaPolicyValidator) validateMulticastIGMP(ingressRules, egressRules []crdv1beta1.Rule) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// validateL7Protocols validates the L7Protocols field set in Antrea-native policy
// rules are valid, and compatible with the ports or protocols fields.
func (v *antreaPolicyValidator) validateL7Protocols(ingressRules, egressRules []crdv1beta1.Rule) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// validateFQDNSelectors validates the toFQDN field set in Antrea-native policy egress rules are valid.
func (v *antreaPolicyValidator) validateFQDNSelectors(egressRules []crdv1beta1.Rule) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// updateValidate validates the UPDATE events of Antrea-native policies.
func (v *antreaPolicyValidator) updateValidate(curObj, oldObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false

	// deleteValidate validates the DELETE events of Antrea-native policies.
}

func (v *antreaPolicyValidator) deleteValidate(oldObj interface{}, userInfo authenticationv1.UserInfo) (string, bool) {
	_ = "STUB: not implemented"

	// createValidate validates the CREATE events of Tier resources.
	return "", false
}

func (t *tierValidator) createValidate(curObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// Tier priority must not overlap reserved tier's priority.

// Tier priority must not overlap existing tier's priority

// updateValidate validates the UPDATE events of Tier resources.
func (t *tierValidator) updateValidate(curObj, oldObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// Retrieve antrea-controller's Namespace

// Allow exception of Tier Priority updates performed by the antrea-controller

// deleteValidate validates the DELETE events of Tier resources.
func (t *tierValidator) deleteValidate(oldObj interface{}, userInfo authenticationv1.UserInfo) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Tier with existing ACNPs/ANNPs cannot be deleted.

// validateAntreaClusterGroupSpec ensures that an IPBlock is not set along with namespaceSelector, nodeSelector and/or a
// podSelector. Similarly, ExternalEntitySelector cannot be set with PodSelector.
func validateAntreaClusterGroupSpec(s crdv1beta1.GroupSpec) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// If two fields are set, only nsSel+pSel and nsSel+eeSel are valid.

func validateAntreaGroupSpec(s crdv1beta1.GroupSpec) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// If two fields are set, only nsSel+pSel and nsSel+eeSel are valid.

func validateGroupIPBlocks(ipbs []crdv1beta1.IPBlock) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// CIDR formats are already validated in validateIPBlock()

func (g *groupValidator) validateChildClusterGroup(s *crdv1beta1.ClusterGroup) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// TODO: relax this constraint when max group nesting level increases.

// the childGroup has not been created yet.

// TODO: relax this constraint when max group nesting level increases.

func (g *groupValidator) validateChildGroup(s *crdv1beta1.Group) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// TODO: relax this constraint when max group nesting level increases.

// the childGroup has not been created yet.

// TODO: relax this constraint when max group nesting level increases.

func (g *groupValidator) validateCG(cg *crdv1beta1.ClusterGroup) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (g *groupValidator) validateG(grp *crdv1beta1.Group) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// createValidate validates the CREATE events of Group, ClusterGroup resources.
func (g *groupValidator) createValidate(curObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "",

		// updateValidate validates the UPDATE events of Group, ClusterGroup resources.
		false
}

func (g *groupValidator) updateValidate(curObj, oldObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "",

		// validateGroup validates the CREATE and UPDATE events of Group, ClusterGroup resources.
		false
}

func (g *groupValidator) validateGroup(curObj interface{}) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

// deleteValidate validates the DELETE events of Group, ClusterGroup resources.
func (g *groupValidator) deleteValidate(oldObj interface{}, userInfo authenticationv1.UserInfo) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (a *adminPolicyValidator) validateAdminNP(anp *v1alpha1.AdminNetworkPolicy) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (a *adminPolicyValidator) validateBANP(banp *v1alpha1.BaselineAdminNetworkPolicy) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (a *adminPolicyValidator) createValidate(curObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

func (a *adminPolicyValidator) updateValidate(curObj, oldObj interface{}, userInfo authenticationv1.UserInfo) ([]string, string, bool) {
	_ = "STUB: not implemented"
	return nil, "", false
}

func (a *adminPolicyValidator) deleteValidate(oldObj interface{}, userInfo authenticationv1.UserInfo) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
