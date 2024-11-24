package conf_test

import (
	"testing"

	"github.com/ljcbaby/domainlist-convert/conf"
	"github.com/ljcbaby/domainlist-convert/log"
)

func TestParseConfig(t *testing.T) {
	conf.Init("config-sample.yaml")

	log.L().Sugar().Infof("log: %+v", conf.Log)
	log.L().Sugar().Infof("convert: %+v", conf.Convert)
	log.L().Sugar().Infof("service: %+v", conf.Service)
}
