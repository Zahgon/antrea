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
	"reflect"

	"github.com/spf13/cobra"
	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type formatterType string

const (
	jsonFormatter  formatterType = "json"
	yamlFormatter  formatterType = "yaml"
	tableFormatter formatterType = "table"
	rawFormatter   formatterType = "raw"
)

// commandGroup is used to group commands, it could be specified in commandDefinition.
// The default commandGroup of a commandDefinition is `flat` which means the command
// is a direct sub-command of the root command. For any other commandGroup, the
// antctl framework will generate a same name sub-command of the root command for
// each of them, any commands specified as one of these group will need to be invoked
// as:
//
//	antctl <commandGroup> <command>
type commandGroup uint
type OutputType uint

// There are two output types: single item or list and the actual type is decided by
// OutputType value here and command's arguments.
const (
	// defaultType represents the output type is single item if there is an argument
	// and its value is provided. If not, the output type is list.
	defaultType OutputType = iota
	// single represents the output type is always single item.
	single
	// multiple represents the output type is always list.
	multiple
)

const (
	flat commandGroup = iota
	get
	query
	mc
	upgrade
	check
)

var groupCommands = map[commandGroup]*cobra.Command{
	get: {
		Use:   "get",
		Short: "Get the status or resource of a topic",
		Long:  "Get the status or resource of a topic",
	},
	query: {
		Use:   "query",
		Short: "Execute a user-provided query",
		Long:  "Execute a user-provided query",
	},
	mc: {
		Use:   "mc",
		Short: "Sub-commands of multi-cluster feature",
		Long:  "Sub-commands of multi-cluster feature",
	},
	upgrade: {
		Use:   "upgrade",
		Short: "Sub-commands for upgrade operations",
		Long:  "Sub-commands for upgrade operations",
	},
	check: {
		Use:   "check",
		Short: "Performs pre and post installation checks",
	},
}

type endpointResponder interface {
	OutputType() OutputType
	flags() []flagInfo
}

type resourceEndpoint struct {
	groupVersionResource *schema.GroupVersionResource
	resourceName         string
	namespaced           bool
	supportSorting       bool
	params               []flagInfo
	parameterTransform   func(args map[string]string) (k8sruntime.Object, error)
	restMethod           restMethod
}

func (e *resourceEndpoint) OutputType() OutputType {
	_ = "STUB: not implemented"
	return *new(OutputType)
}

func (e *resourceEndpoint) flags() []flagInfo { _ = "STUB: not implemented"; return nil }

func getSortByFlag() flagInfo { _ = "STUB: not implemented"; return *new(flagInfo) }

type restMethod uint

const (
	restGet restMethod = iota
	restPost
)

type nonResourceEndpoint struct {
	path       string
	params     []flagInfo
	outputType OutputType
}

func (e *nonResourceEndpoint) flags() []flagInfo { _ = "STUB: not implemented"; return nil }

func (e *nonResourceEndpoint) OutputType() OutputType {
	_ = "STUB: not implemented"
	return *

	// endpoint is used to specified the API for an antctl running against antrea-controller.
	new(OutputType)
}

type endpoint struct {
	resourceEndpoint    *resourceEndpoint
	nonResourceEndpoint *nonResourceEndpoint
	// addonTransform is used to transform or update the response data received
	// from the handler, it must returns an interface which has same type as
	// TransformedResponse.
	addonTransform func(reader io.Reader, single bool, opts map[string]string) (interface{}, error)
	// requestErrorFallback is called when a client request fails, in which
	// case transforms are called on the io.Reader object returned by this
	// function. This is useful if a command still needs to output useful
	// information in case of error.
	requestErrorFallback func() (io.Reader, error)
}

// flagInfo represents a command-line flag that can be provided when invoking an antctl command.
type flagInfo struct {
	name            string
	shorthand       string
	defaultValue    string
	supportedValues []string
	arg             bool
	usage           string
	isBool          bool
}

