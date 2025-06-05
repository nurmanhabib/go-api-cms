package basepath

import (
	"path/filepath"
	"runtime"
)

func Dir(paths ...string) string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return ""
	}

	root := filepath.Join(filepath.Dir(file), "../..")
	paths = append([]string{root}, paths...)

	return filepath.Join(paths...)
}
