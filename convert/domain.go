package convert

import (
	"bufio"
	"os"
	"strings"

	"github.com/ljcbaby/domainlist-convert/conf"
	"github.com/ljcbaby/domainlist-convert/log"
	"github.com/pkg/errors"
)

func convertDomain(s *os.File, t *os.File) error {
	scanner := bufio.NewScanner(s)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.Split(line, "#")[0]
		line = strings.TrimSpace(line)

		if line == "payload:" {
			continue
		}
		line = strings.TrimPrefix(line, `- '`)
		line = strings.TrimSuffix(line, `'`)

		if len(line) == 0 {
			continue
		}

		res := ""

		if strings.Contains(line, "*") || line[0] == '.' {
			if !conf.Convert.EnableRegex {
				log.L().Sugar().Warnf("regex is disabled: %s", line)
				continue
			}

			if line[0] == '.' {
				if !strings.Contains(line, "*") {
					res = "regexp:^.+" + strings.ReplaceAll(line, ".", "\\.") + "$"
				} else {
					res = "regexp:^.+" + strings.ReplaceAll(strings.ReplaceAll(line, ".", "\\."), "*", "[^.]+") + "$"
					log.L().Sugar().Warnf("complex rule, check carefully: %s -> %s", line, res)
				}
			} else {
				if line[0] != '+' {
					res = "regexp:^" + strings.ReplaceAll(strings.ReplaceAll(line, ".", "\\."), "*", "[^.]+") + "$"
				} else {
					res = "regexp:^.*" + strings.ReplaceAll(strings.ReplaceAll(line[2:], ".", "\\."), "*", "[^.]+") + "$"
					log.L().Sugar().Warnf("complex rule, check carefully: %s -> %s", line, res)
				}
			}
		} else {
			if line[0] == '+' {
				res = "domain:" + line[2:]
			} else {
				res = "full:" + line
			}
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
