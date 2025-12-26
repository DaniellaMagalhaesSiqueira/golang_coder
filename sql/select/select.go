package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:senha@localhost:5432/curso")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	//Quey várias linhas, QueryRow uma linha
	rows, error := conn.Query(ctx, `Select * from usuarios where id < 10`)
	if error != nil {
		log.Fatal(error)
	}
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
