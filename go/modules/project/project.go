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
	"strings"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	compute "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	monitoring "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/monitoring"
	service "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
	cloudresourcemanager "github.com/pulumi/pulumi-google-native/sdk/go/google/cloudresourcemanager/v3"
	contacts "github.com/pulumi/pulumi-google-native/sdk/go/google/essentialcontacts/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type ProjectState struct {
	pulumi.ResourceState
}

type LoggingSink struct {
	BQPartitionedTable bool   // Optional
	Description        string // Optional
	Destination        string
	Disabled           bool              // Optional
	Exclusions         map[string]string // Optional
	Filter             string
	IAM                bool // Optional
	Type               string
	UniqueWriter       bool // Optional
}

type OrgPolicyRule struct {
	All    bool     // Optional
	Values []string // Optional
}

type OrgPolicyRules struct {
	Allow []OrgPolicyRule
	Deny  []OrgPolicyRules
}

type OrgPolicyCondition struct {
	Description string // Optional
	Expression  string // Optional
	Location    string // Optional
	Title       string // Optional
}

type OrgPolicy struct {
	InheritFromParent bool // Optional
	Reset             bool
	Rules             []OrgPolicyRules
	Enforce           bool               // Optional (For Boolean Policies Only)
	Conition          OrgPolicyCondition //Optional
}

type ServiceConfigObj struct {
	DisableOnDestroy         bool
	DisableDependentServices bool
}

type SharedVpcHostConfigObj struct {
	Enabled         bool
	ServiceProjects []string //Optional
}

type SharedVpcServiceConfigObj struct {
	HostProject        string
	ServiceIdentityIAM map[string]string //Optional
}

type EssentialContactsObj struct {
	Email       string
	LanguageTag string //Optional
}

type ProjectArgs struct {
	AutoCreateNetwork        bool
	BillingAccount           string
	Contacts                 []EssentialContactsObj
	CustomRoles              map[string]string
	DefaultServiceAccount    string
	DescriptiveName          string
	GroupIAM                 map[string]string
	IAM                      map[string]string
	IAMAdditive              map[string]string
	IAMAdditiveMembers       map[string]string
	Labels                   map[string]string
	LienReason               string
	LoggingExclusions        map[string]string
	LoggingSinks             map[string]LoggingSink
	MetricScopes             []string
	Name                     string
	OrgPolicies              map[string]OrgPolicy
	OrgPoliciesDataPath      string
	OSLogin                  bool
	OSLoginAdmins            []string
	OSLoginUsers             []string
	Parent                   string
	Prefix                   string
	ProjectCreate            bool
	ServiceConfig            ServiceConfigObj
	ServiceEncryptionKeyIds  map[string]string
	ServicePerimeterBridges  []string
	ServicePerimeterStandard string
	Services                 []string
	SharedVpcHostConfig      SharedVpcHostConfigObj
	SharedVpcServiceConfig   SharedVpcServiceConfigObj
	SkipDelete               bool
	TagBindings              map[string]string
}

type ProjectObj struct {
	ProjectId string
	Number    string
	Name      string
}

// Local Variables
var DescriptiveName string
var ParentType string
var ParentId string
var Prefix string
var Project ProjectObj

