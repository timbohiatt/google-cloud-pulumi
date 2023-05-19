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

package gcs

import (
	"fmt"
	"strings"

	pubsub "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
	storage "github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	utils "github.com/timbohiatt/google-cloud-pulumi/go/utils"
)

type ResourceState struct {
	pulumi.ResourceState
}

type CORSArgs struct {
	Origin         []string
	Method         []string
	ResponseHeader []string
	MaxAgeSeconds  []string
}

type LoggingConfigArgs struct {
	LogBucket       string
	LogObjectPrefix string
}

type NotificationConfigArgs struct {
	Enabled          bool
	PayloadFormat    string
	TopicName        string
	SAEmail          string
	EventTypes       []string
	CustomAttributes []string
	ObjectNamePrefix string
}

type RetentionPolicyArgs struct {
	RetentionPeriod int
	IsLocked        bool
}

type WebsiteArgs struct {
	MainPageSuffix string
	NotFoundPage   string
}

type Args struct {
	PulumiExport  bool // Default False
	CORS          *CORSArgs
	EncryptionKey string
	ForceDestroy  bool
	IAM           map[string]string
	//Labels        map[string] // TODO - Enable Labels
	Location                 string
	LoggingConfig            *LoggingConfigArgs
	Name                     string
	NotificationConfig       *NotificationConfigArgs
	Prefix                   string
	ProjectID                string
	RetentionPolicy          *RetentionPolicyArgs
	StorageClass             string
	UniformBucketLevelAccess bool
	Versioning               bool
	Website                  *WebsiteArgs
}

type Locals struct {
	Prefix                   string
	Notification             bool
	UniformBucketLevelAccess bool
}

// Module Configuratuion
// Constants
const moduleName string = "gcs"

// Variables
var urnPrefix string = fmt.Sprintf("module-%s", moduleName)

func New(ctx *pulumi.Context, name string, args *Args, opts pulumi.ResourceOption) (state *ResourceState, err error) {
	fmt.Println("Running Google Cloud Pulumi - Module: GCS")

	// Argument Validation
	// Validate - Argument - Location
	if args.Location == "" {
		args.Location = "EU" // Default Location Value to EU.
	}

	// Validate - Argument - StorageClass
	if args.StorageClass == "" {
		args.StorageClass = "MULTI_REGIONALS" // Default Value - Multi Regioanl
	}
	// Validate - Argument - StorageClass of Correct Option
	if utils.Contains([]string{"STANDARD", "MULTI_REGIONAL", "REGIONAL", "NEARLINE", "COLDLINE", "ARCHIVE"}, args.StorageClass) == false {
		// Error Validating StorageClass argument format.
		err = fmt.Errorf("Format: StorageClass variable must be one of STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.")
		return state, err
	}

	// Local Variable Configuration
	// instanciate - local variables - module
	locals := &Locals{}

	// Derive - local variable - Prefix
	if args.Prefix != "" {
		locals.Prefix = string.ToLower(fmt.Sprintf("%s-", args.Prefix))
	}

	// Derive - local variable - Notification
	if args.NotificationConfig.Enabled == true {
		locals.Notification = true
	}

	// var - Google Cloud Storage Bucket
	var gcpStorageBucket *storage.Bucket

	// instanciate - Resource Arguments - Google Cloud Project
	bucketArgs := &storage.BucketArgs{
		Name:                     pulumi.String(strings.ToLower(fmt.Sprintf("%s%s", locals.Prefix, args.Name))),
		Project:                  pulumi.String(args.ProjectID),
		Location:                 pulumi.String(args.Location),
		StorageClass:             pulumi.String(args.StorageClass),
		ForceDestroy:             pulumi.Bool(args.ForceDestroy),
		UniformBucketLevelAccess: pulumi.Bool(args.UniformBucketLevelAccess),
		Versioning: &storage.BucketVersioningArgs{
			Enabled: pulumi.Bool(args.Versioning),
		},
		Website: &storage.BucketWebsiteArgs{
			MainPageSuffix: pulumi.String(args.Website.MainPageSuffix),
			NotFoundPage:   pulumi.String(args.Website.NotFoundPage),
		},
		//Labels                     //TODO Add labels
	}

	// flag - not null - Add Encryption Key (KMS Name)
	if args.EncryptionKey != "" {
		bucketArgs.Encryption = &storage.BucketEncryptionArgs{
			bucketArgs.DefaultKmsKeyName: pulumi.String(args.EncryptionKey),
		}
	}

	// flag - not null - RetentionPolicyArgs{} - Create Retention Config
	if (args.RetentionPolicy != &RetentionPolicyArgs{}) {
		bucketArgs.BucketRetentionPolicy = &storage.BucketRetentionPolicyArgs{
			RetentionPeriod: pulumi.Int(args.RetentionPolicy.RetentionPeriod),
			IsLocked:        pulumi.Bool(args.RetentionPolicy.IsLocked),
		}
	}

	// flag - not null - LoggingConfig{} - Create Logging Config
	if (args.LoggingConfig != &LoggingConfigArgs{}) {
		bucketArgs.LoggingConfig = &storage.LoggingConfigArgs{
			LogBucket: pulumi.String(args.LoggingConfig.LogBucket),
			IsLocked:  pulumi.Bool(args.LoggingConfig.LogObjectPrefix),
		}
	}

	// TODO - CORS
	// flag - not null - CORS{} - Create CORS Config
	/*if (args.CORS != &CORSArgs{}) {
		bucketArgs.CORS.Origin = pulumi.String(args.CORS.Origin)
		bucketArgs.CORS.Origin = pulumi.String(args.CORS.Origin)
		bucketArgs.CORS.Origin = pulumi.String(args.CORS.Origin)
		bucketArgs.CORS.Origin = pulumi.String(args.CORS.Origin)
	}*/

	// TODO - LifecycleRules

	// resource - [Classic] - Google Cloud Storage Bucket
	gcpStorageBucket, err = storage.NewBucket(ctx, fmt.Sprintf("%s-gcp-storage-bucket", urnPrefix), bucketArgs)
	if err != nil {
		// Error Creating Resource -  Google Cloud Storage Bucket
		return state, err
	}

	// TODO Resource: google_storage_bucket_iam_binding
	// TODO Resource: google_storage_notification
	// TODO Resource: google_pubsub_topic_iam_binding

	// flag - not null - NotificationConfigArgs{} - Create Notification PubSub Topic
	if (args.NotificationConfig != &NotificationConfigArgs{}) {
		// resource - [Classic] - Google Cloud PubSub Topic
		var gcpPubSubTopic *pubsub.Topic
		gcpPubSubTopic, err = pubsub.NewTopic(ctx, fmt.Sprintf("%s-gcp-storage-bucket-pubsub-topic", urnPrefix), &pubsub.TopicArgs{
			Project: pulumi.String(args.ProjectID),
			Name:    pulumi.String(args.NotificationConfig.TopicName),
		})
		if err != nil {
			// Error Creating Resource - Google Cloud PubSub Topic
			return state, err
		}
		if args.PulumiExport {
			// export - resoruce -  Google Cloud PubSub Topic
			ctx.Export(fmt.Sprintf("%s-gcp-storage-bucket-pubsub-topic", urnPrefix), gcpPubSubTopic) // TODO: Fix ARN String, Use Routine
		}
	}

	return state, err
}
