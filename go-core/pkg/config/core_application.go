package core_config

type ApplicationConfig struct {
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Host          string `mapstructure:"host" json:"host" yaml:"host"`
	Name          string `mapstructure:"name" json:"name" yaml:"name"`
	Port          int64  `mapstructure:"port" json:"port" yaml:"port"`
	ReadTimeout   int    `mapstructure:"readtimeout" json:"read_timeout" yaml:"readtimeout"`
	WriterTimeout int    `mapstructure:"writer_timeout" json:"writer_timeout" yaml:"writertimeout"`
}
