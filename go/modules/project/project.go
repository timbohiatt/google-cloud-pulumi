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

	organizations "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	cloudresourcemanager "github.com/pulumi/pulumi-google-native/sdk/go/google/cloudresourcemanager/v3.Project"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceState struct {
	pulumi.ResourceState
}

/*type LoggingSink struct {
	BQPartitionedTable bool   // Optional
	Description        string // Optional
	Destination        string
	Disabled           bool              // Optional
	Exclusions         map[string]string // Optional
	Filter             string
	IAM                bool // Optional
	Type               string
	UniqueWriter       bool // Optional
}*/

/*type OrgPolicyRule struct {
	All    bool     // Optional
	Values []string // Optional
}*/

/*type OrgPolicyRules struct {
	Allow []OrgPolicyRule
	Deny  []OrgPolicyRules
}*/

/*type OrgPolicyCondition struct {
	Description string // Optional
	Expression  string // Optional
	Location    string // Optional
	Title       string // Optional
}*/

/*type OrgPolicy struct {
	InheritFromParent bool // Optional
	Reset             bool
	Rules             []OrgPolicyRules
	Enforce           bool               // Optional (For Boolean Policies Only)
	Conition          OrgPolicyCondition //Optional
}*/

/*type ServiceConfigObj struct {
	DisableOnDestroy         bool
	DisableDependentServices bool
}*/

/*type SharedVpcHostConfigObj struct {
	Enabled         bool
	ServiceProjects []string //Optional
}*/

/*type SharedVpcServiceConfigObj struct {
	HostProject        string
	ServiceIdentityIAM map[string]string //Optional
}*/

/*type EssentialContactsObj struct {
	Email       string
	LanguageTag string //Optional
}*/

type Args struct {
	PulumiExport bool
	//AutoCreateNetwork        bool
	BillingAccount string
	//Contacts                 []EssentialContactsObj
	//CustomRoles              map[string]string
	//DefaultServiceAccount    string
	DescriptiveName string
	//GroupIAM                 map[string]string
	//IAM                      map[string]string
	//IAMAdditive              map[string]string
	//IAMAdditiveMembers       map[string]string
	//Labels                   map[string]string
	//LienReason               string
	//LoggingExclusions        map[string]string
	//LoggingSinks             map[string]LoggingSink
	//MetricScopes             []string
	Name string
	//OrgPolicies              map[string]OrgPolicy
	//OrgPoliciesDataPath      string
	//OSLogin                  bool
	//OSLoginAdmins            []string
	//OSLoginUsers             []string
	Parent        string
	Prefix        string
	ProjectCreate bool
	//ServiceConfig            ServiceConfigObj
	//ServiceEncryptionKeyIds  map[string]string
	//ServicePerimeterBridges  []string
	//ServicePerimeterStandard string
	//Services                 []string
	//SharedVpcHostConfig      SharedVpcHostConfigObj
	//SharedVpcServiceConfig   SharedVpcServiceConfigObj
	//SkipDelete               bool
	//TagBindings              map[string]string
}

/*type ProjectObj struct {
	ProjectId string
	Number    string
	Name      string
}*/

type Locals struct {
	ParentId        string
	ParentType      string
	Prefix          string
	ProjectId       string
	DescriptiveName string
}

// Module Configuratuion
// Constants
const moduleName string = "project"

// Variables
var urnPrefix string = fmt.Sprintf("module-%s", moduleName)

