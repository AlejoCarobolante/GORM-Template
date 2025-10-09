package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost                 string `mapstructure:"DB_HOST"`
	DBPort                 string `mapstructure:"DB_PORT"`
	DBUser                 string `mapstructure:"DB_USER"`
	DBPass                 string `mapstructure:"DB_PASS"`
	DBName                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	MigrationsFolder       string `mapstructure:"MIGRATIONS_FOLDER"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := nil {
		log.Println("Can't find file .env, using OS env instead.")
		return NewEnvWithoutFile(env)
	}

	err = viper.Unmarshal(&env)
	if err !=nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	
	if env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}
	
	return &env
}

func NewEnvWithoutFile(env Env) *Env { // Se usa en caso de que no se provea un .env, se usa las variables del SO en su lugar
	env.DBHost = os.Getenv("DB_HOST")
	env.DBName = os.Getenv("DB_NAME")
	env.DBPass = os.Getenv("DB_PASS")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBUser = os.Getenv("DB_USER")
	env.AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	env.ServerAddress = os.Getenv("SERVER_ADDRESS")
	env.MigrationsFolder = os.Getenv("MIGRATIONS_FOLDER")
	return &env
}
