package main

import (
	"context"
	"github.com/kvendingoldo/cloud-secrets/controller"
	"github.com/kvendingoldo/cloud-secrets/pkg/apis/cloudsecrets"
	"github.com/kvendingoldo/cloud-secrets/pkg/apis/cloudsecrets/validation"
	"github.com/kvendingoldo/cloud-secrets/provider"
	"github.com/kvendingoldo/cloud-secrets/provider/aws"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := cloudsecrets.NewConfig()
	if err := cfg.ParseFlags(os.Args[1:]); err != nil {
		log.Fatalf("flag parsing error: %v", err)
	}
	log.Infof("config: %s", cfg)

	if err := validation.ValidateConfig(cfg); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	if cfg.LogFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	ll, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to parse log level: %v", err)
	}
	log.SetLevel(ll)

	ctx, cancel := context.WithCancel(context.Background())

	go serveMetrics(cfg.MetricsAddress)
	go handleSigterm(cancel)

	var p provider.Provider
	switch cfg.Provider {
	case "aws":
		p, err = aws.NewAWSProvider(
			aws.AWSConfig{
				Region:     cfg.AWSRegion,
				AssumeRole: cfg.AWSAssumeRole,
				APIRetries: cfg.AWSAPIRetries,
			},
		)
	default:
		log.Fatalf("unknown provider: %s", cfg.Provider)
	}
	if err != nil {
		log.Fatal(err)
	}

	ctrl := controller.Controller{
		Provider:   p,
		Interval:   cfg.Interval,
		SecretName: cfg.SecretName,
	}

	if cfg.Once {
		err := ctrl.RunOnce(ctx)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}

	ctrl.ScheduleRunOnce(time.Now())
	ctrl.Run(ctx)
}

func handleSigterm(cancel func()) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)
	<-signals
	log.Info("Received SIGTERM. Terminating...")
	cancel()
}

func serveMetrics(address string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(address, nil))
}
