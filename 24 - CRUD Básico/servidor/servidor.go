package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json"email"`
}

// Cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	request, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler requisição!"))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(request, &usuario); err != nil {
		w.Write([]byte("Erro ao adicionar usuário no struct"))
		return
	}

	db, err := banco.Conn()
	if err != nil{
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO usuario (nome, email) VALUES (?, ?)")
	if err != nil{
		w.Write([]byte("Erro ao preparar inserção de dados no banco!"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil{
		w.Write([]byte("Erro ao inserir dados no banco!"))
		return
	}

	id_inserted, err := insert.LastInsertId()
	if err != nil{
		w.Write([]byte("Erro ao retornar Id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", id_inserted)))
}

// Busca e retorna todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){

	db, err := banco.Conn()
	if err != nil{
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	query, err := db.Query("select * from usuario")
	if err != nil{
		w.Write([]byte("Erro ao fazer consulta no banco de dados!"))
		return
	}
	defer query.Close()

	var usuarios []usuario

	for query.Next(){
		var usuario usuario

		if err := query.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil{
			w.Write([]byte("Erro ao escanear usuário no banco!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil{
		w.Write([]byte("Erro ao converter usuário para json"))
		return
	}
}

// Busca e retorna apenas um usuário do banco de dados 
func BuscarUsuario(w http.ResponseWriter, r *http.Request){

	parametro := mux.Vars(r)

	ID, err := strconv.ParseUint(parametro["id"], 10, 32)
	if err != nil{
		w.Write([]byte("Erro ao converter parâmetro"))
		return
	}

	db, err := banco.Conn()
	if err != nil{
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	defer db.Close()

	linha, err := db.Query("select * from usuario where id = ?", ID)
	if err != nil{
		w.Write([]byte("Erro ao buscar usuario"))
		return
	}

	var usuario usuario

	if linha.Next(){
		if err := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil{
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(usuario); err != nil{
		w.Write([]byte("Erro ao converter usuário para JSON"))
		return
	}

}

// Atualiza um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){

	parametro := mux.Vars(r)

	ID, err := strconv.ParseUint(parametro["id"], 10, 32)
	if err != nil{
		w.Write([]byte("Erro ao converter parâmetro para inteiro"))
		return
	}

	request_body, err := io.ReadAll(r.Body)
	if err != nil{
		w.Write([]byte("Erro ao obter body da requisição"))
		return
	}

	var usuario usuario
	if err := json.Unmarshal(request_body, &usuario); err != nil{
		w.Write([]byte("Erro ao converter usuário para struc"))
		return
	}

	db, err := banco.Conn()
	if err != nil{
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuario set nome = ?, email = ? where id = ?")
	if err != nil{
		w.Write([]byte("Erro ao preparar update"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, ID); err != nil{
		w.Write([]byte("Erro ao atualizar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// Deleta um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request){

	parametro := mux.Vars(r)

	ID, err := strconv.ParseUint(parametro["id"], 10, 32)
	if err != nil{
		w.Write([]byte("Erro ao converter parâmetro para inteiro!"))
		return
	}

	db, err := banco.Conn()
	if err != nil{
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("delete from usuario where id = ?")
	if err != nil{
		w.Write([]byte("Erro ao preparar a query!"))
		return
	}

	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil{
		w.Write([]byte("Erro ao deletar usuário do banco de dados!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}