package eos_cmd

import (
	"errors"
	"fmt"
	"github.com/eager7/go/mlog"
	"os/exec"
	"regexp"
)

var log = mlog.L

func runCmd(shell string) ([]byte, error) {
	cmd := exec.Command("/bin/bash", "-c", shell)
	out, err := cmd.Output()
	if err != nil {
		log.Warn("exec ", cmd.Args, "failed, err:", err.Error(), "result:", string(out))
		return nil, err
	}
	//log.Notice("exec [", num, "]", cmd.Args, "success")
	return out, err
}

func md5Sum(data string) (string, error){
	md5, err := runCmd(fmt.Sprintf("md5 -s %s", data)) //MD5 ("pct") = 973333fd6d3ab217522f3785d5d3e6ea
	if err != nil {
		return "", err
	}
	reg := regexp.MustCompile(`.*= (.*)`)//(.*)分组匹配多个
	ret := reg.FindSubmatch(md5)
	if len(ret) != 2 {
		return "", errors.New("regexp search error")
	}
	return string(ret[1]), nil
}

func compileCode(code []byte) ([]byte, error) {
	return nil, nil
}
