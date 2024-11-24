package convert

import "github.com/ljcbaby/domainlist-convert/conf"

type Task struct {
	conf.File
	Source string
	Target string
}
