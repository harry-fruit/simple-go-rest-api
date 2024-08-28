package config

import (
	"os"
	"strings"
)

func GetArgs() map[string]string {
	args := os.Args

	kvMap := make(map[string]string)

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			parts := strings.SplitN(arg[2:], "=", 2)
			if len(parts) == 2 {
				kvMap[parts[0]] = parts[1]
			}
		}
	}

	return kvMap
}
