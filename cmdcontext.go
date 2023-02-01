package main

import "github.com/aca/x/jsondb"

type cmdContext struct {
    db *jsondb.DB[Config]
    rootPath string
    currentPath string
}
