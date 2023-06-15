package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	//postgres://user:password@host:port/dbname

	databaseurl := "postgres://postgres:admin@localhost:5432/profiling"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseurl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database : %v\n", err)
		os.Exit(1)
	}

	fmt.Println("successfully connect to database")
}
