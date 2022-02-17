package config

type Website struct {
	QiDianUrl string `mapstructure:"qiDian-url" json:"qiDianUrl" yaml:"qiDian-url"`
	BQGUrl    string `mapstructure:"bqg-url" json:"bqg-url" yaml:"bqg-url"`
}
