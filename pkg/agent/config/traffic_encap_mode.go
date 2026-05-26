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

package config

type TrafficEncapModeType int

const (
	TrafficEncapModeEncap TrafficEncapModeType = iota
	TrafficEncapModeNoEncap
	TrafficEncapModeHybrid
	TrafficEncapModeNetworkPolicyOnly
	TrafficEncapModeInvalid = -1
)

var (
	modeStrs = [...]string{
		"encap",
		"noEncap",
		"hybrid",
		"networkPolicyOnly",
	}
)

// GetTrafficEncapModeFromStr returns true and TrafficEncapModeType corresponding to input string.
// Otherwise, false and undefined value is returned
func GetTrafficEncapModeFromStr(str string) (bool, TrafficEncapModeType) {
	_ = "STUB: not implemented"
	return false, *new(TrafficEncapModeType)
}

func GetTrafficEncapModes() []TrafficEncapModeType { _ = "STUB: not implemented"; return nil }

// String returns value in string.
func (m TrafficEncapModeType) String() string { _ = "STUB: not implemented"; return "" }

// IsNetworkPolicyOnly returns true if TrafficEncapModeType is network policy only.
func (m TrafficEncapModeType) IsNetworkPolicyOnly() bool { _ = "STUB: not implemented"; return false }

// SupportsNoEncap returns true if TrafficEncapModeType supports noEncap.
func (m TrafficEncapModeType) SupportsNoEncap() bool { _ = "STUB: not implemented"; return false }

// SupportsEncap returns true if TrafficEncapModeType supports encap.
func (m TrafficEncapModeType) SupportsEncap() bool { _ = "STUB: not implemented"; return false }
