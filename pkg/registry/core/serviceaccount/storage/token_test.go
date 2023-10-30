/*
Copyright 2014 The Kubernetes Authors.

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

package storage

import (
	"testing"

	"k8s.io/apiserver/pkg/registry/generic"
	etcd3testing "k8s.io/apiserver/pkg/storage/etcd3/testing"
	"k8s.io/kubernetes/pkg/registry/registrytest"
)

// func newToken() *TokenREST {
// 	return &TokenREST{
// 		svcaccts:             nil,
// 		pods:                 nil,
// 		secrets:              nil,
// 		issuer:               nil,
// 		auds:                 nil,
// 		audsSet:              nil,
// 		maxExpirationSeconds: 0,
// 		extendExpiration:     false,
// 	}
// }

// func TestGroupVersionKind(t *testing.T) {
// 	token := newToken()
// 	input := schema.GroupVersion{Group: "foo", Version: "bar"}
// 	gkv := token.GroupVersionKind(input)
// 	t.Logf("%v", gkv)

// }

func newToken(t *testing.T) (*TokenREST, *etcd3testing.EtcdTestServer) {
	etcdStorage, server := registrytest.NewEtcdStorage(t, "")
	restOptions := generic.RESTOptions{
		StorageConfig:           etcdStorage,
		Decorator:               generic.UndecoratedStorage,
		DeleteCollectionWorkers: 1,
		ResourcePrefix:          "serviceaccounts",
	}
	return &TokenREST{
		svcaccts: restOptions.,
		// pods:                 nil,
		// secrets:              nil,
		// issuer:               nil,
		// auds:                 nil,
		// audsSet:              nil,
		// maxExpirationSeconds: 0,
		// extendExpiration:     false,
	}, server
}
