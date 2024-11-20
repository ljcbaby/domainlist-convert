package conf_test

import (
	"testing"

	"github.com/dn-11/provider2domainset/conf"
)

func TestParseConfig(t *testing.T) {
	conf.Init("config-sample.yaml")
}
