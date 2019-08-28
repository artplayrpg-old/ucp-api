package config

// JWTConfig is the "jwt" configuration.
type JWTConfig struct {
	Secret string `default:"secret"`
}

// Config is the configuration.
type Config struct {
	JWT JWTConfig

	Bind string `default:"0.0.0.0:7788"`
	Host string `split_words:"true" required:"false"`
	Port int    `default:"3306"`
	User string `split_words:"true" required:"false"`
	Pass string `split_words:"true" required:"false"`
	Name string `split_words:"true" required:"false"`
}

// New is a new instance of configuration.
func New() *Config {
	return &Config{}
}