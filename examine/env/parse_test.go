package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/caarlos0/env/v11"

	"github.com/valteem/examine/env/model"
)

func TestParseEnv(t *testing.T) {

	LoadFromEnvFile(".env")

	cfgActual := model.Config{}

	env.Parse(&cfgActual)

	cfgExpected := model.Config{
		App: model.App{
			Version: "0.0.0",
		},
		DB: model.DB{
			MaxPoolSize: 10,
			Url:         "postgresql://user:secret@localhost",
		},
		HTTP: model.HTTP{
			Port:         ":3001",
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
		},
	}

	if !reflect.DeepEqual(cfgActual, cfgExpected) {
		t.Errorf("get:\n%v\nexpect:\n%v\n", cfgActual, cfgExpected)
	}

}
