package main

import (
	"fmt"

	"github.com/willianscsilva/go-querybuilder"
)

// People entity
type People struct {
	Name string
	Age  int
	Date string
}

func main() {
	e := People{
		Name: "Will",
		Age:  38,
		Date: "2021-01-01'T'09:08:00",
	}
	var queryList []string
	for i := 0; i < 1000; i++ {
		insertQuery := "insert into databse.table (Name, Age, Date) VALUES(:Name, :Age, :Date)"
		buildedInsert := querybuilder.QueryBuilder(insertQuery, e)

		selectQuery := "select * from databse.table where age >= :Age and date >= :Date and name = :Name)"
		buildedSelect := querybuilder.QueryBuilder(selectQuery, e)

		queryList = append(queryList, buildedInsert)
		queryList = append(queryList, buildedSelect)
	}
	fmt.Println(queryList)
	fmt.Println("List length: ", len(queryList))

}