func New(ctx *pulumi.Context, name string, args *Args, opts pulumi.ResourceOption) (state *ResourceState, err error) {
	fmt.Println("Running Google Cloud Pulumi - Module: Project")

	// Argument Validation

	// Validate - Argument - Name
	if args.Name == "" {
		// Error - Argument - Name
		err = fmt.Errorf("Unexpected Nil Value: Argument: 'Name', must not be nil. Please provide 'Name variable'", args.Name)
		return state, err
	}

	// Local Variable Configuration

	// instanciate - local variables - module
	locals := &Locals{}

	// Derive - local variable - Prefix
	if args.Prefix != "" {
		// configure user supplied prefix / else ""
		locals.Prefix = fmt.Sprintf("%s-", args.Prefix)
	}

	// Derive & Validate - local variable - ProjectId
	locals.ProjectId = fmt.Sprintf("%s%s", locals.Prefix, args.Name)
	// validate - ProjectId variable length, must be greater than 6 and less than 30 character in length
	if len(locals.ProjectId) < 6 || len(locals.ProjectId) > 30 {
		// Error Validating ProjectId format.
		err = fmt.Errorf("Format: 'ProjectId' variable, must be greater than 6 ASCII characters and less than 30 ASCII characters in length. Current value '%s' with a length of %d is invalid. ", locals.ProjectId, len(locals.ProjectId))
		return state, err
	}

	// Derive & Validate - local variables - ParentType, ParentId
	const organisations string = "organisations"
	const folders string = "folders"
	var parentOptions = []string{"organizations", "folders"}
	// validate - parent variable format, must contain "/"
	if Contains(strings.Split(args.Parent, ""), "/") {
		parentValues := strings.Split(args.Parent, "/")
		// validate - parent variable format, must contain "orgnanisations or folders"
		if Contains(parentOptions, parentValues[0]) {
			// Parent is Google Cloud Organisation or Folder
			locals.ParentType = parentValues[0] // Set Patent Type as "organizations" or "folders"
			locals.ParentId = parentValues[1]   // Set ID for Parent Organization or Folder
		} else {
			// Error Validating Parent argument format.
			err = fmt.Errorf("Format: parent variable, must be in format 'organizations/xxxxxxxx' or 'folders/xxxxxxxx'. Current value '%s' is invalid. ", args.Parent)
			return state, err
		}
	} else {
		// Error Validating Parent - missing '/' separator
		err = fmt.Errorf("Format: parent variable, must be in format 'organizations/xxxxxxxx' or 'folders/xxxxxxxx'. Current value '%s' does not contain '/' separator. ", args.Parent)
		return state, err
	}

	// Derive & Validate - local variables - DescriptiveName
	if args.DescriptiveName != "" {
		// set - provided DescriptiveName
		locals.DescriptiveName = args.DescriptiveName
	} else {
		// set - default - projectId
		locals.DescriptiveName = locals.ProjectId
	}

	// var - Google Cloud Project resource
	var gcpProject *cloudresourcemanager.Project

	// flag - bool - Create Project?
	if args.ProjectCreate {

		// instanciate - Resource Arguments - Google Cloud Project
		var projectArgs *organizations.ProjectArgs

		// If Parent is a Google Cloud Organization set the parent to this Organization
		if locals.ParentType == organisations {
			projectArgs.OrgId = pulumi.String(locals.ParentId)
		}
		// If Parent is a Google Cloud Folder set the parent to this Folder
		if locals.ParentType == folders {
			projectArgs.FolderId = pulumi.String(locals.ParentId)
		}

		projectArgs{
			ProjectId:         pulumi.String(locals.ProjectId),
			AutoCreateNetwork: pulumi.Bool(false),                 // TODO: Add Variable Support, Add Validation Routine
			BillingAccount:    pulumi.String(args.BillingAccount), // TODO: Add Variable Support, Add Validation Routine
			Name:              pulumi.String(locals.DescriptiveName),
			SkipDelete:        pulumi.Bool(false), // TODO: Add Variable Support, Add Validation Routine
			//Labels: MAP?pulumi.String()			// TODO: Add Variable Support, Add Validation Routine
		}

		// resource - [Classic] - Google Cloud Project
		gcpProject, err := organizations.NewProject(ctx, fmt.Sprintf("%s-gcp-project", urnPrefix), &projectArgs)
		if err != nil {
			// Error Creating Resource -  Google Cloud Project
			return state, err
		}
		// flag - bool - Export Resource?
		if args.PulumiExport {
			// export - resoruce - Google Cloud Project
			ctx.Export(fmt.Sprintf("%s-gcp-project", urnPrefix), gcpProject) // TODO: Fix ARN String, Use Routine
		}

	}

	return state, err
}

// Util Functions

// Contains returns a boolean value;
// Returns True when the the input array contains an element of equal value to the input string.
// Returns False when the the input array does not contain an element of equal value to the input string.
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
