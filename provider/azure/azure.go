package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest/azure"

	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/kvendingoldo/cloud-secrets/provider"
	"os"
)

type AzureProvider struct {
	provider.BaseProvider
	client   *keyvault.BaseClient
	vaultURL string
}

type AzureConfig struct {
	Region        string
	ResourceGroup string
	KeyVault      string
}

func NewAzureProvider(azureConfig AzureConfig) (*AzureProvider, error) {
	// TODO: need check auth method later
	authorizer, err := kvauth.NewAuthorizerFromCLI()
	if err != nil {
		fmt.Printf("unable to create vault authorizer: %v\n", err)
		os.Exit(1)
	}

	keyClient := keyvault.New()
	keyClient.Authorizer = authorizer

	provider := &AzureProvider{
		client:   &keyClient,
		vaultURL: fmt.Sprintf("https://%s.%s", azureConfig.KeyVault, azure.PublicCloud.KeyVaultDNSSuffix),
	}

	return provider, nil
}

func (p *AzureProvider) GetSecret(name string) {

	secretResp, err := p.client.GetSecret(context.Background(), p.vaultURL, name, "")
	if err != nil {
		fmt.Printf("unable to get value for secret: %v\n", err)
	}
	fmt.Println(*secretResp.Value)
}
