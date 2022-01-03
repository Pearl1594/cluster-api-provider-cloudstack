/*
Copyright 2022.

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

package webhook_utilities

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"reflect"
)

func EnsureFieldExists(value string, name string, allErrs field.ErrorList) field.ErrorList {
	if value == "" {
		allErrs = append(allErrs, field.Required(field.NewPath("spec", name), name))
	}
	return allErrs
}

func EnsureStringFieldsAreEqual(new string, old string, name string, allErrs field.ErrorList) field.ErrorList {
	if new != old {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", name), name))
	}
	return allErrs
}

func EnsureStringStringMapFieldsAreEqual(new *map[string]string, old *map[string]string, name string, allErrs field.ErrorList) field.ErrorList {
	if old == nil && new == nil {
		return allErrs
	}
	if new == nil || old == nil {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", name), name))
	}
	if !reflect.DeepEqual(old, new) {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec", name), name))
	}
	return allErrs
}

func AggregateObjErrors(gk schema.GroupKind, name string, allErrs field.ErrorList) error {
	if len(allErrs) == 0 {
		return nil
	}

	return errors.NewInvalid(
		gk,
		name,
		allErrs,
	)
}
