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

package gkecluster

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContainerClusterState struct {
	pulumi.ResourceState
}

type ContainerClusterArgs struct {
	ProjectId   pulumi.StringInput `pulumi:"projectId"`
	Location    pulumi.StringInput `pulumi:"location"`
	Name        pulumi.StringInput `pulumi:"name"`
	Description pulumi.StringInput `pulumi:"description"`
	AutoPilot   pulumi.Bool
	NetConfig   ContainerClusterNetworkConfig
}

type ContainerClusterNetworkConfig struct {
	Network    pulumi.StringInput `pulumi:"network"`
	SubNetwork pulumi.StringInput `pulumi:"subnetwork"`
}

type ContainerNodePoolArgs struct {
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	Location  pulumi.StringInput `pulumi:"location"`
	Cluster   pulumi.StringInput `pulumi:"cluster"`
}

func NewContainerCluster(ctx *pulumi.Context, name string, args ContainerClusterArgs, opts pulumi.ResourceOption) (*ContainerClusterState, error) {
	containerCluster := &ContainerClusterState{}
	err := ctx.RegisterComponentResource("pkg:google:gke-cluster", name, containerCluster, opts)
	if err != nil {
		return nil, err
	}

	svcAcc, err := serviceaccount.NewAccount(ctx, name, &serviceaccount.AccountArgs{
		Project:     args.ProjectId,
		AccountId:   pulumi.String(fmt.Sprintf("svc-%v", name)),
		DisplayName: pulumi.String(name),
	})
	if err != nil {
		return nil, err
	}

	cluster, err := container.NewCluster(ctx, name, &container.ClusterArgs{
		Project:               args.ProjectId,
		Location:              args.Location,
		RemoveDefaultNodePool: pulumi.Bool(true),
		InitialNodeCount:      pulumi.Int(1),
		Network:               args.NetConfig.Network,
		Subnetwork:            args.NetConfig.SubNetwork,
	})
	if err != nil {
		return nil, err
	}
	_, err = container.NewNodePool(ctx, name, &container.NodePoolArgs{
		Project:   args.ProjectId,
		Location:  args.Location,
		Cluster:   cluster.Name,
		NodeCount: pulumi.Int(1),
		NodeConfig: &container.NodePoolNodeConfigArgs{
			Preemptible:    pulumi.Bool(true),
			MachineType:    pulumi.String("e2-medium"),
			ServiceAccount: svcAcc.Email,
			OauthScopes: pulumi.StringArray{
				pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return containerCluster, nil
}

func NewContainerNodePool(ctx *pulumi.Context, name string, args ContainerNodePoolArgs, opts pulumi.ResourceOption) (*ContainerClusterState, error) {
	containerNodePool := &ContainerClusterState{}
	err := ctx.RegisterComponentResource("pkg:google:gke-cluster", name, containerNodePool, opts)
	if err != nil {
		return nil, err
	}

	_, err = container.NewNodePool(ctx, name, &container.NodePoolArgs{
		Project:  args.ProjectId,
		Location: args.Location,
	})
	if err != nil {
		return nil, err
	}
	return containerNodePool, nil
}
