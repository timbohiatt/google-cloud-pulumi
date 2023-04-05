package dns

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
