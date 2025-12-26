package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	conn, err := pgx.Connect(ctx,
		"postgres://postgres:senha@localhost:5432/curso")

	if err != nil {
		// log.Fatal(err)
		panic(err) //para quando o estado do programa fica inválido
		// Não use panic para:
		// Erros esperados
		// Validação de dados
		// Falhas de conexão normais
		// Fluxo de controle

	}
	defer conn.Close(ctx)

	var id int
	nome := "Paulo"
	err = conn.QueryRow(
		ctx, `INSERT INTO usuarios (nome)
		VALUES ($1)
		RETURNING id`, nome,
	).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Insert realizado com sucesso! Id: %d", id)

}
