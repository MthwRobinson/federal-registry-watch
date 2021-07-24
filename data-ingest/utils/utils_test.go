package utils

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateIfNotExists(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gotest")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "testdir")
	assert.NoDirExists(t, path)
	createIfNotExists(path)
	assert.DirExists(t, path)
}

func TestCreateDirectory(t *testing.T) {
	dir, _ := ioutil.TempDir("", "gotest")
	defer os.RemoveAll(dir)

	createDirectory(dir, "2021-02-03", 5)
	expectedPath := filepath.Join(dir, "register-files", "2021", "02", "03", "5")
	assert.DirExists(t, expectedPath)
}
