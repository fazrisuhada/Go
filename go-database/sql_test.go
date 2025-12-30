package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at from customers"

	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("================================")
		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		if birthDate.Valid {
			fmt.Println("birth_date:", birthDate.Time)
		}
		fmt.Println("married:", married)
		fmt.Println("created_at:", createdAt)
	}

}
