/*
Copyright 2017 The Kubernetes Authors.

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

package fuzzer

import (
	fuzz "github.com/google/gofuzz"
	"github.com/onexstack/onex/pkg/apis/apps"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
)

// Funcs returns the fuzzer functions for the apps api group.
var Funcs = func(codecs runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(j *apps.Chain, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again

			// match defaulting
		},
		func(j *apps.MinerSet, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again

			// match defaulting
			if j.Spec.Selector == nil {
				j.Spec.Selector = &metav1.LabelSelector{MatchLabels: j.Spec.Template.Labels}
			}
			if len(j.Labels) == 0 {
				j.Labels = j.Spec.Template.Labels
			}

		},
		func(j *apps.Miner, c fuzz.Continue) {
			c.FuzzNoCustom(s) // fuzz self without calling this function again

			// match defaulting
		},
	}
}
