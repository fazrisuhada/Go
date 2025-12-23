package godatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExcSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customers (id, name) VALUES (1, 'Fazri')"

	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}
