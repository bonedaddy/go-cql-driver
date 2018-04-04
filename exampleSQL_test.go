package cql_test

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/MichaelS11/go-cql-driver"
)

func Example_cqlSqlSelect() {
	db, err := sql.Open("cql", cql.TestHostValid)
	if err != nil {
		fmt.Printf("Open error is not nil: %v", err)
		return
	}
	if db == nil {
		fmt.Println("db is nil")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	rows, err := db.QueryContext(ctx, "select cql_version from system.local")
	cancel()
	if err != nil {
		fmt.Printf("QueryContext error is not nil: %v", err)
		return
	}
	if !rows.Next() {
		fmt.Println("no row data")
		return
	}
	
	dest := make([]interface{}, 1)
	destPointer := make([]interface{}, 1)
	destPointer[0] = &dest[0]
	err = rows.Scan(destPointer...)
	if err != nil {
		fmt.Printf("Scan error is not nil: %v", err)
		return
	}

	if len(dest) != 1 {
		fmt.Println("len dest != 1")
		return
	}
	data, ok := dest[0].(string)
	if !ok {
		fmt.Println("dest type not string")
		return
	}
	if len(data) < 3 {
		fmt.Println("data string len too small")
		return
	}

	fmt.Println("recived cql_version from system.local")

	// output: recived cql_version from system.local
}