package core_database

var (
	DBdd Config
)

type Config struct {
	// App App `mapstructure:"app" yaml:"app"`
	Database DatabaseModel `mapstructure:"database" yaml:"database"`
	// Redis Redis `mapstructure:"redis" yaml:"redis"`
	// Server Server `mapstructure:"server" yaml:"server"`
	// Zap Zap `mapstructure:"zap" yaml:"zap"`
	// Wechat Wechat `mapstructure:"wechat" yaml:"wechat"`
	// Express Express `mapstructure:"express" yaml:"express"`
}

type DatabaseModel struct {
	Driver    string
	Source    string
	Localhost string
	Port      string
	DBname    string
	Username  string
	Password  string
	// ConnMaxIdleTime int
	// ConnMaxLifeTime int
	// MaxIdleConns    int
	// MaxOpenConns    int
	// Registers       []DBResolverConfig
}

type DBResolverConfig struct {
	Sources  []string
	Replicas []string
	Policy   string
	Tables   []string
}

var (
// DatabaseConfig  = new(Database)
// DatabasesConfig = make(map[string]*Database)
)
