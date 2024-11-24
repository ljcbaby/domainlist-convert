package daemon

import (
	"os"
	"path/filepath"

	"github.com/dn-11/provider2domainset/conf"
	"github.com/dn-11/provider2domainset/convert"
	"github.com/dn-11/provider2domainset/log"
)

func start() {
	for _, f := range conf.Convert.ProcessFiles {
		s, err := os.Stat(filepath.Join(conf.Convert.Source, f.Name))
		if err != nil {
			if os.IsNotExist(err) {
				log.L().Sugar().Errorf("File not exist: %s", f.Name)
			} else {
				log.L().Sugar().Errorf("Get %s file info error: %v", f.Name, err)
			}
			continue
		}

		t, err := os.Stat(filepath.Join(conf.Convert.Target, f.Name))
		if err != nil && !os.IsNotExist(err) {
			log.L().Sugar().Errorf("Get %s file info error: %v", f.Name, err)
			continue
		}

		if os.IsNotExist(err) || s.ModTime().After(t.ModTime()) {
			err := convert.Convert(convert.Task{
				Source: conf.Convert.Source,
				Target: conf.Convert.Target,
				File:   f,
			})
			if err != nil {
				log.L().Sugar().Errorf("Convert %s error: %v", f.Name, err)
			}

			refreshChan <- struct{}{}
		}
	}

	log.L().Info("Start check finished.")
}
