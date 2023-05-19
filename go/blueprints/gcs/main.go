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
	gcs "github.com/timbohiatt/google-cloud-pulumi/go/modules/gcs"
	utils "github.com/timbohiatt/google-cloud-pulumi/go/utils"
)

// Blueprint Configuratuion
// Constants
const blueprintName string = "gcs"

// Variables
var urnPrefix string = fmt.Sprintf("blueprint-%s", blueprintName)

// Individual Blueprint Execution
func main() {

	pulumi.Run(func(ctx *pulumi.Context) (err error) {
		fmt.Println("Running Google Cloud Pulumi - Blueprint: GCS")

		// Instanciate Pulumi Provider
		var provider *gcp.Provider

		// Instanciate Pulumi Config
		conf := config.New(ctx, "")

		// Google Cloud Poject - Configuration
		// Required
		ProjectId := conf.Require("GCPProjectID")           // [Required] Google Cloud Project ID [Existing]
		BucketName := conf.Require("GCPBucketName")         // [Required] Google Cloud Storage Bucket Name
		// Optional 
		BucketPrefix := conf.Try("GCPBucketNamePrefix") // [Optional] Google Cloud Storage Bucket Name Prefix

		ExecutionServiceAccountEmail, err := conf.Try("ExecutionServiceAccountEmail") // [Optional] Pulumi Execution Service Account Email
		if err != nil {
			ExecutionServiceAccountEmail = ""
		} else {
			// Execution Service Account Email has been provided.
			provider, err = utils.GetProviderWithServiceAccount(ctx, fmt.Sprintf("%s-google-cloud-pulumi-provider", urnPrefix), ExecutionServiceAccountEmail, []string{"cloud-platform"})
			if err != nil {
				// Error Configuring Pulumi Provider to use Google Service Account
				return err
			} else {
				fmt.Println("Running Google Cloud Pulumi - Blueprint: With Google Service Account")
			}
		}

		// Run's Module: GCS
		_, err = gcs.New(ctx, fmt.Sprintf("%s", urnPrefix), &gcs.Args{
			PulumiExport: true,
			ProjectId:    ProjectId,
			Prefix:       BucketPrefix,
			Name:         BucketName,
		}, pulumi.Provider(provider))
		if err != nil {
			// Error on GCS Creation
			return err
		}
		// GCD Creation Completed
		return err
	})
}
