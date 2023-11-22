package core_config

type Config struct {
	MysqlConfig       MysqlConfig       `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	ApplicationConfig ApplicationConfig `mapstructure:"application" json:"application" yaml:"application"`
}
