package PulumiGoCode

import (
	// "azure/PulumiGoCode"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func StaticWebsite(ctx *pulumi.Context, resourceGroup *resources.ResourceGroup) error {

	skuArgs := storage.SkuArgs{
		Name: pulumi.String("Standard_LRS"),
	}
	storageAccountArgs := storage.StorageAccountArgs{
		ResourceGroupName: resourceGroup.Name,
		Sku:               &skuArgs,
		Kind:              pulumi.String("StorageV2"),
	}

	// Create an Azure resource (Storage Account)
	account, err := storage.NewStorageAccount(ctx, "sa001", &storageAccountArgs)
	if err != nil {
		return err
	}

	// Export the primary key of the Storage Account

	getAccountKey := func(args []interface{}) (string, error) {
		resourceGroupName := args[0].(string)
		accountName := args[1].(string)
		listStorageAccountKeysArgs := &storage.ListStorageAccountKeysArgs{
			ResourceGroupName: resourceGroupName,
			AccountName:       accountName,
		}
		accountKeys, err := storage.ListStorageAccountKeys(
			ctx,
			listStorageAccountKeysArgs,
		)
		if err != nil {
			return "", err
		}
		return accountKeys.Keys[0].Value, nil
	}

	ctx.Export(
		"primaryStorageKey",
		pulumi.All(resourceGroup.Name, account.Name).ApplyT(getAccountKey),
	)

	// Enable static website support
	storageAccountStaticWebsiteArgs := &storage.StorageAccountStaticWebsiteArgs{
		AccountName:       account.Name,
		ResourceGroupName: resourceGroup.Name,
		IndexDocument:     pulumi.String("index.html"),
	}
	staticWebsite, err := storage.NewStorageAccountStaticWebsite(
		ctx, "staticWebsite",
		storageAccountStaticWebsiteArgs,
	)
	if err != nil {
		return err
	}

	// Upload the file
	_, err = storage.NewBlob(ctx, "index.html", &storage.BlobArgs{
		ResourceGroupName: resourceGroup.Name,
		AccountName:       account.Name,
		ContainerName:     staticWebsite.ContainerName,
		Source:            pulumi.NewFileAsset("static_content/index.html"),
		ContentType:       pulumi.String("text/html"),
	})
	if err != nil {
		return err
	}

	// Web endpoint to the website
	ctx.Export("staticEndpoint", account.PrimaryEndpoints.Web())

	// END
	return nil
}
