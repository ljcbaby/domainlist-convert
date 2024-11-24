package daemon

import (
	"os/exec"
	"sync"
	"time"

	"github.com/dn-11/provider2domainset/log"
)

var (
	refreshChan = make(chan struct{}, 1)
	tLock       sync.Mutex
	timer       *time.Timer
)

func delayRestart() {
	for range refreshChan {
		tLock.Lock()
		if timer == nil {
			log.L().Info("Timer start")
			timer = time.NewTimer(60 * time.Second)
			go func() {
				<-timer.C
				log.L().Info("Restart mosdns")
				cmd := exec.Command("/etc/init.d/mosdns", "restart")
				err := cmd.Run()
				if err != nil {
					log.L().Sugar().With(err).Error("Restart mosdns failed")
				}
				tLock.Lock()
				timer = nil
				tLock.Unlock()
			}()
		} else {
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(60 * time.Second)
		}
		tLock.Unlock()
	}
}
