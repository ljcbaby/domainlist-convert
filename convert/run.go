package convert

import (
	"github.com/dn-11/provider2domainset/conf"
	"github.com/dn-11/provider2domainset/log"
)

func RunOnce() {
	log.L().Info("start convert now")

	for _, f := range conf.Convert.ProcessFiles {
		err := Convert(Task{
			Source: conf.Convert.Source,
			Target: conf.Convert.Target,
			File:   f,
		})
		if err != nil {
			log.L().Sugar().With(err).Error("convert failed")
		}
	}

	log.L().Info("convert done")
}
