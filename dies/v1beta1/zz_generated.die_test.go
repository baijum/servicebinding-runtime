//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 the original author or authors.

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

// Code generated by diegen. DO NOT EDIT.

package v1beta1

import (
	testingx "testing"

	testing "dies.dev/testing"
)

func TestClusterWorkloadResourceMappingDie_MissingMethods(t *testingx.T) {
	die := ClusterWorkloadResourceMappingBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ClusterWorkloadResourceMappingDie: %s", diff.List())
	}
}

func TestClusterWorkloadResourceMappingSpecDie_MissingMethods(t *testingx.T) {
	die := ClusterWorkloadResourceMappingSpecBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ClusterWorkloadResourceMappingSpecDie: %s", diff.List())
	}
}

func TestClusterWorkloadResourceMappingTemplateDie_MissingMethods(t *testingx.T) {
	die := ClusterWorkloadResourceMappingTemplateBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ClusterWorkloadResourceMappingTemplateDie: %s", diff.List())
	}
}

func TestClusterWorkloadResourceMappingContainerDie_MissingMethods(t *testingx.T) {
	die := ClusterWorkloadResourceMappingContainerBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ClusterWorkloadResourceMappingContainerDie: %s", diff.List())
	}
}

func TestServiceBindingDie_MissingMethods(t *testingx.T) {
	die := ServiceBindingBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ServiceBindingDie: %s", diff.List())
	}
}

func TestServiceBindingSpecDie_MissingMethods(t *testingx.T) {
	die := ServiceBindingSpecBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ServiceBindingSpecDie: %s", diff.List())
	}
}

func TestServiceBindingWorkloadReferenceDie_MissingMethods(t *testingx.T) {
	die := ServiceBindingWorkloadReferenceBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ServiceBindingWorkloadReferenceDie: %s", diff.List())
	}
}

func TestServiceBindingServiceReferenceDie_MissingMethods(t *testingx.T) {
	die := ServiceBindingServiceReferenceBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ServiceBindingServiceReferenceDie: %s", diff.List())
	}
}

func TestEnvMappingDie_MissingMethods(t *testingx.T) {
	die := EnvMappingBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for EnvMappingDie: %s", diff.List())
	}
}

func TestServiceBindingStatusDie_MissingMethods(t *testingx.T) {
	die := ServiceBindingStatusBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ServiceBindingStatusDie: %s", diff.List())
	}
}

func TestServiceBindingSecretReferenceDie_MissingMethods(t *testingx.T) {
	die := ServiceBindingSecretReferenceBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ServiceBindingSecretReferenceDie: %s", diff.List())
	}
}
