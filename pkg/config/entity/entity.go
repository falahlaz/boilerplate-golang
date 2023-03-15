package entity

type ConfigData struct {
	Server     Server     `env:"server"`
	Redis      Redis      `env:"redis"`
	Multiconns []string   `env:"multiconns"`
	Databases  []Database `env:"databases"`
}

func NewConfigData() *ConfigData {
	return &ConfigData{}
}

type Server struct {
	Name    string `default:"new-application" env:"APP_NAME"`
	Version string `default:"0.0.1" env:"APP_VERSION"`
	Port    string `default:"3000" env:"APP_PORT"`
	Env     string `default:"development" env:"APP_ENV"`

	MaxOpenConn int `default:"200" env:"DB_MAX_OPEN_CONN"`
	MaxIdleConn int `default:"100" env:"DB_MAX_IDLE_CONN"`
	MaxLifetime int `default:"10" env:"DB_MAX_LIFETIME"`
}

type Database struct {
	Connection string `required:"true" env:"DB_CONNECTION"`
	Host       string `required:"true" env:"DB_HOST"`
	User       string `required:"true" env:"DB_USER"`
	Pass       string `required:"true" env:"DB_PASS"`
	Port       string `required:"true" env:"DB_PORT"`
	Name       string `required:"true" env:"DB_NAME"`
	Ssl        string `env:"DB_SSLMODE"`
	Tz         string `env:"DB_TZ"`
	Charset    string `env:"DB_CHARSET"`
	Location   string `env:"DB_LOCATION"`
	ParseTime  string `env:"DB_PARSE_TIME"`
	Alias      string `required:"true" env:"DB_ALIAS"`
	Migration  bool   `default:"false" env:"DB_MIGRATION_ENABLED"`
	Seeder     bool   `default:"false" env:"DB_SEEDER_ENABLED"`
}

type Redis struct {
	Host         string `env:"REDIS_HOST"`
	Pass         string `env:"REDIS_PASS"`
	Port         string `env:"REDIS_PORT"`
	PoolSize     int    `default:"50" env:"REDIS_POOL_SIZE"`
	MaxConnAge   int    `default:"10" env:"REDIS_MAX_CONN_AGE"`
	MinIdleConns int    `default:"10" env:"REDIS_MIN_IDLE_CONNS"`
}
