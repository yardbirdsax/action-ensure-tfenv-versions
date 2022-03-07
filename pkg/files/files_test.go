package files

import (
	"io/fs"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	savedAppFs afero.Fs
)

func setUp() {
	savedAppFs = appFs
	appFs = afero.NewMemMapFs()
}

func tearDown() {
	appFs = savedAppFs
}
func TestReadFile(t *testing.T) {
	setUp()
	defer tearDown()

	content := "hello\nworld"
	expectedContent := []string{
		"hello",
		"world",
	}
	fileName := "filename.txt"
	err := afero.WriteFile(appFs, fileName, []byte(content), fs.ModeAppend)
	require.NoError(t, err)

	actualContent, err := ReadFile(fileName)
	assert.Equal(t, expectedContent, actualContent)
	assert.NoError(t, err)
}

func TestReadFiles(t *testing.T) {
	setUp()
	defer tearDown()

	files := map[string]string{
		"file1.txt": "hello\nworld",
		"file2.txt": "hello\nthere",
	}
	var fileNames []string
	expectedContent := []string{
		"hello",
		"world",
		"hello",
		"there",
	}
	for fileName, content := range files {
		fileNames = append(fileNames, fileName)
		err := afero.WriteFile(appFs, fileName, []byte(content), fs.ModeAppend)
		require.NoError(t, err)
	}

	actualContent, err := ReadFiles(fileNames)
	assert.Equal(t, expectedContent, actualContent)
	assert.NoError(t, err)
}

func TestFindFiles(t *testing.T) {
	setUp()
	defer tearDown()

	fileNames := []string{
		"test/one.txt",
		"test/two.txt",
		"test/three.txt",
		"test/one.log",
		"test2/one.txt",
	}
	expectedFileNames := []string{
		"test/one.txt",
		"test/three.txt",
		"test/two.txt",
		"test2/one.txt",
	}
	content := "hello world"
	for _, fileName := range fileNames {
		err := afero.WriteFile(appFs, fileName, []byte(content), fs.ModeAppend)
		require.NoError(t, err)
	}

	actualFileNames, err := FindFiles(".*\\.txt", ".")
	assert.NoError(t, err)
	assert.Equal(t, expectedFileNames, actualFileNames)
}
