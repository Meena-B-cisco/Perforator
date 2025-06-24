package server

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/yandex/perforator/perforator/pkg/storage/bundle"
	"github.com/yandex/perforator/perforator/pkg/storage/microscope/filter"
)

type TvmAuth struct {
	ID            uint32 `yaml:"id"`
	SecretEnvName string `yaml:"secret_env"`
}

type TLSConfig struct {
	Enabled         bool   `yaml:"enabled"`
	ClientCAFile    string `yaml:"ca_file_path"`
	CertificateFile string `yaml:"certificate_file_path"`
	KeyFile         string `yaml:"key_file_path"`
	VerifyClient    bool   `yaml:"verify_client"`

	CertificateFileDeprecated string `yaml:"certificate_file"`
	KeyFileDeprecated         string `yaml:"key_file"`
}

type Config struct {
	Port                   uint32         `yaml:"port"`
	MetricsPort            uint32         `yaml:"metrics_port"`
	StorageConfig          bundle.Config  `yaml:"storage"`
	TvmAuth                *TvmAuth       `yaml:"tvm"`
	TLS                    TLSConfig      `yaml:"tls"`
	MicroscopePullerConfig *filter.Config `yaml:"microscope_puller"`
}

func ParseConfig(path string, strict bool) (conf *Config, err error) {
	// TODO(PERFORATOR-480): always be strict
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	conf = &Config{}
	dec := yaml.NewDecoder(file)
	dec.KnownFields(strict)
	err = dec.Decode(conf)
	return
}

func (c *Config) FillDefault() {

	// TLS backward compatibility.
	// Previously, agent communicated with storage via only TLS, so you need to enable tls if these values ​​are present.
	if c.TLS.CertificateFileDeprecated != "" {
		c.TLS.Enabled = true
		c.TLS.CertificateFile = c.TLS.CertificateFileDeprecated
	}

	if c.TLS.KeyFileDeprecated != "" {
		c.TLS.Enabled = true
		c.TLS.KeyFile = c.TLS.KeyFileDeprecated
	}
}
