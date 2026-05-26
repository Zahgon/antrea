// Copyright 2021 Antrea Authors
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
// Adapted from https://github.com/kubernetes/kubernetes/blob/master/test/images/agnhost/crd-conversion-webhook/converter/framework.go

package webhook

import (
	"net/http"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

// convertFunc is the user defined function for any conversion. The code in this file is a
// template that can be use for any CR conversion given this function.
type convertFunc func(Object *unstructured.Unstructured, version string) (*unstructured.Unstructured, metav1.Status)

func statusSucceed() metav1.Status { _ = "STUB: not implemented"; return *new(metav1.Status) }

// doConversionV1beta1 converts the requested objects in the v1beta1 ConversionRequest using the given conversion function and
// returns a conversion response. Failures are reported with the Reason in the conversion response.
// Deprecated: apiextensions/v1beta1 is deprecated, use apiextensions/v1 instead
func doConversionV1beta1(convertRequest *v1beta1.ConversionRequest, convert convertFunc) *v1beta1.ConversionResponse {
	_ = "STUB: not implemented"
	return nil
}

// doConversionV1 converts the requested objects in the v1 ConversionRequest using the given conversion function and
// returns a conversion response. Failures are reported with the Reason in the conversion response.
func doConversionV1(convertRequest *v1.ConversionRequest, convert convertFunc) *v1.ConversionResponse {
	_ = "STUB: not implemented"
	return nil
}

func HandleCRDConversion(crdConvertFunc convertFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// reset the request, it is not needed in a response.

// reset the request, it is not needed in a response.

type mediaType struct {
	Type, SubType string
}

var scheme = runtime.NewScheme()

func init() {
	addToScheme(scheme)
}

func addToScheme(scheme *runtime.Scheme) { _ = "STUB: not implemented"; return }

var serializers = map[mediaType]runtime.Serializer{
	{"application", "json"}: json.NewSerializerWithOptions(
		json.DefaultMetaFactory, scheme, scheme, json.SerializerOptions{
			Yaml: false, Pretty: false, Strict: false,
		}),
	{"application", "yaml"}: json.NewSerializerWithOptions(
		json.DefaultMetaFactory, scheme, scheme, json.SerializerOptions{
			Yaml: true, Pretty: false, Strict: false,
		}),
}

func getInputSerializer(contentType string) runtime.Serializer {
	_ = "STUB: not implemented"
	return *new(runtime.Serializer)
}

func getOutputSerializer(accept string) runtime.Serializer {
	_ = "STUB: not implemented"
	return *new(runtime.Serializer)
}
