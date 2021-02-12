package main

import (
	"fmt"
	"time"

	"github.com/willianscsilva/go-querybuilder"
)

// People entity
type People struct {
	Name string
	Age  int
	Date string
}

func main() {
	start := time.Now()
	e := People{
		Name: "John",
		Age:  38,
		Date: "2021-01-01'T'09:08:00",
	}
	var queryList []string
	for i := 0; i < 1000; i++ {
		insertQuery := "insert into database.table (Name, Age, Date) VALUES(:Name, :Age, :Date)"
		buildedInsert := querybuilder.QueryBuilder(insertQuery, e)

		selectQuery := "select * from database.table where age >= :Age and date >= :Date and name = :Name)"
		buildedSelect := querybuilder.QueryBuilder(selectQuery, e)

		selectQuery2 := `select * from database.table where age >= :Age
							and date >= :Date and name = :Name)`
		buildedSelect2 := querybuilder.QueryBuilder(selectQuery2, e)

		insertQuery2 := `INSERT INTO database.table (Name, Age, Date)
		VALUES(:Name, :Age, :Date) RETURNING id`
		buildedInsert2 := querybuilder.QueryBuilder(insertQuery2, e)

		queryList = append(queryList, buildedInsert)
		queryList = append(queryList, buildedSelect)
		queryList = append(queryList, buildedInsert2)
		queryList = append(queryList, buildedSelect2)
	}
	fmt.Println(queryList)
	fmt.Println("List length: ", len(queryList))
	elapsed := time.Since(start)
	fmt.Printf("Execution time %s", elapsed)
}
