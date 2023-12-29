package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	dbUser                  = "db.user"
	dbPassword              = "db.password"
	dbName                  = "db.name"
	dbHost                  = "db.host"
	dbEnableMultiStatements = "db.enable_multi_statement"
)

type Conf struct {
	viperConfig *viper.Viper
}

type DBConnSettings struct {
	User                  string
	Password              string
	Database              string
	Host                  string
	EnableMultiStatements bool
}

func NewConf() (*Conf, error) {
	config := &Conf{viperConfig: viper.New()}
	err := config.readFromConfFile()
	if err != nil {
		return nil, fmt.Errorf("Error while trying to read config file: %v", err)
	}

	err = config.checkRequiredProperties()
	if err != nil {
		return nil, fmt.Errorf("Error while checking properties: %v", err)
	}

	config.viperConfig.SetDefault("listen", "8000")

	return config, nil
}

func (c *Conf) readFromConfFile() error {
	c.viperConfig.SetConfigName("config")
	c.viperConfig.SetConfigType("yaml")
	c.viperConfig.AddConfigPath("./config")
	return c.viperConfig.ReadInConfig()
}

func (c *Conf) checkRequiredProperties() error {
	d := c.DBConnSettings()
	if d.Database == "" {
		return fmt.Errorf("%v is empty", dbName)
	}
	if d.User == "" {
		return fmt.Errorf("%v is empty", dbUser)
	}
	if d.Password == "" {
		return fmt.Errorf("%v is empty", dbPassword)
	}
	if d.Host == "" {
		return fmt.Errorf("%v is empty", dbHost)
	}
	return nil
}

func (c *Conf) DBConnSettings() DBConnSettings {
	return DBConnSettings{
		User:                  c.viperConfig.GetString(dbUser),
		Password:              c.viperConfig.GetString(dbPassword),
		Database:              c.viperConfig.GetString(dbName),
		Host:                  c.viperConfig.GetString(dbHost),
		EnableMultiStatements: c.viperConfig.GetBool(dbEnableMultiStatements),
	}
}
