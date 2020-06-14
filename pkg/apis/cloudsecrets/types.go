package cloudsecrets

import (
	"github.com/alecthomas/kingpin"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var (
	// Version is the current version of the app, generated at build time
	Version = "unknown"
)

type Config struct {
	SecretName     string
	Provider       string
	LogFormat      string
	LogLevel       string
	MetricsAddress string

	Interval time.Duration
	Once     bool

	AWSRegion     string
	AWSAssumeRole string
	AWSAPIRetries int

	AzureRegion        string
	AzureResourceGroup string
	AzureKeyVault      string
}

var defaultConfig = &Config{
	SecretName:     "",
	Provider:       "",
	LogFormat:      "text",
	LogLevel:       logrus.InfoLevel.String(),
	MetricsAddress: ":7979",

	Interval: time.Minute,
	Once:     true,

	AWSRegion:     "us-east-1",
	AWSAssumeRole: "",
	AWSAPIRetries: 3,

	AzureRegion: "centralus",
}

func NewConfig() *Config {
	return &Config{}
}

// allLogLevelsAsStrings returns all logrus levels as a list of strings
func allLogLevelsAsStrings() []string {
	var levels []string
	for _, level := range logrus.AllLevels {
		levels = append(levels, level.String())
	}
	return levels
}

func (cfg *Config) ParseFlags(args []string) error {
	app := kingpin.New("cloud-secrets", "Cloud Secrets is a thin wrapper under Clouds secret services.\n\nNote that all flags may be replaced with env vars - `--flag` -> `CLOUD_SECRETS=1` or `--flag value` -> `CLOUD_SECRETS_FLAG=value`")
	app.Version(Version)
	app.DefaultEnvars()

	// Flags related to processing sources
	app.Flag("secret-name", "Name of secret").Default(defaultConfig.SecretName).StringVar(&cfg.SecretName)

	// Flags related to providers
	app.Flag("provider", "The Cloud provider (required, options: aws, azure)").Required().PlaceHolder("provider").EnumVar(&cfg.Provider, "aws", "azure")
	// AWS
	app.Flag("aws-region", "").Default(defaultConfig.AWSRegion).StringVar(&cfg.AWSRegion)
	app.Flag("aws-assume-role", "When using the AWS provider, assume this IAM role. Useful for hosted zones in another AWS account. Specify the full ARN (optional)").Default(defaultConfig.AWSAssumeRole).StringVar(&cfg.AWSAssumeRole)
	app.Flag("aws-api-retries", "When using the AWS provider, set the maximum number of retries for API calls before giving up.").Default(strconv.Itoa(defaultConfig.AWSAPIRetries)).IntVar(&cfg.AWSAPIRetries)
	// Azure
	app.Flag("azure-region", "").Default(defaultConfig.AzureRegion).StringVar(&cfg.AzureRegion)
	app.Flag("azure-key-vault", "").StringVar(&cfg.AzureKeyVault)
	app.Flag("azure-resource-group", "").StringVar(&cfg.AzureResourceGroup)

	// Miscellaneous flags
	app.Flag("log-format", "The format in which log messages are printed (default: text, options: text, json)").Default(defaultConfig.LogFormat).EnumVar(&cfg.LogFormat, "text", "json")
	app.Flag("metrics-address", "Specify where to serve the metrics and health check endpoint (default: :7979)").Default(defaultConfig.MetricsAddress).StringVar(&cfg.MetricsAddress)
	app.Flag("log-level", "Set the level of logging. (default: info, options: panic, debug, info, warning, error, fatal").Default(defaultConfig.LogLevel).EnumVar(&cfg.LogLevel, allLogLevelsAsStrings()...)

	// Flags related to the main control loop
	app.Flag("interval", "The interval between two consecutive synchronizations in duration format (default: 1m)").Default(defaultConfig.Interval.String()).DurationVar(&cfg.Interval)
	app.Flag("once", "When enabled, exits the synchronization loop after the first iteration (default: disabled)").Default(strconv.FormatBool(defaultConfig.Once)).BoolVar(&cfg.Once)

	_, err := app.Parse(args)
	if err != nil {
		return err
	}

	return nil
}
