package google

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"fmt"

	"github.com/kvendingoldo/cloud-secrets/provider"
	log "github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type GoogleProvider struct {
	provider.BaseProvider
	client        *secretmanager.Client
	projectId     string
	secretVersion string
}

type GoogleConfig struct {
	ProjectId     string
	SecretVersion string
}

func NewGoogleProvider(googleConfig GoogleConfig) (*GoogleProvider, error) {

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	provider := &GoogleProvider{
		client:        client,
		projectId:     googleConfig.ProjectId,
		secretVersion: googleConfig.SecretVersion,
	}

	return provider, nil
}

func (p *GoogleProvider) GetSecret(name string) {
	ctx := context.Background()

	fmt.Println("projects/" + p.projectId + "/secrets/" + name + "/versions/" + p.secretVersion)

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/" + p.projectId + "/secrets/" + name + "/versions/" + p.secretVersion,
	}

	// Call the API.
	result, err := p.client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Infof("failed to get secret: %v", err)
	}

	fmt.Println(result.Payload)
}
