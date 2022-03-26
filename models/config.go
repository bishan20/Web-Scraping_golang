package models

type Config struct {
	DBDriver            string `mapstructure:"DB_DRIVER"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	DBSource            string `mapstructure:"DB_SOURCE"`
	ServerAddress       string `mapstructure:"SERVER_ADDRESS"`
	AccessTokenDuration string `mapstructure:"ACCESS_TOKEN_DURATION"`
}
