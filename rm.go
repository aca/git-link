package main

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

// TODO: handle command run in sub-directory
func newRMCommand(cctx *cmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use: "rm",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := cctx.db

			targetFile := args[0]

			relpath, err := filepath.Rel(cctx.rootPath, filepath.Join(cctx.currentPath, targetFile))
			_ = relpath
			if err != nil {
				return err
			}

			// cfg.Data.Links = lo.Filter(cfg.Data.Links, func(l Link, _ int) bool {
			//              if l.Source == relpath {
			//                  log.Printf("removed %q", targetFile)
			//              }
			// 	return l.Source != relpath
			// })

			return cfg.Save()
		},
	}
	return cmd
}
