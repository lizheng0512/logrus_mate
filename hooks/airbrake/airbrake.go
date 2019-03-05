package airbrake

import (
	"github.com/gogap/config"
	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"
)

type AirbrakeHookConfig struct {
	ProjectId int64
	APIKey    string
	Env       string
}

func init() {
	logrus_mate.RegisterHook("airbrake", NewAirbrakeHook)
}

func NewAirbrakeHook(config config.Configuration) (hook logrus.Hook, err error) {

	conf := AirbrakeHookConfig{}
	if config != nil {
		conf.ProjectId = config.GetInt64("project-id")
		conf.APIKey = config.GetString("api-key")
		conf.Env = config.GetString("env")
	}

	hook = airbrake.NewHook(
		conf.ProjectId,
		conf.APIKey,
		conf.Env,
	)

	return
}
