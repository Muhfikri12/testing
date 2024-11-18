package util

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	AppName string
	Port    string
	Debug   bool
	DB      DatabaseConfig
}

type DatabaseConfig struct {
	Name     string
	Username string
	Password string
	Host     string
}

func ReadConfiguration() (Configuration, error) {
	// Mengatur agar Viper membaca variabel environment secara otomatis
	viper.AutomaticEnv()

	// Menentukan nama file konfigurasi tanpa ekstensi
	viper.SetConfigName(".env")

	// Menentukan format file konfigurasi sebagai env
	viper.SetConfigType("env")

	// Menentukan jalur direktori tempat file konfigurasi berada
	viper.AddConfigPath(".")

	// Membaca file konfigurasi
	if err := viper.ReadInConfig(); err != nil {
		return Configuration{}, err
	}

	// Membaca nilai konfigurasi dan mengembalikan hasil
	return Configuration{
		AppName: viper.GetString("APP_NAME"),
		Port:    viper.GetString("PORT"),
		Debug:   viper.GetBool("DEBUG"),
		DB: DatabaseConfig{
			Name:     viper.GetString("DATABASE_NAME"),
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Host:     viper.GetString("DATABASE_HOST"),
		},
	}, nil
}
