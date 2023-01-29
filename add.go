package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	xxhash "github.com/cespare/xxhash/v2"

	"github.com/spf13/cobra"
)

func newAddCmd(cctx *cmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use: "add",
		RunE: func(cmd *cobra.Command, args []string) (rerr error) {
			cfg := cctx.db

			destination := args[0]
			source := filepath.Base(destination)

            log.Printf("Calculing hash of %q", destination)
			h := xxhash.New()
			f, err := os.Open(destination)
			if err != nil {
				return err
			}
			if _, err := io.Copy(h, f); err != nil {
				return err
			}

			err = os.Symlink(destination, source)
			if err != nil {
				return err
			}

			defer func() {
				if rerr != nil {
					err := os.Remove(source)
					rerr = errors.Join(err, rerr)
				}
			}()

			cfg.Data.Links = append(cfg.Data.Links, Link{
				Source:      source,
				Destination: args[0],
				XXH64:       fmt.Sprintf("%016x", h.Sum64()),
			})

			err = cfg.Save()
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
