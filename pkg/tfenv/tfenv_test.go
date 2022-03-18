package tfenv

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yardbirdsax/ensure-tfenv-versions/mocks"
)

func TestFindInStringList_Found(t *testing.T) {
	t.Parallel()
	list := []string{
		"hello",
		"world",
	}
	stringToMatch := "hello"

	found := findInStringList(stringToMatch, list)

	assert.True(t, found)
}

func TestIsTFEnvVersionInstalledE_Installed(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockExec := mocks.NewMockExec(ctrl)
	expectedArgs := []interface{}{
		"list",
	}
	version := "0.15.5"
	mockExec.EXPECT().ExecCommand("tfenv", false, expectedArgs...).Times(1).Return(
		version,
		nil,
	)

	result, err := isTFEnvVersionInstalledE(version, mockExec)
	assert.True(t, result)
	assert.NoError(t, err)
}

func TestIsTFEnvVersionInstalledE_NotInstalled(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	mockExec := mocks.NewMockExec(ctrl)
	expectedArgs := []interface{}{
		"list",
	}
	version := "0.15.5"
	actualVersion := "0.15.6"
	mockExec.EXPECT().ExecCommand("tfenv", false, expectedArgs...).Times(1).Return(
		actualVersion,
		nil,
	)

	result, err := isTFEnvVersionInstalledE(version, mockExec)
	assert.False(t, result)
	assert.NoError(t, err)
}

func TestUninstallTFEnvVersion_Present(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockExec := mocks.NewMockExec(ctrl)

	version := "0.15.5"

	mockExec.EXPECT().ExecCommand("tfenv", false, "list").Times(1).Return(
		version,
		nil,
	)
	mockExec.EXPECT().ExecCommand("tfenv", true, "uninstall", version).Times(1).Return("", nil)

	err := uninstallTFEnvVersion(version, mockExec)
	assert.NoError(t, err)
}

func TestUninstallTFEnvVersion_NotPresent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockExec := mocks.NewMockExec(ctrl)

	version := "0.15.5"

	mockExec.EXPECT().ExecCommand("tfenv", false, "list").Times(1).Return(
		"",
		nil,
	)
	mockExec.EXPECT().ExecCommand("tfenv", true, "uninstall", version).Times(0).Return("", nil)

	err := uninstallTFEnvVersion(version, mockExec)
	assert.NoError(t, err)
}

func TestInstallTFEnvVersion_Present(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockExec := mocks.NewMockExec(ctrl)

	version := "0.15.5"

	mockExec.EXPECT().ExecCommand("tfenv", true, "install", version).Times(1).Return("", nil)

	err := installTFEnvVersion(version, mockExec)
	assert.NoError(t, err)
}

func TestInstallTFEnvVersion_NotPresent(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockExec := mocks.NewMockExec(ctrl)

	version := "0.15.5"

	mockExec.EXPECT().ExecCommand("tfenv", true, "install", version).Times(1).Return("", nil)

	err := installTFEnvVersion(version, mockExec)
	assert.NoError(t, err)
}

func TestGetUniqueVersions(t *testing.T) {
	t.Parallel()

	inputVersions := []string{
		"0.15.5",
		"0.15.5",
		"1.0.1",
		"1.1.0",
	}
	expectedVersions := []string{
		"0.15.5",
		"1.0.1",
		"1.1.0",
	}

	actualVersions := GetUniqueVersions(inputVersions)
	assert.Equal(t, expectedVersions, actualVersions)
}
