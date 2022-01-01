package common

import "time"

type AppConfig struct {
	General struct {
		Timezone          string         `mapstructure:"timezone" validate:"required"`
		Location          *time.Location `mapstructure:"-"`
		DeleteBeforeWrite bool           `mapstructure:"-"`
	} `mapstructure:"general"`
	Influxdb struct {
		Bucket      string `mapstructure:"bucket" validate:"required"`
		Org         string `mapstructure:"org" validate:"required"`
		Measurement string `mapstructure:"measurement" validate:"required"`
		Url         string `mapstructure:"url" validate:"required"`
		Token       string `mapstructure:"token" validate:"required"`
	} `mapstructure:"influxdb"`
	SolarEdge struct {
		ApiKey     string `mapstructure:"api-key" validate:"required"`
		SiteId     string `mapstructure:"site-id" validate:"required"`
		InverterSn string `mapstructure:"inverter-sn" validate:"required"`
	} `mapstructure:"solaredge"`
}
