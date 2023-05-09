package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
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

func NewProject(ctx *pulumi.Context, name string, args *ProjectArgs, opts pulumi.ResourceOption) (state *ResourceState, err error) {
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
