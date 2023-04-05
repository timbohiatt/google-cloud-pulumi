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

package dns

import (
	"github.com/go-playground/validator/v10"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func NewDNS(ctx *pulumi.Context, name string, args DNSArgs, opts pulumi.ResourceOption) (state *DNSState, err error) {

	err := validate.Struct(args)
	validationErrors := err.(validator.ValidationErrors)
	
	// Structure Load is Completed HERE

	for key, attrs := range(args.RecordSets){
		
	}

	/*_recordsets_0 = {
		for key, attrs in var.recordsets :
		key => merge(attrs, zipmap(["type", "name"], split(" ", key)))
	}*/

	
	/* Handling TF Locals */ 




	locals {
		# split record name and type and set as keys in a map
		_recordsets_0 = {
		  for key, attrs in var.recordsets :
		  key => merge(attrs, zipmap(["type", "name"], split(" ", key)))
		}
		# compute the final resource name for the recordset
		_recordsets = {
		  for key, attrs in local._recordsets_0 :
		  key => merge(attrs, {
			resource_name = (
			  attrs.name == ""
			  ? var.domain
			  : (
				substr(attrs.name, -1, 1) == "."
				? attrs.name
				: "${attrs.name}.${var.domain}"
			  )
			)
		  })
		}
		# split recordsets between regular, geo and wrr
		geo_recordsets = {
		  for k, v in local._recordsets :
		  k => v
		  if v.geo_routing != null
		}
		regular_recordsets = {
		  for k, v in local._recordsets :
		  k => v
		  if v.records != null
		}
		wrr_recordsets = {
		  for k, v in local._recordsets :
		  k => v
		  if v.wrr_routing != null
		}
		zone = (
		  var.zone_create
		  ? try(
			google_dns_managed_zone.non-public.0, try(
			  google_dns_managed_zone.public.0, null
			)
		  )
		  : try(data.google_dns_managed_zone.public.0, null)
		)
		dns_keys = try(
		  data.google_dns_keys.dns_keys.0, null
		)
	  }




	
	/* Two Standard Routines
		 - Get Defaults
		 - Suplement Defaults (Compare incoming, add missing, exclude overidden Defaults)
		 - Struct with Perfect Values... Run Build. 


		 - Two YAMLs (Defaults & Data) ---> Merged Sruct Per Data YAML. Execute.
		 - 
	*/


	//Resource google_dns_managed_zone (Non-Public)

	//Resource google_dns_managed_zone (Public)

	//Resource google_dns_record_set (cloud static records)

	//Resource google_dns_record_set (cloud geo records)

	//Resource google_dns_record_set (cloud wrr records)

	return state, err
}
