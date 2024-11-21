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
	ProcessFiles []*File
	EnableRegex  bool
}

var Service struct {
	Enable bool
	Delay  int
	DbFile string
}

var Log struct {
	Level zapcore.Level
}

func Init(file string) {
	if _, err := os.Stat(file); err != nil {
		if !os.IsNotExist(err) {
			log.L().Sugar().With(err).Errorf("get stat of %s failed", file)
		}
		log.L().Sugar().Infof("config not existed, creating at %s", file)
		created, err := os.Create(file)
		if err != nil {
			log.L().Sugar().With(err).Errorf("create config at %s failed", file)
		}
		if _, err := created.Write(configSample); err != nil {
			log.L().Sugar().With(err).Errorf("write config at %s failed", file)
		}
	}

	viper.SetConfigFile(file)
	err := viper.ReadInConfig()

	update()
	if err != nil {
		log.L().Sugar().With(err).Errorf("read config from %s failed", file)
	}
}

func update() {
	Convert.Source = viper.GetString("convert.source")
	Convert.Target = viper.GetString("convert.target")
	Convert.ProcessFiles = make([]*File, 0)
	for _, file := range viper.GetStringMap("convert.files") {
		f := file.(map[string]interface{})
		Convert.ProcessFiles = append(Convert.ProcessFiles, &File{
			Name: f["name"].(string),
			Type: f["type"].(string),
		})
	}
	Convert.EnableRegex = viper.GetBool("convert.enable_regex")

	Service.Enable = viper.GetBool("service.enable")
	Service.Delay = viper.GetInt("service.delay")
	Service.DbFile = viper.GetString("service.db_file")

	if level, err := zapcore.ParseLevel(viper.GetString("log.level")); err == nil {
		Log.Level = level
	}
}
