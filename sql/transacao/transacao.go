package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx,
		"postgres://postgres:senha@localhost:5432/curso")
	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// Se o commit não acontecer, faz rollback
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()
	/*
			Mais seguro, se der panic
			defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback(ctx)
				panic(p)
			}
			if err != nil {
				_ = tx.Rollback(ctx)
			}
		}()
	*/

	var id int

	id_erro := 2

	nome := "Paulo"

	err = tx.QueryRow(
		ctx,
		`INSERT INTO usuarios (nome)
	 VALUES ($1, $2)
	 RETURNING id`,
		id_erro,
		nome,
	).Scan(&id)

	if err != nil {
		log.Fatal(err)
		return // dispara o rollback no defer
	}
	err = tx.Commit(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
