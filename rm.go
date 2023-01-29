package main

import (
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// TODO: handle command run in sub-directory
func newRMCommand(cctx *cmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use: "rm",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := cctx.db

            cfg.Data.Links = lo.Filter(cfg.Data.Links, func(l Link, _ int) bool {
                return l.Source == args[0]
            })

            return cfg.Save()
		},
	}
	return cmd
}
