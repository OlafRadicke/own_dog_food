package goprojects

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateTheIndependentFriendDe(ctx *pulumi.Context, nameSpaceName string) error {
	_, err := helm.NewChart(ctx, "the-independent-friend-de", helm.ChartArgs{
		Namespace: pulumi.String(nameSpaceName),
		Chart:     pulumi.String("the_independent_friend_de"),
		Version:   pulumi.String("0.1.1"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://olafradicke.github.io/own_dog_food/helm/charts/"),
		},
		Values: pulumi.Map{
			"app": pulumi.Map{
				"name": pulumi.String("the-independent-friend-de"),
			},
			"image": pulumi.Map{
				"repository": pulumi.String("olafradicke/the-independent-friend-de"),
				"pullPolicy": pulumi.String("IfNotPresent"),
				"tag":        pulumi.String("4.8.3"),
			},
		},
	})
	return err
}

// func CreateTheIndependentFriendDeIngress(ctx *pulumi.Context, nameSpaceName string) error {
// 	yamlConfig := &yaml.ConfigFileArgs{
// 		File: "yaml/the-independent-friend-de/ingress.yaml",
// 	}
// 	_, err := yaml.NewConfigFile(ctx, "the-independent-friend-de-ingress", yamlConfig)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
