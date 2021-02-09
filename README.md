# go-querybuilder
### Build your queries in golang

This library does not implement golang's sql package. It just feeds your query with values received in a struct. Soon it will not be possible to execute the query through this library.

That said, consult the sample package to see the execution performance. For better performance, run the compiled application.

### See an example of how to use
```go
package main
import (
	"github.com/willianscsilva/go-querybuilder"
)
type People struct {
  Name string
  Age  int
  Date string
}
func main() {
  e := People{
    Name: "John",
    Age:  38,
    Date: "2021-01-01'T'00:00:00",
  }
  
  insertQuery := "insert into databse.table (Name, Age, Date) VALUES(:Name, :Age, :Date)"
  buildedInsert := querybuilder.QueryBuilder(insertQuery, e)

  selectQuery := "select * from databse.table where age >= :Age and date >= :Date and name = :Name)"
  buildedSelect := querybuilder.QueryBuilder(selectQuery, e)
}
```

Thank you very much for using.