func NewProject(ctx *pulumi.Context, name string, args *ProjectArgs, opts pulumi.ResourceOption) (state *ProjectState, err error) {

	// Register Component Resource
	googleCloudProject := &ProjectState{}
	err = ctx.RegisterComponentResource("pkg:google-cloud-pulumi:go:module:project", name, googleCloudProject, opts)
	if err != nil {
		return nil, err
	}

	/*
		=========================================================================
		==		Project Module Variables & Validations
		=========================================================================
	*/

	// Calculate Prefix
	if args.Prefix != "" {
		Prefix = fmt.Sprintf("%s-", args.Prefix)
	}

	// Calculate Parent Type & Parent ID [ Format: {ORGANIZATION|FOLDER}/{ID} ]
	if args.Parent != "" {
		ParentType = strings.ToLower(strings.Split(args.Parent, "/")[0]) // ORGANIZATION||FOLDER
		ParentId = strings.ToLower(strings.Split(args.Parent, "/")[1])   // ID

		if ParentType != "organizations" || ParentType != "folders" {
			// Invalid Parent Type; Must be 'organizations' or 'folders'
			ctx.Log.Error("Project Creation Error; Parent Type must be 'organizations' or 'folders'.", nil)
			return state, err
		}
	}

	// Construct Project Descriptive Name
	if args.DescriptiveName != "" {
		DescriptiveName = DescriptiveName
	} else {
		DescriptiveName = fmt.Sprintf("%s%s", Prefix, args.Name)
	}

	/*
		=========================================================================
		==		Project Module Resources
		=========================================================================
	*/

	// Instanciate a New Google Cloud Project Object
	var newProject *cloudresourcemanager.Project
	// [Pulumi Native] - Google Cloud Project
	if args.ProjectCreate {
		newProject, err = cloudresourcemanager.NewProject(ctx, fmt.Sprintf("project"), &cloudresourcemanager.ProjectArgs{
			Parent:      pulumi.String(fmt.Sprintf("%s/%s", ParentType, ParentId)), // Organization or Folder to Create Project
			ProjectId:   pulumi.String(fmt.Sprintf("%s%s", Prefix, args.Name)),     // Google Project ID
			DisplayName: pulumi.String(DescriptiveName),                            // Google Project Descriptive Name
			//Labels // TODO

		})
		if err != nil {
			// Error Creating New Google Cloud Project
			return state, err
		}
	}

	/*
		TODO Notes:
			- Needs Billing Account Linking: (Not Supported With Native?)
			- Needs AutoCreateNetworks FALSE: (Not Supported With Native?)
	*/

	// [Pulumi Classic] - Google Project Service API
	if args.ProjectCreate {
		for serviceIdx, serviceName := range args.Services {
			_, err := service.NewService(ctx, fmt.Sprintf("project-%d-%s", serviceIdx, serviceName), &service.ServiceArgs{
				Project:                  newProject.ProjectId,
				DisableDependentServices: pulumi.Bool(true),
				DisableOnDestroy:         pulumi.Bool(true),
				Service:                  pulumi.String(serviceName),
			})
			if err != nil {
				// Error Enabling Project Service API
				return state, err
			}
		}
	}

	// [Pulumi Classic] - Google Project Metadata Item (Enable: OS Login)
	if args.OSLogin {
		_, err := compute.NewProjectMetadata(ctx, "metadata-enable-oslogin", &compute.ProjectMetadataArgs{
			Metadata: pulumi.StringMap{
				"enable-oslogin": pulumi.String("TRUE"),
			},
		})
		if err != nil {
			// Error Creating Project Metadata & Enabling OS Login
			return state, err
		}
	}

	// [Pulumi Native] - Google Cloud Resource Lien
	//if args.LienReason != "" {
	//	lien, err := cloudresourcemanager.NewLien(ctx, "lien", &cloudresourcemanager.LienArgs{
	//Parent: // PROJECT NUMBER "projects/XXXXXX"
	//Name: // TODO
	//Reason: //VAR TODO
	//Origin: "created-by-pulumi"
	//CreateTime: // TODO
	//Restrictions: // TODO ["resourcemanager.projects.delete"]
	//	})
	//	if err != nil {
	// Error Creating Lien
	//		return state, err
	//	}
	//}

	// [Pulumi Native] - Google Essential Contacts (Project)
	for contactIdx, contactDetails := range args.Contacts {
		_, err := contacts.NewContact(ctx, fmt.Sprintf("essential-contact-project-%d", contactIdx), &contacts.ContactArgs{
			Project:     newProject.ProjectId,
			Email:       pulumi.String(contactDetails.Email),
			LanguageTag: pulumi.String(contactDetails.LanguageTag),
			//NotificationCategorySubscriptions: &contacts.ContactArgs{}
		})
		if err != nil {
			// Error Creating Project Essential Contacts
			return state, err
		}
	}

	// [Pulumi Classic] - Google Monitored Project (Metric Scope)
	for metricScopeIdx, metricScopeData := range args.MetricScopes {
		_, err := monitoring.NewMonitoredProject(ctx, fmt.Sprintf("metric-scope-%d", metricScopeIdx), &monitoring.MonitoredProjectArgs{
			Name:         newProject.ProjectId,
			MetricsScope: pulumi.String(metricScopeData),
		})
		if err != nil {
			return state, err
		}
	}

	// Resource - Google Tags Tag Bindings (Binding)

	return state, err
}

// Individual Module Execution
func main() {
	pulumi.Run(func(ctx *pulumi.Context) (err error) {

		conf := config.New(ctx, "")

		ExecutionServiceAccount := conf.Require("ExecutionServiceAccount")
		BillingAccount := conf.Require("BillingAccount")
		// Google Cloud Poject - Configuration
		Name := conf.Require("GCPProject:Name")
		DescriptiveName := conf.Require("GCPProject:DescriptiveName")
		// Folder or Organization in which to deploy
		Parent := conf.Require("Parent")

		var provider *gcp.Provider
		if ExecutionServiceAccount != "" {
			accessToken, err := serviceaccount.GetAccountAccessToken(ctx, &serviceaccount.GetAccountAccessTokenArgs{
				TargetServiceAccount: ExecutionServiceAccount,
				Scopes:               []string{"cloud-platform"},
			})
			if err != nil {
				return err
			}
			// Create provider config for billing account user
			provider, err = gcp.NewProvider(ctx, "executionServiceAccountUser", &gcp.ProviderArgs{
				AccessToken: pulumi.String(accessToken.AccessToken),
			})
			if err != nil {
				return err
			}
		}

		// Run's Module: Project
		_, err = NewProject(ctx, "sample-project", &ProjectArgs{
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
			Parent: Parent,
			//Prefix                   string
			ProjectCreate: true,
			//ServiceConfig            ServiceConfigObj
			//ServiceEncryptionKeyIds  map[string]string
			//ServicePerimeterBridges  []string
			//ServicePerimeterStandard string
			//Services                 []string
			//SharedVpcHostConfig      SharedVpcHostConfigObj
			//SharedVpcServiceConfig   SharedVpcServiceConfigObj
			//SkipDelete               bool
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
