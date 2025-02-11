package conf

import (
	_ "embed"
	"os"

	"github.com/ljcbaby/domainlist-convert/log"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

//go:embed config-sample.yaml
var configSample []byte

var Convert struct {
	Source       string
	Target       string
	ProcessFiles []File
	EnableRegex  bool
}

var Service struct {
	Enable bool
	Delay  int
}

var Log struct {
	Level zapcore.Level
}

func Init(file string) {
	if _, err := os.Stat(file); err != nil {
		if !os.IsNotExist(err) {
			log.L().Sugar().With(err).Errorf("get stat of %s failed", file)
		}
		log.L().Sugar().Errorf("config not existed, creating at %s", file)
		created, err := os.Create(file)
		if err != nil {
			log.L().Sugar().With(err).Errorf("create config at %s failed", file)
		}
		if _, err := created.Write(configSample); err != nil {
			log.L().Sugar().With(err).Errorf("write config at %s failed", file)
		}
		log.L().Fatal("config created, please edit it.")
	}

	viper.SetConfigFile(file)
	err := viper.ReadInConfig()

	viper.SetDefault("service.delay", 60)

	update()
	if err != nil {
		log.L().Sugar().With(err).Errorf("read config from %s failed", file)
	}
}

func update() {
	Convert.Source = viper.GetString("convert.source")
	Convert.Target = viper.GetString("convert.target")

	if Convert.Source == "" || Convert.Target == "" {
		log.L().Sugar().Fatal("source or target is empty")
	}

	Convert.ProcessFiles = make([]File, 0)
	for _, file := range viper.Get("convert.processFiles").([]interface{}) {
		t := file.(map[string]interface{})
		f := File{
			Name: t["name"].(string),
			Type: t["type"].(string),
		}

		switch f.Type {
		case TypeClassical, TypeDomain:
		default:
			log.L().Sugar().Fatalf("%s unknown type, skip it: %s", f.Type, f.Name)
			continue
		}

		Convert.ProcessFiles = append(Convert.ProcessFiles, f)
	}

	if len(Convert.ProcessFiles) == 0 {
		log.L().Sugar().Fatal("no files to process")
	}

	Convert.EnableRegex = viper.GetBool("convert.enable_regex")

	Service.Enable = viper.GetBool("service.enable")
	Service.Delay = viper.GetInt("service.delay")

	if level, err := zapcore.ParseLevel(viper.GetString("log.level")); err == nil {
		if log.L().Level.Level() == zapcore.InfoLevel {
			Log.Level = level
			log.L().SetLogLevel(level)
		}
	}
}
