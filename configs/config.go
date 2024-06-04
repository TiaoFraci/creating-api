package configs

import "github.com/spf13/viper"

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassWord    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JwtExperesIn  int
	TokenAuth     string
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName(path)
	viper.SetConfigType("env")

	return nil, nil

}
