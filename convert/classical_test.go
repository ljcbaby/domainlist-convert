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

func TestConvertClassicalTxt(t *testing.T) {
	log.L().SetLogLevel(zapcore.DebugLevel)

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	task := convert.Task{
		Source: filepath.Join(pwd, "test_case"),
		Target: filepath.Join(pwd, "test_output"),
		File: conf.File{
			Name: "classical.txt",
			Type: "classical",
		},
	}

	err = convert.Convert(task)
	if err != nil {
		t.Error(err)
	}
}

func TestConvertClassicalYaml(t *testing.T) {
	log.L().SetLogLevel(zapcore.DebugLevel)

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	task := convert.Task{
		Source: filepath.Join(pwd, "test_case"),
		Target: filepath.Join(pwd, "test_output"),
		File: conf.File{
			Name: "classical.yaml",
			Type: "classical",
		},
	}

	err = convert.Convert(task)
	if err != nil {
		t.Error(err)
	}
}
