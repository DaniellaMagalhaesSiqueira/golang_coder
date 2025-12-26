package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

// Conectar em postgres

// Criar o banco se não existir

// Fechar a conexão

// Abrir nova conexão apontando para curso

// Criar tabelas

func main() {
	ctx := context.Background()

	// Conecta no banco padrão "postgres"
	conn, err := pgx.Connect(ctx,
		"postgres://postgres:senha@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx) // defer vai fechar no final da função

	dbName := "curso"

	// Verifica se o banco já existe
	var exists bool
	err = conn.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT 1
			FROM pg_database
			WHERE datname = $1
		)
	`, dbName).Scan(&exists)

	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		_, err = conn.Exec(ctx, "CREATE DATABASE "+dbName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Banco criado:", dbName)
	} else {
		fmt.Println("Banco já existe:", dbName)
	}

	conn.Close(ctx)
	appConn, err := pgx.Connect(ctx,
		"postgres://postgres:senha@localhost:5432/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer appConn.Close(ctx)
	// 5️⃣ Criação da tabela
	_, err = appConn.Exec(ctx, `
		DROP TABLE IF EXISTS usuarios;
		CREATE TABLE usuarios (
			id SERIAL PRIMARY KEY,
			nome VARCHAR(80)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tabela usuarios criada com sucesso")
}
