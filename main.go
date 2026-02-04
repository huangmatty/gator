package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/huangmatty/gator/internal/config"
	"github.com/huangmatty/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to open connection to database at %s: %w", cfg.DBUrl, err))
	}
	dbQueries := database.New(db)

	progState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	cmds := commands{
		commandsToHandlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerListUserFeeds))

	if len(os.Args) < 2 {
		log.Fatal("usage: gator <command> [args...]")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	if err := cmds.run(progState, cmd); err != nil {
		log.Fatal(err)
	}
}
