package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/semidesnatada/gator/internal/commands"
	"github.com/semidesnatada/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	s, err := commands.InitialiseState()
	if err != nil {
		fmt.Println(err.Error())
	}

	db, err := sql.Open("postgres", s.Config.DBUrl)

	dbQueries := database.New(db)

	s.DB = dbQueries

	cmds := commands.GetCommands()

	args := os.Args
	if len(args) < 2 {
		fmt.Println("error: not a valid command")
        os.Exit(1)
	} else {
		fmt.Printf("args are: %v and %v\n", args[1], args[2:])
		runErr := cmds.Run(&s, commands.Command{Name:args[1], Argument: args[2:]})
		if runErr != nil {
			fmt.Printf("error: %v\n", runErr)
			os.Exit(1)
		}
	}

}