// rawCommand defines a full function cobra.Command which lets developers
// write complex client-side tasks. Only the global flags of the antctl framework will
// be passed to the cobra.Command.
type rawCommand struct {
	cobraCommand          *cobra.Command
	supportAgent          bool
	supportController     bool
	supportFlowAggregator bool
	commandGroup          commandGroup
}

// commandDefinition defines options to create a cobra.Command for an antctl client.
type commandDefinition struct {
	// Cobra related
	use     string
	aliases []string
	short   string
	long    string
	example string // It will be filled with generated examples if it is not provided.
	// commandGroup represents the group of the command.
	commandGroup           commandGroup
	controllerEndpoint     *endpoint
	agentEndpoint          *endpoint
	flowAggregatorEndpoint *endpoint
	// transformedResponse is the final response struct of the command. If the
	// AddonTransform is set, TransformedResponse is not needed to be used as the
	// response struct of the handler, but it is still needed to guide the formatter.
	// It should always be filled.
	transformedResponse reflect.Type
}

func (cd *commandDefinition) namespaced() bool { _ = "STUB: not implemented"; return false }

func (cd *commandDefinition) getAddonTransform() func(reader io.Reader, single bool, opts map[string]string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil
}

func (cd *commandDefinition) getEndpoint() endpointResponder {
	_ = "STUB: not implemented"
	return *new(endpointResponder)
}

func (cd *commandDefinition) getRequestErrorFallback() func() (io.Reader, error) {
	_ = "STUB: not implemented"
	return nil
}

// applySubCommandToRoot applies the commandDefinition to a cobra.Command with
// the client. It populates basic fields of a cobra.Command and creates the
// appropriate RunE function for it according to the commandDefinition.
func (cd *commandDefinition) applySubCommandToRoot(root *cobra.Command, client AntctlClient, out io.Writer) {
	_ = "STUB: not implemented"
	return
}

// when antctl runs outside the Controller/Agent/FlowAggregator Pod. This check ensures that
// the log-level command is not added to the list of available commands.

// validate checks if the commandDefinition is valid.
func (cd *commandDefinition) validate() []error { _ = "STUB: not implemented"; return nil }

// decode parses the data in reader and converts it to one or more
// TransformedResponse objects. If single is false, the return type is
// []TransformedResponse. Otherwise, the return type is TransformedResponse.
func (cd *commandDefinition) decode(r io.Reader, single bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// output reads bytes from the resp and outputs the data to the writer in desired
// format. If the AddonTransform is set, it will use the function to transform
// the data first. It will try to output the resp in the format ft specified after
// doing transform.
func (cd *commandDefinition) output(resp io.Reader, writer io.Writer, ft formatterType, single bool, args map[string]string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Decode the data if there is no AddonTransform.

// No response returned.

// If the transformed response is of type []byte, just output
// the raw bytes.

// Output structure data in format

func (cd *commandDefinition) collectFlags(cmd *cobra.Command, args []string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cd *commandDefinition) validateFlagValue(val string, supportedValues []string) bool {
	_ = "STUB: not implemented"
	return false
}

// newCommandRunE creates the RunE function for the command. The RunE function
// checks the args according to argOption and flags.
func (cd *commandDefinition) newCommandRunE(c AntctlClient, out io.Writer) func(*cobra.Command, []string) error {
	_ = "STUB: not implemented"
	return nil
}

// applyFlagsToCommand sets up args and flags for the command.
func (cd *commandDefinition) applyFlagsToCommand(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	return
}

// When the flag is a boolean, the default value will always be false.

// applyExampleToCommand generates examples according to the commandDefinition.
// It only creates for commands which specified TransformedResponse. If the singleObject
// is specified, it only creates one example to retrieve the single object. Otherwise,
// it will generates examples about retrieving single object according to the key
// argOption and retrieving the object list.
func (cd *commandDefinition) applyExampleToCommand(cmd *cobra.Command) {
	_ = "STUB: not implemented"
	return
}
