package project

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectState struct {
	pulumi.ResourceState
}

type ProjectArgs struct {
}

func NewProject(ctx *pulumi.Context, name string, args *ProjectArgs, opts pulumi.ResourceOption) (state *ProjectState, err error) {
	return state, err
}

// Individual Module Execution
func main() {

	var provider *gcp.Provider

	pulumi.Run(func(ctx *pulumi.Context) (err error) {

		// Run's Module: Project
		_, err = NewProject(ctx, "sample-project", &ProjectArgs{}, pulumi.Provider(provider))
		if err != nil {
			// Error on Project Creation
			return err
		}
		// Project Creation Completed
		return err
	})
}
