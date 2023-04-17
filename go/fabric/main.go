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
	project "go/modules/project"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Get Stack Configs
		conf := config.New(ctx, "")
		billingSvcAcc := conf.Require("billingSA")

		// Get a access token to assume identity to a service account that has the ability
		// To provision & attach projects to a billing account.
		//
		//   - Prerequisites:  a service account in a project with projectCreator and
		//                     billingAccountUser permissions on the organization

		accessToken, err := serviceaccount.GetAccountAccessToken(ctx, &serviceaccount.GetAccountAccessTokenArgs{
			TargetServiceAccount: billingSvcAcc,
			Scopes:               []string{"cloud-platform"},
		})
		if err != nil {
			return err
		}

		// Create provider config for billing account user
		googleBillingUser, err := gcp.NewProvider(ctx, "googlebillinguser", &gcp.ProviderArgs{
			AccessToken: pulumi.String(accessToken.AccessToken),
		})

		// Create the project using our credentials from the above provider configuration
		newProject, err := project.NewProject(ctx, "sample-project", &project.ProjectArgs{}, pulumi.Provider(googleBillingUser))
		if err != nil {
			return err
		}

		return nil
	})
}
