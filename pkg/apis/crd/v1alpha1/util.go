// Copyright 2025 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/conversion"
)

// ConditionEqualsIgnoreLastTransitionTime checks equality of two conditions, ignoring the
// LastTransitionTime field. It must be exported because it is used directly by the
// packetcapture controller's mergeConditions logic.
func ConditionEqualsIgnoreLastTransitionTime(a, b PacketCaptureCondition) bool {
	_ = "STUB: not implemented"
	return false
}

func ConditionSliceEqualsIgnoreLastTransitionTime(as, bs []PacketCaptureCondition) bool {
	_ = "STUB: not implemented"
	return false
}

// Flow exporter protocol name constants
const (
	FlowExporterProtocolGRPC  = "grpc"
	FlowExporterProtocolIPFIX = "ipfix"
)

var semanticIgnoreLastTransitionTime = conversion.EqualitiesOrDie(
	ConditionSliceEqualsIgnoreLastTransitionTime,
)

// PacketCaptureStatusEqual performs a semantic deep equality check between two
// PacketCaptureStatus objects.
func PacketCaptureStatusEqual(oldStatus, newStatus PacketCaptureStatus) bool {
	_ = "STUB: not implemented"
	return false
}

func (proto *FlowExporterGRPCConfig) Name() string { _ = "STUB: not implemented"; return "" }

func (proto *FlowExporterIPFIXConfig) Name() string { _ = "STUB: not implemented"; return "" }

func (proto *FlowExporterGRPCConfig) TransportProtocol() FlowExporterTransportProtocol {
	_ = "STUB: not implemented"
	return *new(FlowExporterTransportProtocol)
}

func (proto *FlowExporterIPFIXConfig) TransportProtocol() FlowExporterTransportProtocol {
	_ = "STUB: not implemented"
	return *new(FlowExporterTransportProtocol)
}
