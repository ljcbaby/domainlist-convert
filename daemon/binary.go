package daemon

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dn-11/provider2domainset/log"
	"go.uber.org/zap"
)

const installed = "/usr/sbin/provider2domainset"

// Install copy binary to /usr/sbin/ (openwrt path)
func Install() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil && !errors.Is(err, exec.ErrDot) {
		log.Logger.Error("fetch current binary path failed", zap.Error(err))
		return
	}

	absFile, err := filepath.Abs(file)
	if err != nil {
		log.Logger.Error("The absPath failed", zap.Error(err))
		return
	}
	log.Logger.Sugar().Infof("current binary: %v", absFile)

	originFp, err := os.Open(absFile)
	if err != nil {
		log.Logger.Error("open current binary failed", zap.Error(err))
		return
	}
	defer originFp.Close()

	if _, err := os.Stat(installed); err != nil {
		if !os.IsNotExist(err) {
			log.Logger.Error("fetch binary stat failed", zap.Error(err))
			return
		}
	} else {
		if err := os.RemoveAll(installed); err != nil {
			log.Logger.Error("remove old binary failed", zap.Error(err))
			return
		}
	}

	fp, err := os.OpenFile(installed, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Logger.Sugar().Errorf("write to %v", installed)
		return
	}
	defer fp.Close()
	_, err = io.Copy(fp, originFp)
	if err != nil {
		_ = os.RemoveAll(installed)
		log.Logger.Sugar().With(err).Errorf("copy binary to %s", installed)
		return
	}
	log.Logger.Info("installed provider2domainset")
}

func Uninstall() {
	file, err := exec.LookPath("provider2domainset")
	if err != nil {
		log.Logger.Error("find provider2domainset failed", zap.Error(err))
		return
	}

	if err := os.RemoveAll(file); err != nil {
		log.Logger.Error("remove binary failed", zap.Error(err), zap.String("path", file))
		return
	}
}