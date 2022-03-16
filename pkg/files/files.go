package files

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/afero"
	"go.uber.org/zap"
)

var (
	appFs = afero.NewOsFs()
)

func FindFiles(pattern string, path string) (files []string, err error) {
	err = afero.Walk(appFs, path, func(path string, info os.FileInfo, inErr error) (outErr error) {
		if inErr != nil {
			outErr = inErr
			return
		}
		if matched, outErr := regexp.MatchString(pattern, info.Name()); matched && outErr == nil {
			files = append(files, path)
		}
		return
	})
	return
}

func ReadFile(path string) (contents []string, err error) {
	byteContent, err := afero.ReadFile(appFs, path)
	if err != nil {
		return
	}
	contents = strings.Split(string(byteContent), "\n")
	return
}

func ReadFiles(paths []string) (contents []string, err error) {
	for _, path := range paths {
		zap.S().Debugf("Reading file %s", path)
		content, fileErr := ReadFile(path)
		if err != nil {
			zap.S().Errorf("Error reading file '%s': %v", path, fileErr)
			err = fmt.Errorf("errors encountered reading one or more files, please see output")
			continue
		}
		contents = append(contents, content...)
	}
	return
}
