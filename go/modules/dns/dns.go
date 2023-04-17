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

type DNSState struct {
	pulumi.ResourceState
}

type DNSSecConfigKeySigningKeyArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	KeyLength pulumi.StringInput `pulumi:"keyLength"`
}

type DNSSecConfigZoneSigningKeyArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	KeyLength pulumi.StringInput `pulumi:"keyLength"`
}

type DNSSecConfigArgs struct {
	NonExistence   pulumi.StringInput             `pulumi:"nonExistence"`
	State          pulumi.StringInput             `pulumi:"state"`
	KeySigningKey  DNSSecConfigKeySigningKeyArgs  `pulumi:"keySigningKey"`
	ZoneSigningKey DNSSecConfigZoneSigningKeyArgs `pulumi:"zoneSigningKey"`
}

type GeoRoute struct {
	Location pulumi.StringInput   `pulumi:"location"`
	records  []pulumi.StringInput `pulumi:"records"`
}

type WrrRoute struct {
	Weight  pulumi.StringInput   `pulumi:"weight"`
	records []pulumi.StringInput `pulumi:"records"`
}

type RecordSet struct {
	TTL        pulumi.Integer       `pulumi:"ttl"`
	Records    []pulumi.StringInput `pulumi:"records"`
	GeoRouting []GeoRoute           `pulumi:"geoRouting"`
	WrrRouting []WrrRoute           `pulumi:"wrrRouting"`
}

type DNSArgs struct {
	ProjectId                 pulumi.StringInput   `pulumi:"projectId"`
	Description               pulumi.StringInput   `pulumi:"description"`
	ClientNetworks            []pulumi.StringInput `pulumi:"clientNetworks"`
	DNSSecConfig              DNSSecConfigArgs     `pulumi:"dnsSecConfig"`
	Domain                    pulumi.StringInput   `pulumi:"domain"`
	EnableLogging             pulumi.BoolInput     `pulumi:"enableLogging"`
	Forwarders                []pulumi.StringInput `pulumi:"forwarders"`
	Name                      pulumi.StringInput   `pulumi:"name"`
	PeerNetwork               pulumi.StringInput   `pulumi:"peerNetwork"`
	ServiceDirectoryNamespace pulumi.StringInput   `pulumi:"serviceDirectoryNamespace"`
	Type                      pulumi.StringInput   `pulumi:"type"`
	ZoneCreate                pulumi.BoolInput     `pulumi:"zoneCreate"`
	RecordSets                []RecordSet          `pulumi:"recordSets"`
}



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
