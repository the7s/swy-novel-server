package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Website Website `mapstructure:"website" json:"website" yaml:"website"`
}
