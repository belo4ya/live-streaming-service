package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

const (
	rootDir        = "../third_party"
	swaggerVersion = "v4.15.0"
	swaggerGitUrl  = "https://github.com/swagger-api/swagger-ui.git"
	swaggerDir     = "swagger-ui"
	cacheDir       = ".cache"
	openapiDir     = "../api"
)

var swaggerStaticCmd = &cobra.Command{
	Use: "swagger-static",
	Run: func(cmd *cobra.Command, args []string) {
		tmpDir, _ := os.MkdirTemp("", "")
		defer func() { _ = os.RemoveAll(tmpDir) }()

		gitClone := exec.Command(
			"git",
			"clone",
			"--depth",
			"1",
			"--branch",
			swaggerVersion,
			swaggerGitUrl,
			tmpDir,
		)
		log.Println(gitClone)
		if err := gitClone.Run(); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(swaggerStaticCmd)
}
