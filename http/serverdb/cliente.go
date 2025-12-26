package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)
	//esse swicth recebe true como parâmetro, por isso não tem parâmetro
	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:root@localhost:5432/curso")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	//Quey várias linhas, QueryRow uma linha
	var u Usuario
	conn.QueryRow(ctx, `Select * from usuarios where id = $1`, id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request, id int) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:root@localhost:5432/curso")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	rows, _ := conn.Query(ctx, "select id, nome from usuarios")

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}
	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
