package tfenv

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/yardbirdsax/ensure-tfenv-versions/pkg/exec"

	"go.uber.org/zap"
)

func GetUniqueVersions(versions []string) (uniqueVersions []string) {
	versionMap := map[string]bool{}

	for _, version := range versions {
		if _, exists := versionMap[version]; !exists {
			uniqueVersions = append(uniqueVersions, version)
			versionMap[version] = true
		}
	}

	return
}

func isTFEnvVersionInstalledE(version string, exec exec.Exec) (isInstalled bool, err error) {

	output, err := exec.ExecCommand("tfenv", false, "list")
	if err != nil {
		return
	}
	outputSplit := strings.Split(output, "\n")

	for _, line := range outputSplit {
		if matched, _ := regexp.Match(version, []byte(line)); matched {
			isInstalled = true
		}
	}

	return
}

func findInStringList(stringToMatch string, list []string) bool {
	for _, s := range list {
		if matched, _ := regexp.Match(stringToMatch, []byte(s)); matched {
			return true
		}
	}
	return false
}

func uninstallTFEnvVersion(version string, exec exec.Exec) (err error) {
	isInstalled, err := isTFEnvVersionInstalledE(version, exec)
	if err != nil {
		return
	}
	if isInstalled {
		_, err = exec.ExecCommand("tfenv", true, "uninstall", version)
	}
	return
}

func installTFEnvVersion(version string, exec exec.Exec) (err error) {
	zap.S().Debugf("Checking if version '%s' is currently installed", version)
	isInstalled, err := isTFEnvVersionInstalledE(version, exec)
	if err != nil {
		return
	}
	if !isInstalled {
		zap.S().Infof("Installing Terraform version '%s'", version)
		_, err = exec.ExecCommand("tfenv", true, "install", version)
	} else {
		zap.S().Debugf("Version '%s' is already installed", version)
	}
	return
}

func InstallTFEnvVersions(versions []string, exec exec.Exec) (err error) {
	for _, version := range versions {
		versionErr := installTFEnvVersion(version, exec)
		if versionErr != nil {
			zap.S().Errorf("Error installing version '%s': %v", version, versionErr)
			err = fmt.Errorf("error installing one or more versions, please see previous output")
			continue
		}
	}
	return
}
