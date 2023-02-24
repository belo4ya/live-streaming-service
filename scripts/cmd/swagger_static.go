package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var swaggerStaticCmd = &cobra.Command{
	Use: "swagger-static",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("swaggerStatic called")
	},
}

func init() {
	rootCmd.AddCommand(swaggerStaticCmd)
}
