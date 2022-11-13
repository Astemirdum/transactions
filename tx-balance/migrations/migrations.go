package migrations

import "embed"

//go:embed *.sql
var MigrationFiles embed.FS

const Label = "transaction"
