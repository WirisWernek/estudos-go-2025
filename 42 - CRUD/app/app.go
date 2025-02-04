package app

import (
	"crud/banco"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Insert é utilizado para criar usuários
func Insert(w http.ResponseWriter, r *http.Request) {

	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o body da request"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(requestBody, &usuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter usuário"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios(nome, email) VALUES(?, ?)")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao setar parametros do statement"))
		return
	}

	idUsuario, erro := insercao.LastInsertId()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar o id do usuário inserido"))
		return
	}

	// versão mais usada
	// w.WriteHeader(201)
	// w.Write([]byte(fmt.Sprintf("Usuário cadastrado com sucesso! ID: %d", idUsuario)))

	usuario.ID = uint32(idUsuario)
	usuarioJSON, erro := json.Marshal(usuario)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter idUsuario em JSON"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(usuarioJSON)

}

// GetAll retorna todos os usuários
func GetAll(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar usuarios no banco"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear os usuários"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os usuários em JSON"))
		return
	}
}

// GetById retorna o usuário ao qual o id passado pertence
func GetById(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// a função ParseUint recebe 3 parâmetros: 1º a variavel a ser convertida, 2º a base utilizada, 3º o tamanho em bits
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parÂmetro ID"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar usuarios no banco"))
		return
	}

	defer linha.Close()

	var usuario usuario

	if linha.Next() {

		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o usuário em JSON"))
		return
	}
}

// Update atualiza os dados de um usuário
func Update(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parÂmetro ID"))
		return
	}

	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o body da request"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(requestBody, &usuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter usuário"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao setar parametros do statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete apaga um usuário
func Delete(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parÂmetro ID"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao setar parametros do statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
