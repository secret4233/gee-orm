package main

import (
	"fmt"

	"geeorm"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string
}

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()

	userNow := User{"Tom"}
	if err := s.Find(&userNow); err != nil {
		fmt.Printf("Error!!")
	}
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
