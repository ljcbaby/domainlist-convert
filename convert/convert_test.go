package convert_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dn-11/provider2domainset/conf"
	"github.com/dn-11/provider2domainset/convert"
)

func TestConvert(t *testing.T) {
	// assert pwd
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	t.Log(pwd)

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
