// Copyright 2022 Antrea Authors
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

type IPsecAuthenticationMode int

const (
	IPsecAuthenticationModePSK IPsecAuthenticationMode = iota
	IPsecAuthenticationModeCert
	IPsecAuthenticationModeInvalid = -1
)

var supportedIPsecAuthenticationModeStrs = [...]string{
	"psk",
	"cert",
}

func GetIPsecConfigModes() []IPsecAuthenticationMode { _ = "STUB: not implemented"; return nil }

// String returns value in string.
func (am IPsecAuthenticationMode) String() string { _ = "STUB: not implemented"; return "" }

// GetIPsecAuthenticationModeFromStr returns true and IPsecAuthenticationModeType corresponding to input string.
// Otherwise, false and undefined value is returned
func GetIPsecAuthenticationModeFromStr(str string) (bool, IPsecAuthenticationMode) {
	_ = "STUB: not implemented"
	return false, *new(IPsecAuthenticationMode)
}
