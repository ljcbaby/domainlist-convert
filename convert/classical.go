package convert

import (
	"bufio"
	"os"
	"strings"

	"github.com/dn-11/provider2domainset/conf"
	"github.com/dn-11/provider2domainset/log"
	"github.com/pkg/errors"
)

func convertClassical(s *os.File, t *os.File) error {
	scanner := bufio.NewScanner(s)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.Split(line, "#")[0]
		line = strings.TrimSpace(line)

		if line == "payload:" {
			continue
		}
		line = strings.TrimPrefix(line, "- ")

		if len(line) == 0 {
			continue
		}

		args := strings.Split(line, ",")
		if len(args) < 2 {
			log.L().Sugar().Warnf("invalid line: %s", line)
		}

		res := ""

		switch args[0] {
		case "DOMAIN":
			res = "full:" + args[1]
		case "DOMAIN-SUFFIX":
			res = "domain:" + args[1]
		case "DOMAIN-KEYWORD":
			res = "keyword:" + args[1]
		case "DOMAIN-REGEX":
			if conf.Convert.EnableRegex {
				res = "regex:" + args[1]
			} else {
				log.L().Sugar().Warnf("regex is disabled: %s", line)
				continue
			}
		default:
			log.L().Sugar().Warnf("unknown type: %s", line)
			continue
		}

		log.L().Sugar().Debugf("convert line: %s -> %s", line, res)

		if _, err := t.WriteString(res + "\n"); err != nil {
			return errors.Wrap(err, "write target file failed")
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "scan source file failed")
	}

	return nil
}
