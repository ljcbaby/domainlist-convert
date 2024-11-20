package convert

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func Convert(t Task) error {
	source, err := os.ReadFile(filepath.Join(t.Source, t.Name))
	if err != nil {
		return errors.Wrap(err, "read source file failed")
	}

	err = os.MkdirAll(t.Target, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "create target directory failed")
	}

	target, err := os.Create(filepath.Join(t.Target, t.Name))
	if err != nil {
		return errors.Wrap(err, "create target file failed")
	}

	_, err = target.Write(source)
	if err != nil {
		return errors.Wrap(err, "write target file failed")
	}

	return nil
}
