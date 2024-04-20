package util

import (
	"path"
	"runtime"
)

func GetProjectDir(relatePath string) string {
	projectDir := ""
	_, filename, _, _ := runtime.Caller(0)
	projectDir = path.Join(path.Dir(filename), relatePath)
	return projectDir
}
