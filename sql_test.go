package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close() // recomendation for issue connection leak

	ctx := context.Background()

	scriptSql := "INSERT INTO customer(id, name) VALUES('Dandi', 'Dean')"

	// implement exec context *context and query sql, and params is variadic argument (opsional)
	// untuk insert, delete, update
	_, err := db.ExecContext(ctx, scriptSql)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close() // recomendation for issue connection leak

	ctx := context.Background()

	scriptSql := "SELECT id, name FROM customer"

	// implement Query context *context and query sql, and params is variadic argument (opsional)
	// for result query like SELECT query
	rows, err := db.QueryContext(ctx, scriptSql)

	if err != nil {
		panic(err)
	}

	defer rows.Close() // recomendation close for issue result leak

	// (result)-> read data
	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name) // use pointer (&) for resulting data

		if err != nil {
			panic(err)
		}

		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
	}

	defer rows.Close() // recomendation close for result leak
}
