package configs

import (
	"core-auth-api/helper"
)

type AppSettings struct {
	Port string `json:"Port" env:"PORT"`

	JWT struct {
		Issuer            string `json:"Issuer" env:"SSO_ISSUER"`
		ExpirationMinutes int    `json:"ExpirationMinutes" env:"SSO_EXPIRATION_MINUTES"`
		Secret            string `json:"Secret" env:"SSO_SECRET"`
	} `json:"SSO_Config"`
}

func Load() AppSettings {
	return helper.Load[AppSettings]()
}
