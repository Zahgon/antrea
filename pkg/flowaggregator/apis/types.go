// Copyright 2024 Antrea Authors
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

package apis

// FlowRecordsResponse is the response struct of flowrecords command.
type FlowRecordsResponse map[string]interface{}

func (r FlowRecordsResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r FlowRecordsResponse) GetTableRow(maxColumnLength int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (r FlowRecordsResponse) SortRows() bool {
	_ = "STUB: not implemented"

	// RecordMetricsResponse is the response struct of recordmetrics command.
	return false
}

type RecordMetricsResponse struct {
	NumRecordsExported     int64 `json:"numRecordsExported,omitempty"`
	NumRecordsReceived     int64 `json:"numRecordsReceived,omitempty"`
	NumRecordsDropped      int64 `json:"numRecordsDropped,omitempty"`
	NumFlows               int64 `json:"numFlows,omitempty"`
	NumConnToCollector     int64 `json:"numConnToCollector,omitempty"`
	WithClickHouseExporter bool  `json:"withClickHouseExporter,omitempty"`
	WithS3Exporter         bool  `json:"withS3Exporter,omitempty"`
	WithLogExporter        bool  `json:"withLogExporter,omitempty"`
	WithIPFIXExporter      bool  `json:"withIPFIXExporter,omitempty"`
}

func (r RecordMetricsResponse) GetTableHeader() []string { _ = "STUB: not implemented"; return nil }

func (r RecordMetricsResponse) GetTableRow(maxColumnLength int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (r RecordMetricsResponse) SortRows() bool { _ = "STUB: not implemented"; return false }
