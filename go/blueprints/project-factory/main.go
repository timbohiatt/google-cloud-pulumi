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

package main

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	projectfactory "github.com/timbohiatt/google-cloud-pulumi/go/factories/project-factory"
	billingBudget "github.com/timbohiatt/google-cloud-pulumi/go/modules/billing-budget"
	dns "github.com/timbohiatt/google-cloud-pulumi/go/modules/dns"
	iamServiceAccount "github.com/timbohiatt/google-cloud-pulumi/go/modules/iam-service-account"
	project "github.com/timbohiatt/google-cloud-pulumi/go/modules/project"
)

// Individual Blueprint Execution
func main() {

	pulumi.Run(func(ctx *pulumi.Context) (err error) {
		fmt.Println("Running Google Cloud Pulumi - Blueprint: Project Factory")

		var provider *gcp.Provider

		conf := config.New(ctx, "")

		// Google Cloud Poject - Configuration
		Name := conf.Require("GCPProject:Name")

		// Run's Module: Project
		_, err = projectfactory.New(ctx, "sample-project-factory", &projectfactory.Args{
			Project: project.Args{
				Name: Name,
			},
			DNS:               dns.Args{},
			BillingBudget:     billingBudget.Args{},
			IAMServiceAccount: iamServiceAccount.Args{},
		}, pulumi.Provider(provider))
		if err != nil {
			// Error on Project Creation
			return err
		}
		// Project Creation Completed
		return err
	})
}
