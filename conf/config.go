package conf

import (
	_ "embed"
	"os"

	"github.com/dn-11/provider2domainset/log"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

//go:embed config-sample.yaml
var configSample []byte

var Convert struct {
	Source       string
	Target       string
	ProcessFiles []string
}

var Service struct {
	Enable bool
	DbFile string
}

var Log struct {
	Level zapcore.Level
}

func Init(file string) {
	if _, err := os.Stat(file); err != nil {
		if !os.IsNotExist(err) {
			log.Logger.Sugar().With(err).Errorf("get stat of %s failed", file)
		}
		log.Logger.Sugar().Infof("config not existed, creating at %s", file)
		created, err := os.Create(file)
		if err != nil {
			log.Logger.Sugar().With(err).Errorf("create config at %s failed", file)
		}
		if _, err := created.Write(configSample); err != nil {
			log.Logger.Sugar().With(err).Errorf("write config at %s failed", file)
		}
	}

	viper.SetConfigFile(file)
	err := viper.ReadInConfig()

	update()
	if err != nil {
		log.Logger.Sugar().With(err).Errorf("read config from %s failed", file)
	}
}

func update() {
	Convert.Source = viper.GetString("convert.source")
	Convert.Target = viper.GetString("convert.target")
	Convert.ProcessFiles = viper.GetStringSlice("convert.process_files")

	Service.Enable = viper.GetBool("service.enable")
	Service.DbFile = viper.GetString("service.db_file")

	if level, err := zapcore.ParseLevel(viper.GetString("log.level")); err == nil {
		Log.Level = level
	}
}
