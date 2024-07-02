// github.com/spf13/viper
// github.com/go-chi/jwtauth

package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"` // o map é um funcionalidade do viper, mapeio igual o json // qual o banco que vai ser usado
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  int64  `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(filePath string) (*conf, error) {
	var cfg *conf
	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(filePath)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		//return nil
		//nao quero que aplicacao suba se tiver erro
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		//return nil
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)

	return cfg, nil
}
