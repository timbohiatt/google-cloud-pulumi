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
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	billingBudget "github.com/timbohiatt/google-cloud-pulumi/go/modules/billing-budget"
	dns "github.com/timbohiatt/google-cloud-pulumi/go/modules/dns"
	iamServiceAccount "github.com/timbohiatt/google-cloud-pulumi/go/modules/iam-service-account"
	project "github.com/timbohiatt/google-cloud-pulumi/go/modules/project"
)

type ResourceState struct {
	pulumi.ResourceState
}

type Args struct {
	Project           *project.Args
	DNS               *dns.Args
	BillingBudget     *billingBudget.Args
	IAMServiceAccount *iamServiceAccount.Args
}

// Create a Single Project from the Project Factory
func New(ctx *pulumi.Context, name string, args Args, opts pulumi.ResourceOption) (state *ResourceState, err error) {
	fmt.Println("Running Google Cloud Pulumi - Factory: Project")

	var provider *gcp.Provider

	// Module: Billing Alert
	_, err = billingBudget.New(ctx, "project-factory-billing-alert", &billingBudget.Args{}, pulumi.Provider(provider))
	if err != nil {
		// Error Creating Billing Alert
		return state, err
	}

	// Module: DNS
	_, err = dns.New(ctx, "project-factory-dns", &args.DNS, pulumi.Provider(provider))
	if err != nil {
		// Error Creating DNS
		return state, err
	}

	// Module: Project
	_, err = project.New(ctx, "project-factory-project", &args.Project, pulumi.Provider(provider))
	if err != nil {
		// Error Creating Project
		return state, err
	}

	// Module: Service Account
	_, err = iamServiceAccount.New(ctx, "project-factory-iam-service-account", &args.IAMServiceAccount, pulumi.Provider(provider))
	if err != nil {
		// Error Creating Service Account
		return state, err
	}

	// Resource: Compute Subnetwork IAM Member
	// TODO

	return state, err
}
