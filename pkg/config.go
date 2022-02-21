package pkg

import (
	phoneapi "moment/pkg/phoneAPI"
	"time"

	"gitlab.com/knopkalab/go/config"
	"gitlab.com/knopkalab/go/database/reindexer"
	"gitlab.com/knopkalab/go/http"
	"gitlab.com/knopkalab/go/logger"
)

const configFileName = "moment.conf"

type Config struct {
	config.Value
	DBReindexer reindexer.Config
	Ucaller     phoneapi.Ucaller
	Log         logger.Config
	Server      http.ServerConfig

	Sessions struct {
		CookieName   string `config:"-"`
		CookieMaxAge config.Duration
	}
}

func (c *Config) setDefaults() {
	c.Log.Level = logger.LevelTrace
	c.Log.Stdout = true
	c.Log.Filename = configFileName
	c.Server.Port = 7000

	c.DBReindexer.Addr = "127.0.0.1:6534"
	c.DBReindexer.Name = "moment"

	c.Ucaller.ServiceID = "131537"
	c.Ucaller.SecretKey = "0bz4hbS0LvuSRR6nkGJK2pHIKRYSIlU9"
	c.Ucaller.Debug = false

	c.Sessions.CookieName = "session_id"
	c.Sessions.CookieMaxAge.Duration = time.Hour * 24 * 7

}

func (c *Config) Repair() error {
	c.Log.Repair()
	c.Server.Repair()
	err := c.DBReindexer.ValidAndRepair()

	return err
}

func (c *Config) Save() error {
	if err := c.Repair(); err != nil {
		return err
	}
	return config.Save(c, configFileName)
}

func LoadConfig() (*Config, error) {
	c := new(Config)
	c.setDefaults()
	c.Repair()

	err := config.Load(c, configFileName)
	return c, err
}
