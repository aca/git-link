package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	xxhash "github.com/cespare/xxhash/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/cobra"
)

func newFSCKCommand(cctx *cmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "fsck",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := cctx.db

			targetFile, err := filepath.Rel(cctx.rootPath, filepath.Join(cctx.currentPath, args[0]))
			if err != nil {
				return err
			}

			found := false
			for _, link := range cfg.Data.Links {
				if link.Source == targetFile {
					found = true

					destination := link.Destination
					log.Printf("Calculing hash of %q", destination)
					h := xxhash.New()
					f, err := os.Open(destination)
					if err != nil {
						return err
					}
					if _, err := io.Copy(h, f); err != nil {
						return err
					}

					hashString := fmt.Sprintf("%016x", h.Sum64())

					diffString := cmp.Diff(hashString, link.XXH64)
					if diffString == "" {
						log.Printf("%q OK\n", link.Source)
						return nil
					} else {
						return fmt.Errorf("checksum mismatched for %q, %s", link.Source, diffString)
					}
				}
			}

			if !found {
				return fmt.Errorf("%q not found", targetFile)
			}

			return nil
		},
	}
	return cmd
}
