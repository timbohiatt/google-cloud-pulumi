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

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceState struct {
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

type Args struct {
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

func New(ctx *pulumi.Context, name string, args *Args, opts pulumi.ResourceOption) (state *ResourceState, err error) {
	fmt.Println("Running Google Cloud Pulumi - Module: Project")
	return state, err
}
