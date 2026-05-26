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

package antctl

import (
	"io"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

// commandList organizes commands definitions.
// It is the protocol for a pair of antctl client and server.
type commandList struct {
	definitions []commandDefinition
	rawCommands []rawCommand
	codec       serializer.CodecFactory
}

func (cl *commandList) applyPersistentFlagsToRoot(root *cobra.Command) {
	_ = "STUB: not implemented"
	return
}

// applyToRootCommand is the "internal" version of ApplyToRootCommand, used for testing
func (cl *commandList) applyToRootCommand(root *cobra.Command, client AntctlClient, out io.Writer) {
	_ = "STUB: not implemented"
	return
}

// ApplyToRootCommand applies the commandList to the root cobra command, it applies
// each commandDefinition of it to the root command as a sub-command.
func (cl *commandList) ApplyToRootCommand(root *cobra.Command) { _ = "STUB: not implemented"; return }

// validate checks the validation of the commandList.
func (cl *commandList) validate() []error { _ = "STUB: not implemented"; return nil }

// GetDebugCommands returns all commands supported by Controller, Agent or Flow Aggregator
// that are used for debugging purpose.
func (cl *commandList) GetDebugCommands(mode string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// TODO: incorporate query commands into e2e testing once proxy access is implemented

// log-level command does not support remote execution.

// proxy will keep running until interrupted so it
// cannot be used as is in e2e tests. For packetcapture, the default values didn't
// make much sense in e2e tests.

// renderDescription replaces placeholders ${component} in Short and Long of a command
// to the determined component during runtime.
func renderDescription(command *cobra.Command) { _ = "STUB: not implemented"; return }
