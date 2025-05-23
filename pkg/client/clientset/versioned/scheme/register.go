/*
Copyright 2020 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package scheme

import (
	apisixv2 "github.com/fluxcd/flagger/pkg/apis/apisix/v2"
	appmeshv1beta1 "github.com/fluxcd/flagger/pkg/apis/appmesh/v1beta1"
	appmeshv1beta2 "github.com/fluxcd/flagger/pkg/apis/appmesh/v1beta2"
	flaggerv1beta1 "github.com/fluxcd/flagger/pkg/apis/flagger/v1beta1"
	gatewayapiv1 "github.com/fluxcd/flagger/pkg/apis/gatewayapi/v1"
	gatewayapiv1beta1 "github.com/fluxcd/flagger/pkg/apis/gatewayapi/v1beta1"
	gloov1 "github.com/fluxcd/flagger/pkg/apis/gloo/v1"
	gatewayv1 "github.com/fluxcd/flagger/pkg/apis/gloogateway/v1"
	networkingv1beta1 "github.com/fluxcd/flagger/pkg/apis/istio/v1beta1"
	kedav1alpha1 "github.com/fluxcd/flagger/pkg/apis/keda/v1alpha1"
	kumav1alpha1 "github.com/fluxcd/flagger/pkg/apis/kuma/v1alpha1"
	projectcontourv1 "github.com/fluxcd/flagger/pkg/apis/projectcontour/v1"
	splitv1alpha1 "github.com/fluxcd/flagger/pkg/apis/smi/v1alpha1"
	splitv1alpha2 "github.com/fluxcd/flagger/pkg/apis/smi/v1alpha2"
	splitv1alpha3 "github.com/fluxcd/flagger/pkg/apis/smi/v1alpha3"
	traefikv1alpha1 "github.com/fluxcd/flagger/pkg/apis/traefik/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	apisixv2.AddToScheme,
	appmeshv1beta1.AddToScheme,
	appmeshv1beta2.AddToScheme,
	flaggerv1beta1.AddToScheme,
	gatewayapiv1.AddToScheme,
	gatewayapiv1beta1.AddToScheme,
	gloov1.AddToScheme,
	gatewayv1.AddToScheme,
	networkingv1beta1.AddToScheme,
	kedav1alpha1.AddToScheme,
	kumav1alpha1.AddToScheme,
	projectcontourv1.AddToScheme,
	splitv1alpha1.AddToScheme,
	splitv1alpha2.AddToScheme,
	splitv1alpha3.AddToScheme,
	traefikv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//	import (
//	  "k8s.io/client-go/kubernetes"
//	  clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//	  aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//	)
//
//	kclientset, _ := kubernetes.NewForConfig(c)
//	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
