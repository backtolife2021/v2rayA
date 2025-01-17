package where

import (
	"fmt"
	"github.com/v2rayA/v2rayA/global"
	"os/exec"
	"strings"
)

var NotFoundErr = fmt.Errorf("not found")
var ServiceNameList = []string{"v2ray", "xray"}

/* get the version of v2ray-core without 'v' like 4.23.1 */
func GetV2rayServiceVersion() (ver string, err error) {
	v2rayPath, err := GetV2rayBinPath()
	if err != nil || len(v2rayPath) <= 0 {
		return "", newError("cannot find v2ray executable binary")
	}
	out, err := exec.Command("sh", "-c", fmt.Sprintf("%v -version", v2rayPath)).Output()
	var fields []string
	if fields = strings.Fields(strings.TrimSpace(string(out))); len(fields) < 2 {
		return "", newError("cannot parse version of v2ray")
	}
	ver = fields[1]
	if strings.ToUpper(fields[0]) != "V2RAY" {
		ver = "UnknownClient"
	}
	return
}

func GetV2rayBinPath() (string, error) {
	v2rayBinPath := global.GetEnvironmentConfig().V2rayBin
	if v2rayBinPath == "" {
		return getV2rayBinPathAnyway()
	}
	return v2rayBinPath, nil
}

func getV2rayBinPathAnyway() (path string, err error) {
	for _, target := range ServiceNameList {
		if path, err = getV2rayBinPath(target); err == nil {
			return
		}
	}
	return
}

func getV2rayBinPath(target string) (string, error) {
	var pa string
	//从环境变量里找
	out, err := exec.Command("sh", "-c", "which "+target).CombinedOutput()
	if err != nil {
		return "", NotFoundErr
	}
	pa = strings.TrimSpace(string(out))
	if pa == "" {
		return "", NotFoundErr
	}
	return pa, nil
}
