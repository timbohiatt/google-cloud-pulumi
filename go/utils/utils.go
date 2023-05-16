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

package utils

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	serviceaccount "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

// GetProviderWithServiceAccount returns a Pulumi Povider;
// Returns a Puluimi Provider that is instancicated to use a Google Cloud Service Account with Specific Scopes
func GetProviderWithServiceAccount(ctx *pulumi.Context, name string, serviceAccountEmail string, scopes []string) (provider *gcp.Provider, err error) {

	if name != "" {
		// Set Default URN
		name = "google-cloud-pulumi-provider"
	}

	if len(scopes) <= 0 {
		// Set Default Scopes
		scopes = append(scopes, "cloud-platform")
	}

	// Get Service Account Access Token for Specified Service Account Email Address
	accessToken, err := serviceaccount.GetAccountAccessToken(ctx, &serviceaccount.GetAccountAccessTokenArgs{
		TargetServiceAccount: serviceAccountEmail,
		Scopes:               scopes,
	})
	if err != nil {
		// Error Getting Service Account Access Token
		return provider, err
	}

	// Create pulumi provider config for billing account user
	provider, err = gcp.NewProvider(ctx, name, &gcp.ProviderArgs{
		AccessToken: pulumi.String(accessToken.AccessToken),
	})
	if err != nil {
		// Error Creating Provider
		return provider, err
	}

	// Provider Created Successfully
	return provider, err
}
