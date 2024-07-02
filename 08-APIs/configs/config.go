// github.com/spf13/viper
// github.com/go-chi/jwtauth

package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

func NewConfig() *conf {
	return cfg
}

type conf struct {
	dbDriver      string `mapstructure:"DB_DRIVER"` // o map é um funcionalidade do viper, mapeio igual o json // qual o banco que vai ser usado
	dbHost        string `mapstructure:"DB_HOST"`
	dbPort        string `mapstructure:"DB_PORT"`
	dbUser        string `mapstructure:"DB_USER"`
	dbPassword    string `mapstructure:"DB_PASSWORD"`
	dbName        string `mapstructure:"DB_NAME"`
	webServerPort string `mapstructure:"WEB_SERVER_PORT"`
	jwtSecret     string `mapstructure:"JWT_SECRET"`
	jwtExpiresIn  int64  `mapstructure:"JWT_EXPIRES_IN"`
	tokenAuth     *jwtauth.JWTAuth
}

func (c *conf) GetDBDriver() string {
	return c.dbDriver
}

func (c *conf) GetDBHost() string {
	return c.dbHost
}

func (c *conf) GetDBPort() string {
	return c.dbPort
}

func (c *conf) GetDBUser() string {
	return c.dbUser
}

func (c *conf) GetDBPassword() string {
	return c.dbPassword
}

func (c *conf) GetDBName() string {
	return c.dbName
}

func (c *conf) GetWebServerPort() string {
	return c.webServerPort
}

func (c *conf) GetJWTSecret() string {
	return c.jwtSecret
}

func (c *conf) GetJWTExpiresIn() int64 {
	return c.jwtExpiresIn
}

func init() {
	var cfg *conf

	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.IsSet("dbDriver") // olhar documentacao para linkar nome da conf no viper, pois ele nao trabalha com variavel nao exportavel
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

	cfg.tokenAuth = jwtauth.New("HS256", []byte(cfg.jwtSecret), nil)

}
