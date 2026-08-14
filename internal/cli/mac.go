package cli

import (
	"github.com/b4moss/slz/internal/mac"
	"github.com/spf13/cobra"
)

func newMacCmd(opts Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mac",
		Short: "macOS helpers",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommand(&cobra.Command{
		Use:   "dim",
		Short: "Dim the display without sleeping",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return mac.Dim(opts.GOOS, opts.Runner)
		},
	})
	return cmd
}
