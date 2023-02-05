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
		Use:           "add",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) (rerr error) {
			cfg := cctx.db

			// source -> destination
			// a.log -> /mnt/nas/a.log
			var destination string
			var source string
			if len(args) == 1 {
				destination = args[0]
				source = filepath.Base(destination)
			} else if len(args) == 2 {
				destination = args[0]
				source = args[1]
			} else {
				return fmt.Errorf("invalid argument count: %v", len(args))
			}

			f, err := os.Open(destination)
			if err != nil {
				return err
			}
			defer f.Close()

			stat, err := f.Stat()
			if err != nil {
				return err
			}

			if stat.IsDir() {
				return errors.New("destination should be file")
			}

			log.Printf("Calculing hash of %q", destination)
			h := xxhash.New()

			err = os.Symlink(destination, source)
			if os.IsExist(err) {
				log.Printf("Skip %q: %v", source, err)
				return nil
			} else if err != nil {
				return err
			}

			defer func() {
				if rerr != nil {
					err := os.Remove(source)
					rerr = errors.Join(err, rerr)
				}
			}()

			if _, err := io.Copy(h, f); err != nil {
				return err
			}

			relpath, err := filepath.Rel(cctx.rootPath, filepath.Join(cctx.currentPath, source))
			if err != nil {
				return err
			}

			hashv := fmt.Sprintf("%016x", h.Sum64())

			if hashv != "" {
				return fmt.Errorf("failed to get hash value")
			}

			cfg.Data.Links = append(cfg.Data.Links, Link{
				Source:      relpath,
				Destination: args[0],
				XXH64:       hashv,
				Size:        stat.Size(),
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
