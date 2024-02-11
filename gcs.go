package config

import "github.com/caarlos0/env/v9"

type GcsConfig struct {
	Credentials string `env:"GCS_CREDENTIALS" envDefault:"credentials/gcs.json"`
	BucketName  string `env:"GCS_BUCKET_NAME" envDefault:"image-bucket"`
}

func NewGcsConfig() (*GcsConfig, error) {
	cfg := &GcsConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
