/*
 * Copyright 2018 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package crd

import (
	"github.com/jinzhu/copier"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ResourceChecks struct {
	Kind       string               `json:"kind,omitempty"`
	Namespace  string               `json:"namespace,omitempty"`
	Selector   metav1.LabelSelector `json:"selector,omitempty"`
	JsonPath   string               `json:"jsonpath,omitempty"`
	Pattern    string               `json:"pattern,omitempty"`
}

type RiffResources struct {
	Path       string           `json:"path,omitempty"`
	Name       string           `json:"name,omitempty"`
	Checks     []ResourceChecks `json:"checks,omitempty"`
}

type RiffSpec struct {
	Images    []string        `json:"images,omitempty"`
	Resources []RiffResources `json:"resources,omitempty"`
}

type RiffStatus struct {
	Status string `json:"status,omitempty"`
}

type RiffManifest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RiffSpec   `json:"spec,omitempty"`
	Status RiffStatus `json:"status,omitempty"`
}

func (orig *RiffManifest) DeepCopyObject() runtime.Object {
	result := &RiffManifest{}
	copier.Copy(result, orig)
	return result
}

var schemeGroupVersion = schema.GroupVersion{
	Group:    Group,
	Version:  Version,
}
