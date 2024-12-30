package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Viper *viper.Viper

func init() {

	// Directly read the environment variable using os.Getenv
	env := os.Getenv("APP_ENV")
	if env == "" {
		log.Println("APP_ENV not set, defaulting to 'development'")
		env = "development"
	}

	// Initialize Viper
	v := viper.New()

	if err := checkFileExists(); err != nil {
		log.Println(err)
		log.Printf("Loaded From Enivronment Variables. APP_ENV: %s \n", env)
		v.AutomaticEnv()
	} else {
		log.Print("File .env exists\n")
		log.Print("Loaded .env file\n")

		// Set the configuration file based on the environment
		v.SetConfigFile(fmt.Sprint(".env"))
		v.SetConfigType("dotenv")
		v.AddConfigPath("./")

		_ = v.ReadInConfig()
	}

	Viper = v
}

func checkFileExists() error {
	fileName := fmt.Sprint(".env")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", fileName)
	}
	return nil
}
