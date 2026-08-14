package cli

import (
	"fmt"
	"os"
	"runtime"

	"github.com/b4moss/slz/internal/mac"
	"github.com/spf13/cobra"
)

const Version = "0.2.0"

type Options struct {
	GOOS   string
	Runner mac.Runner
}

func NewRoot(opts Options) *cobra.Command {
	if opts.GOOS == "" {
		opts.GOOS = runtime.GOOS
	}
	if opts.Runner == nil {
		opts.Runner = mac.NewRunner()
	}

	rootCmd := &cobra.Command{
		Use:     "slz",
		Version: Version,
		Args:    cobra.NoArgs,
		Run:     func(cmd *cobra.Command, args []string) {},
	}
	rootCmd.SetVersionTemplate("slz version {{.Version}}\n")
	rootCmd.Flags().Bool("version", false, "version for slz")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help [command]",
		Short: "Help about any command",
		RunE: func(c *cobra.Command, args []string) error {
			target, leftover, err := c.Root().Find(args)
			if err != nil {
				return err
			}
			if len(leftover) > 0 {
				return fmt.Errorf("unknown help topic %q", args)
			}
			return target.Help()
		},
	})

	rootCmd.AddCommand(newSetupCmd())
	rootCmd.AddCommand(newConfigCmd())
	rootCmd.AddCommand(newMacCmd(opts))
	return rootCmd
}

func Execute() {
	if err := NewRoot(Options{}).Execute(); err != nil {
		os.Exit(1)
	}
}
