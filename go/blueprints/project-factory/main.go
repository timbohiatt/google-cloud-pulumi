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
	projectfactory "github.com/timbohiatt/google-cloud-pulumi/go/factories/project"
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
		_, err = projectfactory.New(ctx, "sample-project-factory", &projectfactory.ProjectFactoryArgs{
			ProjectArgs: &project.ProjectArgs{
				//AutoCreateNetwork: false,
				//BillingAccount:    BillingAccount,
				//Contacts                 []EssentialContactsObj
				//CustomRoles              map[string]string
				//DefaultServiceAccount: string,
				//DescriptiveName: DescriptiveName,
				//GroupIAM                 map[string]string
				//IAM                      map[string]string
				//IAMAdditive              map[string]string
				//IAMAdditiveMembers       map[string]string
				//Labels                   map[string]string
				//LienReason               string
				//LoggingExclusions        map[string]string
				//LoggingSinks             map[string]LoggingSink
				//MetricScopes             []string
				Name: Name,
				//OrgPolicies              map[string]OrgPolicy
				//OrgPoliciesDataPath      string
				//OSLogin                  bool
				//OSLoginAdmins            []string
				//OSLoginUsers             []string
				//Parent: Parent,
				//Prefix                   string
				//ProjectCreate: true,
				//ServiceConfig            ServiceConfigObj
				//ServiceEncryptionKeyIds  map[string]string
				//ServicePerimeterBridges  []string
				//ServicePerimeterStandard string
				//Services                 []string
				//SharedVpcHostConfig      SharedVpcHostConfigObj
				//SharedVpcServiceConfig   SharedVpcServiceConfigObj
				//SkipDelete               bool
				//TagBindings              map[string]string
			},
		}, pulumi.Provider(provider))
		if err != nil {
			// Error on Project Creation
			return err
		}
		// Project Creation Completed
		return err
	})
}
