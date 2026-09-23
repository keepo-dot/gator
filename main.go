package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/keepo-dot/gator/internal/config"
	"github.com/keepo-dot/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %s\n", err)
		os.Exit(1)
	}
	cliState := state{
		cfg: &cfg,
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("error connecting to database: %s\n", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	cliState.db = dbQueries
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	userCmds := os.Args
	if len(userCmds) < 2 {
		fmt.Println("command name required")
		os.Exit(1)
	}
	cmd := command{
		name:      userCmds[1],
		arguments: userCmds[2:],
	}
	err = cmds.run(&cliState, cmd)
	if err != nil {
		fmt.Printf("error running %s command: %s\n", cmd.name, err)
		os.Exit(1)
	}
}
