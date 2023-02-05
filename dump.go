package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newDumpCmd(cctx *cmdContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "dump",
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) (rerr error) {
			// cfg := cctx.db

			dumpRoot := args[0]
			basePath := args[1]

			err := filepath.WalkDir(dumpRoot, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				destination, err := filepath.Rel(basePath, path)
				if err != nil {
					return err
				}

				if d.IsDir() {
					log.Printf("mkdir %q", destination)
					err := os.MkdirAll(destination, 0o777)
					if os.IsExist(err) {
						return nil
					} else if err != nil {
						return err
					}
					return nil
				}

				addcmd := []string{"git", "link", "add", path, destination}
				log.Println(addcmd)
				cmd := exec.Command(addcmd[0], addcmd[1:]...)
				output, err := cmd.CombinedOutput()
				log.Println(string(output))
				if err != nil {
					return err
				}

				return nil
			})
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
