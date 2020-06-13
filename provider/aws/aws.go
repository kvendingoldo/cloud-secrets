package aws

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/kvendingoldo/cloud-secrets/provider"
	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"
	"strings"
)

const ()

type SecretsManagerAPI interface {
	CancelRotateSecret(*secretsmanager.CancelRotateSecretInput) (*secretsmanager.CancelRotateSecretOutput, error)
	CancelRotateSecretWithContext(aws.Context, *secretsmanager.CancelRotateSecretInput, ...request.Option) (*secretsmanager.CancelRotateSecretOutput, error)
	CancelRotateSecretRequest(*secretsmanager.CancelRotateSecretInput) (*request.Request, *secretsmanager.CancelRotateSecretOutput)

	CreateSecret(*secretsmanager.CreateSecretInput) (*secretsmanager.CreateSecretOutput, error)
	CreateSecretWithContext(aws.Context, *secretsmanager.CreateSecretInput, ...request.Option) (*secretsmanager.CreateSecretOutput, error)
	CreateSecretRequest(*secretsmanager.CreateSecretInput) (*request.Request, *secretsmanager.CreateSecretOutput)

	DeleteResourcePolicy(*secretsmanager.DeleteResourcePolicyInput) (*secretsmanager.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyWithContext(aws.Context, *secretsmanager.DeleteResourcePolicyInput, ...request.Option) (*secretsmanager.DeleteResourcePolicyOutput, error)
	DeleteResourcePolicyRequest(*secretsmanager.DeleteResourcePolicyInput) (*request.Request, *secretsmanager.DeleteResourcePolicyOutput)

	DeleteSecret(*secretsmanager.DeleteSecretInput) (*secretsmanager.DeleteSecretOutput, error)
	DeleteSecretWithContext(aws.Context, *secretsmanager.DeleteSecretInput, ...request.Option) (*secretsmanager.DeleteSecretOutput, error)
	DeleteSecretRequest(*secretsmanager.DeleteSecretInput) (*request.Request, *secretsmanager.DeleteSecretOutput)

	DescribeSecret(*secretsmanager.DescribeSecretInput) (*secretsmanager.DescribeSecretOutput, error)
	DescribeSecretWithContext(aws.Context, *secretsmanager.DescribeSecretInput, ...request.Option) (*secretsmanager.DescribeSecretOutput, error)
	DescribeSecretRequest(*secretsmanager.DescribeSecretInput) (*request.Request, *secretsmanager.DescribeSecretOutput)

	GetRandomPassword(*secretsmanager.GetRandomPasswordInput) (*secretsmanager.GetRandomPasswordOutput, error)
	GetRandomPasswordWithContext(aws.Context, *secretsmanager.GetRandomPasswordInput, ...request.Option) (*secretsmanager.GetRandomPasswordOutput, error)
	GetRandomPasswordRequest(*secretsmanager.GetRandomPasswordInput) (*request.Request, *secretsmanager.GetRandomPasswordOutput)

	GetResourcePolicy(*secretsmanager.GetResourcePolicyInput) (*secretsmanager.GetResourcePolicyOutput, error)
	GetResourcePolicyWithContext(aws.Context, *secretsmanager.GetResourcePolicyInput, ...request.Option) (*secretsmanager.GetResourcePolicyOutput, error)
	GetResourcePolicyRequest(*secretsmanager.GetResourcePolicyInput) (*request.Request, *secretsmanager.GetResourcePolicyOutput)

	GetSecretValue(*secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
	GetSecretValueWithContext(aws.Context, *secretsmanager.GetSecretValueInput, ...request.Option) (*secretsmanager.GetSecretValueOutput, error)
	GetSecretValueRequest(*secretsmanager.GetSecretValueInput) (*request.Request, *secretsmanager.GetSecretValueOutput)

	ListSecretVersionIds(*secretsmanager.ListSecretVersionIdsInput) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecretVersionIdsWithContext(aws.Context, *secretsmanager.ListSecretVersionIdsInput, ...request.Option) (*secretsmanager.ListSecretVersionIdsOutput, error)
	ListSecretVersionIdsRequest(*secretsmanager.ListSecretVersionIdsInput) (*request.Request, *secretsmanager.ListSecretVersionIdsOutput)

	ListSecretVersionIdsPages(*secretsmanager.ListSecretVersionIdsInput, func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool) error
	ListSecretVersionIdsPagesWithContext(aws.Context, *secretsmanager.ListSecretVersionIdsInput, func(*secretsmanager.ListSecretVersionIdsOutput, bool) bool, ...request.Option) error

	ListSecrets(*secretsmanager.ListSecretsInput) (*secretsmanager.ListSecretsOutput, error)
	ListSecretsWithContext(aws.Context, *secretsmanager.ListSecretsInput, ...request.Option) (*secretsmanager.ListSecretsOutput, error)
	ListSecretsRequest(*secretsmanager.ListSecretsInput) (*request.Request, *secretsmanager.ListSecretsOutput)

	ListSecretsPages(*secretsmanager.ListSecretsInput, func(*secretsmanager.ListSecretsOutput, bool) bool) error
	ListSecretsPagesWithContext(aws.Context, *secretsmanager.ListSecretsInput, func(*secretsmanager.ListSecretsOutput, bool) bool, ...request.Option) error

	PutResourcePolicy(*secretsmanager.PutResourcePolicyInput) (*secretsmanager.PutResourcePolicyOutput, error)
	PutResourcePolicyWithContext(aws.Context, *secretsmanager.PutResourcePolicyInput, ...request.Option) (*secretsmanager.PutResourcePolicyOutput, error)
	PutResourcePolicyRequest(*secretsmanager.PutResourcePolicyInput) (*request.Request, *secretsmanager.PutResourcePolicyOutput)

	PutSecretValue(*secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error)
	PutSecretValueWithContext(aws.Context, *secretsmanager.PutSecretValueInput, ...request.Option) (*secretsmanager.PutSecretValueOutput, error)
	PutSecretValueRequest(*secretsmanager.PutSecretValueInput) (*request.Request, *secretsmanager.PutSecretValueOutput)

	RestoreSecret(*secretsmanager.RestoreSecretInput) (*secretsmanager.RestoreSecretOutput, error)
	RestoreSecretWithContext(aws.Context, *secretsmanager.RestoreSecretInput, ...request.Option) (*secretsmanager.RestoreSecretOutput, error)
	RestoreSecretRequest(*secretsmanager.RestoreSecretInput) (*request.Request, *secretsmanager.RestoreSecretOutput)

	RotateSecret(*secretsmanager.RotateSecretInput) (*secretsmanager.RotateSecretOutput, error)
	RotateSecretWithContext(aws.Context, *secretsmanager.RotateSecretInput, ...request.Option) (*secretsmanager.RotateSecretOutput, error)
	RotateSecretRequest(*secretsmanager.RotateSecretInput) (*request.Request, *secretsmanager.RotateSecretOutput)

	TagResource(*secretsmanager.TagResourceInput) (*secretsmanager.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *secretsmanager.TagResourceInput, ...request.Option) (*secretsmanager.TagResourceOutput, error)
	TagResourceRequest(*secretsmanager.TagResourceInput) (*request.Request, *secretsmanager.TagResourceOutput)

	UntagResource(*secretsmanager.UntagResourceInput) (*secretsmanager.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *secretsmanager.UntagResourceInput, ...request.Option) (*secretsmanager.UntagResourceOutput, error)
	UntagResourceRequest(*secretsmanager.UntagResourceInput) (*request.Request, *secretsmanager.UntagResourceOutput)

	UpdateSecret(*secretsmanager.UpdateSecretInput) (*secretsmanager.UpdateSecretOutput, error)
	UpdateSecretWithContext(aws.Context, *secretsmanager.UpdateSecretInput, ...request.Option) (*secretsmanager.UpdateSecretOutput, error)
	UpdateSecretRequest(*secretsmanager.UpdateSecretInput) (*request.Request, *secretsmanager.UpdateSecretOutput)

	UpdateSecretVersionStage(*secretsmanager.UpdateSecretVersionStageInput) (*secretsmanager.UpdateSecretVersionStageOutput, error)
	UpdateSecretVersionStageWithContext(aws.Context, *secretsmanager.UpdateSecretVersionStageInput, ...request.Option) (*secretsmanager.UpdateSecretVersionStageOutput, error)
	UpdateSecretVersionStageRequest(*secretsmanager.UpdateSecretVersionStageInput) (*request.Request, *secretsmanager.UpdateSecretVersionStageOutput)
}

type AWSProvider struct {
	provider.BaseProvider
	client SecretsManagerAPI
}

// AWSConfig contains configuration to create a new AWS provider.
type AWSConfig struct {
	Region     string
	APIRetries int
	AssumeRole string
}

func NewAWSProvider(awsConfig AWSConfig) (*AWSProvider, error) {
	config := aws.NewConfig().WithMaxRetries(awsConfig.APIRetries).WithRegion(awsConfig.Region)

	config.WithHTTPClient(
		instrumented_http.NewClient(config.HTTPClient, &instrumented_http.Callbacks{
			PathProcessor: func(path string) string {
				parts := strings.Split(path, "/")
				return parts[len(parts)-1]
			},
		}),
	)

	session, err := session.NewSessionWithOptions(session.Options{
		Config:            *config,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	if awsConfig.AssumeRole != "" {
		log.Infof("Assuming role: %s", awsConfig.AssumeRole)
		session.Config.WithCredentials(stscreds.NewCredentials(session, awsConfig.AssumeRole))
	}
	provider := &AWSProvider{
		client: secretsmanager.New(session),
	}

	return provider, nil
}

func (p *AWSProvider) GetSecret(name string) {


	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(name),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := p.client.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
	}

	fmt.Println(secretString)
	fmt.Println(decodedBinarySecret)
}
