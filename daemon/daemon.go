package daemon

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/dn-11/provider2domainset/conf"
	"github.com/dn-11/provider2domainset/convert"
	"github.com/dn-11/provider2domainset/log"
	"github.com/fsnotify/fsnotify"
)

func Serve() {
	go delayRestart()
	start()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.L().Sugar().Fatalf("Create watcher failed: %v", err)
	}

	defer func() {
		err := watcher.Close()
		if err != nil {
			log.L().Sugar().Errorf("Close watcher failed: %v", err)
		}
	}()

	err = watcher.Add(conf.Convert.Source)
	if err != nil {
		log.L().Sugar().Errorf("Add watch failed: %v", err)
	}

	go watch(watcher)

	select {}
}

func watch(watcher *fsnotify.Watcher) {
	var sLock sync.Mutex
	lastEvent := make(map[string]time.Time)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				for _, f := range conf.Convert.ProcessFiles {
					if event.Name == filepath.Join(conf.Convert.Source, f.Name) {
						sLock.Lock()
						now := time.Now()

						if last, ok := lastEvent[f.Name]; ok && now.Sub(last) < time.Second {
							sLock.Unlock()
							continue
						}

						lastEvent[f.Name] = now
						sLock.Unlock()

						log.L().Sugar().Infof("File %s modified", f.Name)
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
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}

			log.L().Sugar().Errorf("Watcher error: %v", err)
		}
	}
}
