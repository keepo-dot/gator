package main

import (
	"github.com/keepo-dot/gator/internal/config"
	"github.com/keepo-dot/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
