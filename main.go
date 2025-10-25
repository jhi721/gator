package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jhi721/gator/internal/config"
	"github.com/jhi721/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	config  *config.Config
	conn    *sql.DB
	queries *database.Queries
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	queries := database.New(db)

	s := &state{
		config:  &cfg,
		conn:    db,
		queries: queries,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("users", handlerUsers)

	commands.register("reset", handlerReset)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		log.Fatal("not enough args were provided")
	}

	cmdName := cliArgs[1]
	cmdArgs := cliArgs[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	if err := commands.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
