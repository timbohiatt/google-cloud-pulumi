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

package project

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectFactoryState struct {
	pulumi.ResourceState
}

type ProjectFactoryArgs struct {
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	Region    pulumi.StringInput `pulumi:"region"`
}

// Create a Single Project from the Project Factory
func NewProjectFactory(ctx *pulumi.Context, name string, args ProjectFactoryArgs, opts pulumi.ResourceOption) (state *ProjectFactoryState, err error) {

	// Module: Billing Alert

	// Module: DNS

	// Module: Project

	// Module: Service Account

	// Resource: Compute Subnetwork IAM Member

	return state, err
}
