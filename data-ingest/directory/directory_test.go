package directory

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateIfNotExists(t *testing.T) {
	dir, _ := ioutil.TempDir("", "fed-registry-test")
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "testdir")
	assert.NoDirExists(t, path)
	createIfNotExists(path)
	assert.DirExists(t, path)
}

func TestCreateDirectoryForDate(t *testing.T) {
	dir, _ := ioutil.TempDir("", "fed-registry-test")
	defer os.RemoveAll(dir)

	CreateDirectoryForDate(dir, "2021-02-03")
	expectedPath := filepath.Join(dir, "register-files", "2021", "02", "03")
	assert.DirExists(t, expectedPath)
}

func TestReadAndWriteJSON(t *testing.T) {
	dir, _ := ioutil.TempDir("", "fed-registry-test")
	defer os.RemoveAll(dir)

	type Animal struct {
		Height int    `json:height`
		Weight int    `json:weight`
		Paws   int    `json:paws`
		Sound  string `json:sound`
	}

	bear := Animal{Height: 8, Weight: 750, Paws: 4, Sound: "Roar!"}
	filename := filepath.Join(dir, "bear.json")
	WriteJSON(bear, filename)

	sameBear := Animal{}
	ReadJSON(&sameBear, filename)
	assert.Equal(t, bear, sameBear)
}
