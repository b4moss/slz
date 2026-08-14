package cli

import (
	"os"

	"github.com/spf13/cobra"
)

const Version = "0.1.0"

func Execute() {
	rootCmd := &cobra.Command{
		Use:     "slz",
		Version: Version,
	}
	rootCmd.SetVersionTemplate("slz version {{.Version}}\n")
	rootCmd.Flags().Bool("version", false, "version for slz")
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
