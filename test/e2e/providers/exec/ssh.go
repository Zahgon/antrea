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

package exec

import (
	"golang.org/x/crypto/ssh"
)

// RunSSHCommand runs the provided SSH command on the specified host. Returns the exit code of the
// command, along with the contents of stdout and stderr as strings. Note that if the command
// returns a non-zero error code, this function does not report it as an error.
func RunSSHCommand(host string, config *ssh.ClientConfig, cmd string, envs map[string]string, stdin string, sudo bool) (
	code int, stdout, stderr string, err error) {
	_ = "STUB: not implemented"
	return 0, "", "", nil
}

// Set environment variables.

// Session.Setenv() requires that the remote host's sshd configuration accepts
// environment variables set by clients. So, just pre-appending environment
// variables to the command.

// SSH operation successful, but command returned error code

// command is successful
