package azure

import (
	"context"
	"fmt"
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/kvendingoldo/cloud-secrets/provider"
	"os"
)

type AzureProvider struct {
	provider.BaseProvider
	client   keyvault.BaseClient
	vaultURL string
}

type AzureConfig struct {
	Region        string
	ResourceGroup string
	KeyVault      string
}

func NewAzureProvider(azureConfig AzureConfig) (*AzureProvider, error) {

	provider := &AzureProvider{
		client:   keyvault.New(),
		vaultURL: "https://" + azureConfig.KeyVault + ".vault.azure.net",
	}

	authorizer, err := kvauth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Printf("unable to create vault authorizer: %v\n", err)
		os.Exit(1)
	}

	provider.client.Authorizer = authorizer

	return provider, nil
}

func (p *AzureProvider) GetSecret(name string) {

	fmt.Println("test")
	secretResp, err := p.client.GetSecret(context.Background(), p.vaultURL, name, "")
	if err != nil {
		fmt.Printf("unable to get value for secret: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(*secretResp.Value)
}
