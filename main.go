package main

import (
	"fmt"

	"github.com/semidesnatada/gator/internal/config"
)

func main() {
	// fmt.Println("Hello world")
	con, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	// fmt.Println()
	// fmt.Println(con.CurrentUserName)
	// fmt.Println(con.DBUrl)

	setErr := con.SetUser("derek pickles")
	if setErr != nil {
		fmt.Println(setErr.Error())
		fmt.Println("the beans are in the tin")
		return
	}
	con2, err2 := config.Read()
	if err2 != nil {
		fmt.Println(err2.Error())
		return 
	}
	// fmt.Println()
	fmt.Println(con2.CurrentUserName)
	fmt.Println(con2.DBUrl)

}
