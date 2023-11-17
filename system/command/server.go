package command

type Database struct {
	Driver string
	Source string
	// ConnMaxIdleTime int
	// ConnMaxLifeTime int
	// MaxIdleConns    int
	// MaxOpenConns    int
	// Registers       []DBResolverConfig
}

// Config 配置集合
type Config struct {
	// Application *Application          `yaml:"application"`
	// Ssl         *Ssl                  `yaml:"ssl"`
	// Logger      *Logger               `yaml:"logger"`
	// Jwt         *Jwt                  `yaml:"jwt"`
	Database *Database `yaml:"database"`
	// Databases   *map[string]*Database `yaml:"databases"`
	// Gen         *Gen                  `yaml:"gen"`
	// Cache       *Cache                `yaml:"cache"`
	// Queue       *Queue                `yaml:"queue"`
	// Locker      *Locker               `yaml:"locker"`
	// Extend      interface{}           `yaml:"extend"`
}

var (
	DatabaseConfig = new(Database)
)

func InitServer() {

}
