package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Integralist/go-findroot/find"
	"github.com/aca/x/jsondb"
	"github.com/spf13/cobra"
)

func cmdMain() error {
	log.SetFlags(0)

	rootCmd, _ := newRootCmd(os.Stdout, os.Args[1:])
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func newRootCmd(out io.Writer, args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "git-link",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	f := cmd.PersistentFlags()
	f.Parse(args)

	cctx := &cmdContext{}

	root, err := find.Repo()
	if err != nil {
		return nil, err
	}

    cctx.currentPath, err = os.Getwd()
	if err != nil {
		return nil, err
	}

    cctx.rootPath = root.Path

	cfg, err := jsondb.Open[Config](filepath.Join(cctx.rootPath, ".gitlinks"))
	if err != nil {
		return nil, err
	}

	if cfg.Data.Version == "" {
		cfg.Data.Version = version
	}

	cctx.db = cfg

	cmd.AddCommand(
		newAddCmd(cctx),
		newFSCKCommand(cctx),
		newRMCommand(cctx),
	)

	return cmd, nil
}
