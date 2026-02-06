package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/CybrRonin/gator/internal/config"
	"github.com/CybrRonin/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	//fmt.Printf("Read config: %v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connectign to database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	currState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)

	commandArgs := os.Args

	if len(commandArgs) < 2 {
		log.Fatal("Not enough arugments to proceed")
	}

	cmdName := commandArgs[1]
	args := commandArgs[2:]

	cmd := command{Name: cmdName, Args: args}

	err = cmds.run(currState, cmd)
	if err != nil {
		log.Fatal(err)
	}
	/*
	   cfg, err = config.Read()

	   	if err != nil {
	   		log.Fatalf("error reading config: %v", err)
	   	}

	   fmt.Printf("Read config again: %v\n", cfg)
	*/
}
