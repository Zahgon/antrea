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

package output

import (
	"bytes"
	"io"
)

const (
	maxTableOutputColumnLength int = 50
)

func TableOutput(obj interface{}, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// break after one iteration intentionally (we are just retrieving attribute
// names to use as the table header in the output)
// nolint:staticcheck

func JsonOutput(obj interface{}, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func YamlOutput(obj interface{}, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// Comment copied from: sigs.k8s.io/yaml
// We are using yaml.Unmarshal here (instead of json.Unmarshal) because the
// Go JSON library doesn't try to pick the right number type (int, float,
// etc.) when unmarshalling to interface{}, it just picks float64
// universally. go-yaml does go through the effort of picking the right
// number type, so we can preserve number type throughout this process.

// RawOutput is an output formatter whose output is similar to fmt.Print(responseString)
// to better display multiple-line string responses.
func RawOutput(obj interface{}, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// TableOutputForGetCommands formats the table output for "get" commands.
func TableOutputForGetCommands(obj interface{}, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the elements and headers of table.

func GetColumnWidths(numRows int, numCols int, rows [][]string) []int {
	_ = "STUB: not implemented"
	return nil
}

// Do not limit the column length for a single column table.
// This is for the case a single column table can have long rows which cannot
// fit into a single line (one example is the ovsflows outputs).

// Get the width of every column.

func ConstructTable(numRows int, numCols int, widths []int, rows [][]string, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// ConstructFormattedTable constructs a table with aligned column widths that displays
// all the contents. rows always includes both header and body, and is never empty.
func ConstructFormattedTable(rows [][]string, sortRows bool, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func writeSingleLine(body string, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func jsonEncode(obj interface{}, output *bytes.Buffer) error { _ = "STUB: not implemented"; return nil }

// respTransformer collects output fields in original transformedResponse
// and flattens them. respTransformer realizes this by turning obj into
// JSON and unmarshalling it.
// E.g. agent's transformedVersionResponse will only have two fields after
// transforming: agentVersion and antctlVersion.
func respTransformer(obj interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TableOutputForQueryEndpoint formats the table output for "query endpoint"
// command, utilizing constructTable to implement printing sub tables.
func TableOutputForQueryEndpoint(obj interface{}, writer io.Writer) error {
	_ = "STUB: not implemented"
	// construct sections of sub tables for responses (applied, ingressSrc, egressDst)
	return nil
}

// transform egress and ingress rules to string representation

// iterate through each endpoint and construct response

// indicate each endpoint Namespace/Name

// output applied policies to section

// output rules referencing endpoint as egress destination section

// output rules referencing endpoint as ingress source section
