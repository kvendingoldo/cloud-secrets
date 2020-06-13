package validation

import (
	"errors"
	"fmt"
	"github.com/kvendingoldo/cloud-secrets/pkg/apis/cloudsecrets"
)

func ValidateConfig(cfg *cloudsecrets.Config) error {
	if cfg.LogFormat != "text" && cfg.LogFormat != "json" {
		return fmt.Errorf("unsupported log format: %s", cfg.LogFormat)
	}
	if cfg.Provider == "" {
		return errors.New("no provider specified")
	}

	if cfg.SecretName == "" {
		return errors.New("no secret name specified")
	}

	return nil
}
