package cmd

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sfs",
	Short: "sfs is a Static File Server",
	Run: func(cmd *cobra.Command, args []string) {
    e := echo.New()
    e.Static("/", ".")
    port := fmt.Sprintf(":%d", Port)
    e.Logger.Fatal(e.Start(port))
  },
}

var Port int

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
  rootCmd.Flags().IntVarP(&Port, "port", "p", 1323, "specify port")
}
