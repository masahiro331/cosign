//
// Copyright 2021 The Sigstore Authors.
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

package oci

import v1 "github.com/google/go-containerregistry/pkg/v1"

// Signatures represents a set of signatures that are associated with a particular
// v1.Image.
type Signatures interface {
	v1.Image // The low-level representation of the signatures

	// TODO(mattmoor): Accessors that build on `v1.Image` to provide
	// higher-level accessors for the signature data that is embedded
	// in the wrapped `v1.Image`
}