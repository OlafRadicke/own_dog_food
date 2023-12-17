package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func createOauth2Proxy(ctx *pulumi.Context, nameSpaceName string) error {
	conf := config.New(ctx, "")
	oauth2ProxyclientSecret := conf.GetSecret("oauth2ProxyclientSecret")
	// strOauth2ProxyclientSecret := fmt.Sprintf("%v", oauth2ProxyclientSecret)

	_, err := helm.NewChart(ctx, "oauth2-proxy", helm.ChartArgs{
		Namespace: pulumi.String(nameSpaceName),
		Chart:     pulumi.String("oauth2-proxy"),
		Version:   pulumi.String("6.18.0"),
		FetchArgs: helm.FetchArgs{
			Repo: pulumi.String("https://oauth2-proxy.github.io/manifests"),
		},
		Values: pulumi.Map{
			"extraArgs": pulumi.Map{
				"provider":        pulumi.String("keycloak-oidc"),
				"client-id":       pulumi.String("<your client's id>"),
				"client-secret":   oauth2ProxyclientSecret,
				"redirect-url":    pulumi.String("https://internal.yourcompany.com/oauth2/callback"),
				"oidc-issuer-url": pulumi.String("https://<keycloak host>/auth/realms/<your realm>"),
				"email-domain":    pulumi.String("<yourcompany.com>"),
				// "allowed-role":          pulumi.String("<realm role name>"),
				// "allowed-role":          pulumi.String("<client id>:<client role name>"),
				"allowed-group":         pulumi.String("</group name>"),
				"code-challenge-method": pulumi.String("S256"),
			},
		},
	})
	return err
}
