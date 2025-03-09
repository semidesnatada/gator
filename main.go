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
	// fmt.Println("Hello world")
	// con, err := config.Read()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return 
	// }
	// fmt.Println()
	// fmt.Println(con.CurrentUserName)
	// fmt.Println(con.DBUrl)

	// setErr := con.SetUser("derek pickles")
	// if setErr != nil {
	// 	fmt.Println(setErr.Error())
	// 	fmt.Println("the beans are in the tin")
	// 	return
	// }
	// con2, err2 := config.Read()
	// if err2 != nil {
	// 	fmt.Println(err2.Error())
	// 	return 
	// }
	// // fmt.Println()
	// fmt.Println(con2.CurrentUserName)
	// fmt.Println(con2.DBUrl)



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
