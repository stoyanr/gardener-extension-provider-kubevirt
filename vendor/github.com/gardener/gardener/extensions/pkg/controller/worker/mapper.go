// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package worker

import (
	"context"

	extensionshandler "github.com/gardener/gardener/extensions/pkg/handler"
	extensionspredicate "github.com/gardener/gardener/extensions/pkg/predicate"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
)

// ClusterToWorkerMapper returns a mapper that returns requests for Worker whose
// referenced clusters have been modified.
func ClusterToWorkerMapper(predicates []predicate.Predicate) handler.Mapper {
	return extensionshandler.ClusterToObjectMapper(func() runtime.Object { return &extensionsv1alpha1.WorkerList{} }, predicates)
}

// MachineSetToWorkerMapper returns a mapper that returns requests for Worker whose
// referenced MachineSets have been modified.
func MachineSetToWorkerMapper(predicates []predicate.Predicate) handler.Mapper {
	return newMachineSetToObjectMapper(func() runtime.Object { return &extensionsv1alpha1.WorkerList{} }, predicates)
}

// MachineToWorkerMapper returns a mapper that returns requests for Worker whose
// referenced Machines have been modified.
func MachineToWorkerMapper(predicates []predicate.Predicate) handler.Mapper {
	return newMachineToObjectMapper(func() runtime.Object { return &extensionsv1alpha1.WorkerList{} }, predicates)
}

type machineSetToObjectMapper struct {
	client         client.Client
	newObjListFunc func() runtime.Object
	predicates     []predicate.Predicate
}

func (m *machineSetToObjectMapper) InjectClient(c client.Client) error {
	m.client = c
	return nil
}

func (m *machineSetToObjectMapper) InjectFunc(f inject.Func) error {
	for _, p := range m.predicates {
		if err := f(p); err != nil {
			return err
		}
	}
	return nil
}

func (m *machineSetToObjectMapper) Map(obj handler.MapObject) []reconcile.Request {
	ctx := context.TODO()

	if obj.Object == nil {
		return nil
	}

	machineSet, ok := obj.Object.(*machinev1alpha1.MachineSet)
	if !ok {
		return nil
	}

	objList := m.newObjListFunc()
	if err := m.client.List(ctx, objList, client.InNamespace(machineSet.Namespace)); err != nil {
		return nil
	}

	return getReconcileRequestsFromObjectList(objList, m.predicates)
}

// newMachineSetToObjectMapper returns a mapper that returns requests for objects whose
// referenced MachineSets have been modified.
func newMachineSetToObjectMapper(newObjListFunc func() runtime.Object, predicates []predicate.Predicate) handler.Mapper {
	return &machineSetToObjectMapper{newObjListFunc: newObjListFunc, predicates: predicates}
}

type machineToObjectMapper struct {
	client         client.Client
	newObjListFunc func() runtime.Object
	predicates     []predicate.Predicate
}

func (m *machineToObjectMapper) InjectClient(c client.Client) error {
	m.client = c
	return nil
}

func (m *machineToObjectMapper) InjectFunc(f inject.Func) error {
	for _, p := range m.predicates {
		if err := f(p); err != nil {
			return err
		}
	}
	return nil
}

func (m *machineToObjectMapper) Map(obj handler.MapObject) []reconcile.Request {
	ctx := context.TODO()

	if obj.Object == nil {
		return nil
	}

	machine, ok := obj.Object.(*machinev1alpha1.Machine)
	if !ok {
		return nil
	}

	objList := m.newObjListFunc()
	if err := m.client.List(ctx, objList, client.InNamespace(machine.Namespace)); err != nil {
		return nil
	}

	return getReconcileRequestsFromObjectList(objList, m.predicates)
}

// newMachineToObjectMapper returns a mapper that returns requests for objects whose
// referenced Machines have been modified.
func newMachineToObjectMapper(newObjListFunc func() runtime.Object, predicates []predicate.Predicate) handler.Mapper {
	return &machineToObjectMapper{newObjListFunc: newObjListFunc, predicates: predicates}
}

func getReconcileRequestsFromObjectList(objList runtime.Object, predicates []predicate.Predicate) []reconcile.Request {
	var requests []reconcile.Request

	utilruntime.HandleError(meta.EachListItem(objList, func(obj runtime.Object) error {
		accessor, err := meta.Accessor(obj)
		if err != nil {
			return err
		}

		if !extensionspredicate.EvalGeneric(obj, predicates...) {
			return nil
		}

		requests = append(requests, reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: accessor.GetNamespace(),
				Name:      accessor.GetName(),
			},
		})
		return nil
	}))
	return requests
}
