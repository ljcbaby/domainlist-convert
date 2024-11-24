package daemon

import (
	_ "embed"
	"errors"

	"github.com/ljcbaby/domainlist-convert/log"
	"go.uber.org/zap"

	"os"
	"os/exec"
)

const ServicePath = "/etc/init.d/provider2domainset"

//go:embed provider2domainset
var ServiceFile []byte

func AddService() {
	_, err := exec.LookPath("provider2domainset")
	if err != nil {
		if !errors.Is(err, exec.ErrDot) {
			log.L().Error("fetch current binary path failed", zap.Error(err))
		}
		log.L().Warn("provider2domainset hasn't been installed to path, let's turn to install it")
		Install()
	}
	if _, err := os.Stat(ServicePath); err == nil {
		err := os.Remove(ServicePath)
		if err != nil {
			log.L().Sugar().Warnf("remove %s failed", ServicePath)
		}
	}
	file, err := os.OpenFile(ServicePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.L().Sugar().With(err).Fatalf("open %s failed", ServicePath)
	}
	defer file.Close()
	if _, err := file.Write(ServiceFile); err != nil {
		log.L().Sugar().With(err).Fatalf("write %s failed", ServicePath)
	}
	log.L().Info("add provider2domainset to init.d success")
}

func RmService() {
	err := os.Remove(ServicePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.L().Error("delete service failed", zap.Error(err))
	}
}
