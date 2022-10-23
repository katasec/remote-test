package main

import (
	"context"
	"fmt"

	helper "github.com/katasec/pulumi-helper"
)

func main() {
	fmt.Println("hello")

	// Setup Pulumi run parameters
	args := &helper.PulumiRunRemoteParameters{
		OrgName:     "acme",
		ProjectName: "helloazure",
		StackName:   "dev",
		Destroy:     false,
		Plugins: []map[string]string{
			{
				"name":    "azure-native",
				"version": "v1.64.1",
			},
		},
		Config: []map[string]string{
			{
				"name":  "azure-native:location",
				"value": "EastAsia",
			},
		},
		GitURL:      "https://github.com/katasec/library.git",
		ProjectPath: "azure-storageaccount-sample",
	}

	// Run pulumi
	ctx := context.Background()
	helper.RunPulumiRemote(ctx, args)

}
