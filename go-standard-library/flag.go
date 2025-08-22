package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		username = flag.String("username", "", "database username")
		password = flag.String("password", "", "database password")
		host     = flag.String("host", "root", "database root")
		port     = flag.Int("port", 0, "database port")
	)

	flag.Parse()

	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
}

//  go run flag.go --username=fazrisuhada --password="Rahasia Banget" --host="123.134.2.2" --port=5000
