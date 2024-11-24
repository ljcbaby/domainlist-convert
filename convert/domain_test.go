package convert_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ljcbaby/domainlist-convert/conf"
	"github.com/ljcbaby/domainlist-convert/convert"
	"github.com/ljcbaby/domainlist-convert/log"
	"go.uber.org/zap/zapcore"
)

func TestConvertDomainTxt(t *testing.T) {
	log.L().SetLogLevel(zapcore.DebugLevel)
	conf.Convert.EnableRegex = true

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	task := convert.Task{
		Source: filepath.Join(pwd, "test_case"),
		Target: filepath.Join(pwd, "test_output"),
		File: conf.File{
			Name: "domain.txt",
			Type: "domain",
		},
	}

	err = convert.Convert(task)
	if err != nil {
		t.Error(err)
	}
}

func TestConvertDomainYaml(t *testing.T) {
	log.L().SetLogLevel(zapcore.DebugLevel)
	conf.Convert.EnableRegex = true

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	task := convert.Task{
		Source: filepath.Join(pwd, "test_case"),
		Target: filepath.Join(pwd, "test_output"),
		File: conf.File{
			Name: "domain.yaml",
			Type: "domain",
		},
	}

	err = convert.Convert(task)
	if err != nil {
		t.Error(err)
	}
}
