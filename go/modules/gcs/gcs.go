// Copyright 2023 Google LLC
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

package gcs

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceState struct {
	pulumi.ResourceState
}

type Args struct {
	PulumiExport bool // Default False
}

type Locals struct {
}

// Module Configuratuion
// Constants
const moduleName string = "gcs"

// Variables
var urnPrefix string = fmt.Sprintf("module-%s", moduleName)

func New(ctx *pulumi.Context, name string, args *Args, opts pulumi.ResourceOption) (state *ResourceState, err error) {
	fmt.Println("Running Google Cloud Pulumi - Module: GCS")

	// Argument Validation

	// Local Variable Configuration

	// instanciate - local variables - module
	locals := &Locals{}

	return state, err
}
