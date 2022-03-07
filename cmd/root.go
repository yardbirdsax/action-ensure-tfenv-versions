/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yardbirdsax/ensure-tfenv-versions/pkg/exec"
	"github.com/yardbirdsax/ensure-tfenv-versions/pkg/files"
	"github.com/yardbirdsax/ensure-tfenv-versions/pkg/tfenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootDirectory string
	debugFlag     bool
	rootCmd       = &cobra.Command{
		Use:   "ensure-tfenv-versions",
		Short: "Ensures that required versions of Terraform are installed using the tfenv tool.",
		Long:  ``,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		RunE:          runMe,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func runMe(cmd *cobra.Command, args []string) (err error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if debugFlag {
		config.Level.SetLevel(zap.DebugLevel)
	} else {
		config.Level.SetLevel(zap.InfoLevel)
	}
	logger, err := config.Build()
	if err != nil {
		zap.Error(err)
		os.Exit(1)
	}
	zap.ReplaceGlobals(logger)

	zap.S().Infof("Starting at directory %s", rootDirectory)
	versionFiles, err := files.FindFiles("\\.terraform-version", rootDirectory)
	if err != nil {
		zap.S().Error(err)
		return
	}
	zap.S().Infof("Found %d total files", len(versionFiles))
	fileContents, err := files.ReadFiles(versionFiles)
	if err != nil {
		zap.S().Error(err)
		return
	}
	uniqueVersions := tfenv.GetUniqueVersions(fileContents)
	executor := exec.NewExecutor()
	err = tfenv.InstallTFEnvVersions(uniqueVersions, executor)
	if err != nil {
		zap.S().Error(err)
	}
	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootDirectory, "root-directory", "d", ".", "directory at which to being searching for .terraform-version files.")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "verbose", "v", false, "enable verbose logging")
}
