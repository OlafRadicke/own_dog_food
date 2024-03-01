package main

import (
	"atlantic-ocean/gocode"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var atlanticOcean = func(ctx *pulumi.Context) error {
	nameSpaceName := "pulumi-apps"

	err := gocode.PulumiNamespace(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = gocode.LocalStorage(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	err = gocode.DebugContainer(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	// err = gocode.UptimeKuma(ctx, nameSpaceName)
	// if err != nil {
	// 	return err
	// }

	err = gocode.Hugo(ctx, nameSpaceName)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	pulumi.Run(atlanticOcean)
}
