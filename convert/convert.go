package convert

import (
	"os"
	"path/filepath"

	"github.com/ljcbaby/domainlist-convert/conf"
	"github.com/ljcbaby/domainlist-convert/log"
	"github.com/pkg/errors"
)

func Convert(t Task) error {
	log.L().Sugar().Infof("convert %s of %s", t.Name, t.Type)

	source, err := os.Open(filepath.Join(t.Source, t.Name))
	if err != nil {
		return errors.Wrap(err, "read source file failed")
	}
	defer source.Close()

	err = os.MkdirAll(t.Target, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "create target directory failed")
	}

	target, err := os.Create(filepath.Join(t.Target, t.Name))
	if err != nil {
		return errors.Wrap(err, "create target file failed")
	}
	defer target.Close()

	switch t.Type {
	case conf.TypeClassical:
		err = convertClassical(source, target)
		if err != nil {
			return errors.Wrap(err, "convert classical failed")
		}
	case conf.TypeDomain:
		err = convertDomain(source, target)
		if err != nil {
			return errors.Wrap(err, "convert domain failed")
		}
	default:
		return errors.New("unknown convert type")
	}

	return nil
}
