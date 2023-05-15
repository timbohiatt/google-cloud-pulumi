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
	project "github.com/timbohiatt/google-cloud-pulumi/go/modules/project"
)

// Blueprint Configuratuion
// Constants
const blueprintName string = "project"

// Variables
var urnPrefix string = fmt.Sprintf("blueprint-%s", blueprintName)

// Individual Blueprint Execution
func main() {

	pulumi.Run(func(ctx *pulumi.Context) (err error) {
		fmt.Println("Running Google Cloud Pulumi - Blueprint: Project")

		var provider *gcp.Provider

		conf := config.New(ctx, "")

		// Google Cloud Poject - Configuration

		// Required
		Name := conf.Require("GCPName")                     // Google Cloud Project Name
		BillingAccount := conf.Require("GCPBillingAccount") // Google Cloud Billing Account
		Parent := conf.Require("GCPParent")                 // Google Cloud Parent Organisation or Folder

		// Optional
		Prefix, err := conf.TryString("GCPPrefix")
		if err != nil {
			Prefix = ""
		}
		DescriptiveName, err := conf.TryString("GCPDescriptiveName")
		if err != nil {
			DescriptiveName = ""
		}
		LienReason, err := conf.TryString("GCPLienReason")
		if err != nil {
			LienReason = ""
		}

		// Run's Module: Project
		_, err = project.New(ctx, fmt.Sprintf("%s", urnPrefix), &project.ProjectArgs{
			PulumiExport:      true,
			AutoCreateNetwork: false,
			BillingAccount:    BillingAccount,
			//Contacts                 []EssentialContactsObj
			//CustomRoles              map[string]string
			//DefaultServiceAccount: string,
			DescriptiveName: DescriptiveName,
			//GroupIAM                 map[string]string
			//IAM                      map[string]string
			//IAMAdditive              map[string]string
			//IAMAdditiveMembers       map[string]string
			//Labels                   map[string]string
			LienReason: LienReason,
			//LoggingExclusions        map[string]string
			//LoggingSinks             map[string]LoggingSink
			MetricScopes: []string{
				"AllEnvironments",
			},
			Name: Name,
			//OrgPolicies              map[string]OrgPolicy
			//OrgPoliciesDataPath      string
			OSLogin: true,
			//OSLoginAdmins            []string
			//OSLoginUsers             []string
			Parent: Parent,
			Prefix: Prefix,
			//Prefix                   string
			ProjectCreate: true,
			ServiceConfig: &project.ServiceConfigArgs{
				DisableOnDestroy:         false,
				DisableDependentServices: false,
			},
			//ServiceEncryptionKeyIds  map[string]string
			//ServicePerimeterBridges  []string
			//ServicePerimeterStandard string
			Services: []string{
				"storage.googleapis.com",
				"stackdriver.googleapis.com",
				"compute.googleapis.com",
			},
			//SharedVpcHostConfig      SharedVpcHostConfigObj
			//SharedVpcServiceConfig   SharedVpcServiceConfigObj
			SkipDelete: false,
			//TagBindings              map[string]string
		}, pulumi.Provider(provider))
		if err != nil {
			// Error on Project Creation
			return err
		}
		// Project Creation Completed
		return err
	})
}
