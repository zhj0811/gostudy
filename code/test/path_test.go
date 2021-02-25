package test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	curPath, err := filepath.Abs("./path_test.go")
	assert.NoError(t, err)
	dir := filepath.Dir(curPath)
	base := filepath.Base(dir)
	//filepath.Match()
	//filepath.Walk
	filename := fmt.Sprint(base, "-cert.pem")
	pemPath := filepath.Join(dir, filename)
	t.Logf("pem path: %s", pemPath)
	pattern := fmt.Sprint(dir, "/*.go")
	files, err := filepath.Glob(pattern)
	assert.NoError(t, err)
	t.Logf("files: %#v", files)

}
