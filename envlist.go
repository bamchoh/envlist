package envlist

import (
	"strings"
	"syscall"

	win "golang.org/x/sys/windows/registry"
)

var systemENV = `SYSTEM\CurrentControlSet\Control\Session Manager\Environment`
var cuENV = `Environment`

func splitKeyValue(environment string) (key, value string) {
	e := strings.Split(environment, "=")
	return e[0], e[1]
}

func getEnvList(key win.Key, path string) (map[string]string, error) {
	envList := make(map[string]string, 0)
	k, err := win.OpenKey(key, path, win.QUERY_VALUE)
	if err != nil {
		return envList, err
	}
	defer k.Close()
	keys, err := k.ReadValueNames(0)
	if err != nil {
		return envList, err
	}
	for _, key := range keys {
		val, _, err := k.GetStringValue(key)
		if err != nil {
			return envList, err
		}
		envList[key] = val
	}
	return envList, nil
}

func GetAllEnvList() map[string]string {
	envList := make(map[string]string, 0)
	for i, kv := range syscall.Environ() {
		if i == 0 {
			continue
		}
		k, v := splitKeyValue(kv)
		envList[k] = v
	}
	return envList
}

func GetCurrentUserEnvList() (map[string]string, error) {
	return getEnvList(win.LOCAL_MACHINE, systemENV)
}

func GetSystemEnvList() (map[string]string, error) {
	return getEnvList(win.CURRENT_USER, cuENV)
}
