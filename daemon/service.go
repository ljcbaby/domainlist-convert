package daemon

import (
	_ "embed"
	"errors"

	"github.com/dn-11/provider2domainset/log"
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
			log.Logger.Error("fetch current binary path failed", zap.Error(err))
		}
		log.Logger.Warn("provider2domainset hasn't been installed to path, let's turn to install it")
		Install()
	}
	if _, err := os.Stat(ServicePath); err == nil {
		err := os.Remove(ServicePath)
		if err != nil {
			log.Logger.Sugar().Warnf("remove %s failed", ServicePath)
		}
	}
	file, err := os.OpenFile(ServicePath, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Logger.Sugar().With(err).Fatalf("open %s failed", ServicePath)
	}
	defer file.Close()
	if _, err := file.Write(ServiceFile); err != nil {
		log.Logger.Sugar().With(err).Fatalf("write %s failed", ServicePath)
	}
	log.Logger.Info("add provider2domainset to init.d success")
}

func RmService() {
	err := os.Remove(ServicePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Logger.Error("delete service failed", zap.Error(err))
	}
}
