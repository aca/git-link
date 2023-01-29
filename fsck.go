package main

import (
	"fmt"
	"io"
	"log"
	"os"

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

			found := false
			for _, link := range cfg.Data.Links {
				if link.Source == args[0] {
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
						return nil
					} else {
						return fmt.Errorf("checksum mismatched for %q, %s", link.Source, diffString)
					}
				}
			}

			if !found {
				return fmt.Errorf("%q not found", args[0])
			}

			return nil
		},
	}
	return cmd
}
