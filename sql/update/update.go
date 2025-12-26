package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:root@localhost:5432/curso")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	id := 2
	nome := "Vitor"

	cmdTag, err := conn.Exec(
		ctx,
		`UPDATE usuarios SET nome = $1 WHERE id = $2`,
		nome,
		id,
	)
	if err != nil {
		log.Fatal(err)
	}

	if cmdTag.RowsAffected() == 0 {
		log.Fatal("usuário não encontrado")
	}

	log.Printf("Usuário %d atualizado com sucesso", id)

	id = 3
	cmdTag2, err2 := conn.Exec(
		ctx,
		`DELETE FROM usuarios WHERE id = $1`,
		id,
	)
	if err2 != nil {
		log.Fatal(err)
	}

	if cmdTag2.RowsAffected() == 0 {
		log.Fatal("usuário não encontrado")
	}
	log.Printf("Usuário %d apagado", id)
}
