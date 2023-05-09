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
	project "github.com/timbohiatt/google-cloud-pulumi/modules/go/project"
)

type ResourceState struct {
	pulumi.ResourceState
}

type ProjectFactoryArgs struct {
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	Region    pulumi.StringInput `pulumi:"region"`
}

// Create a Single Project from the Project Factory
func NewProjectFactory(ctx *pulumi.Context, name string, args ProjectFactoryArgs, opts pulumi.ResourceOption) (state *ResourceState, err error) {

	// Module: Billing Alert

	// Module: DNS

	// Module: Project
	factoryProject, err := project.NewProject(ctx*pulumi.Context, "project-factory-project", project.ProjectArgs{
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
		//Name: Name,
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
	}, pulumi.ResourceOption)
	if err != nil {
		// Error Creating Project
		return state, err
	}

	// Module: Service Account

	// Resource: Compute Subnetwork IAM Member

	return state, err
}
