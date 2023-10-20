package config

import "github.com/uptrace/bun"

type DBInterface interface {
	NewDBConnection() (*bun.DB, error)
}
