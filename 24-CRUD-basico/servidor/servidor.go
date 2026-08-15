package servidor

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"go_aprenda_do_zero/24-CRUD-basico/banco"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responderErro(w, http.StatusBadRequest, "Falha ao ler o corpo da requisição:", erro.Error())
		return
	}
	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		responderErro(w, http.StatusBadRequest, "Erro ao converter o usuário para struct:", erro.Error())
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao conectar no banco de dados:", erro.Error())
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao criar o statement:", erro.Error())
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao inserir o usuário:", erro.Error())
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao obter o ID do usuário:", erro.Error())
		return
	}

	var responseIdInserido struct {
		ID int64 `json:"id"`
	}
	responseIdInserido.ID = idInserido
	fmt.Println(responseIdInserido)
	responderJSON(w, http.StatusCreated, responseIdInserido)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao conectar com o banco:", erro.Error())
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar os usuários:", erro.Error())
		return
	}
	defer linhas.Close()
	var usuarios []usuario

	for linhas.Next() {
		if erro := linhas.Err(); erro != nil {
			responderErro(w, http.StatusInternalServerError, "Erro durante a leitura dos usuários:", erro.Error())
			return
		}
		var usuario usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			responderErro(w, http.StatusInternalServerError, "Erro ao escanear o usuário:", erro.Error())
			return
		}
		usuarios = append(usuarios, usuario)
	}

	responderJSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		responderErro(w, http.StatusBadRequest, "Formato do id incorreto:", erro.Error())
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao conectar com o banco:", erro.Error())
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao buscar o usuário:", erro.Error())
		return
	}
	defer linha.Close()

	var usuario usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			responderErro(w, http.StatusInternalServerError, "Erro ao escanear o usuário:", erro.Error())
			return
		}
	} else {
		responderErro(w, http.StatusNotFound, "ID de usuário não encontrado: ", parametros["id"])
		return
	}

	responderJSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		responderErro(w, http.StatusBadRequest, "Formato do id incorreto: ", erro.Error())
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		responderErro(w, http.StatusBadRequest, "Falha ao ler o corpo da requisição: ", erro.Error())
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		responderErro(w, http.StatusBadRequest, "Erro ao converter o usuário para struct: ", erro.Error())
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao conectar com o banco: ", erro.Error())
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")

	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao criar o statement: ", erro.Error())
		return
	}
	defer statement.Close()

	qtdeUpdates, erro := statement.Exec(usuario.Nome, usuario.Email, ID)
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao atualizar o usuário: ", erro.Error())
		return
	}

	// Captura a quantidade de linhas alteradas
	linhasAfetadas, erro := qtdeUpdates.RowsAffected()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao obter linhas afetadas: ", erro.Error())
		return
	}

	// Se nenhuma linha foi alterada, o ID provavelmente não existe no banco
	if linhasAfetadas == 0 {
		responderErro(w, http.StatusNotFound, "Nenhum usuário foi encontrado com o ID informado ou os dados já estão atualizados. ", "")
		return
	}

	var respondeIdAtualizado struct {
		ID uint64 `json:"id"`
	}
	respondeIdAtualizado.ID = ID
	responderJSON(w, http.StatusOK, respondeIdAtualizado)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		responderErro(w, http.StatusBadRequest, "Formato do id incorreto:", erro.Error())
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao conectar com o banco:", erro.Error())
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao criar o statement: ", erro.Error())
		return
	}
	defer statement.Close()

	qtdeUpdates, erro := statement.Exec(ID)
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao excluir o usuário: ", erro.Error())
		return
	}

	// Captura a quantidade de linhas alteradas
	linhasAfetadas, erro := qtdeUpdates.RowsAffected()
	if erro != nil {
		responderErro(w, http.StatusInternalServerError, "Erro ao obter linhas afetadas: ", erro.Error())
		return
	}

	// Se nenhuma linha foi alterada, o ID provavelmente não existe no banco
	if linhasAfetadas == 0 {
		responderErro(w, http.StatusNotFound, "Nenhum usuário foi encontrado com o ID informado. ", parametros["id"])
		return
	}

	var respondeIdAtualizado struct {
		ID uint64 `json:"id"`
	}
	respondeIdAtualizado.ID = ID
	responderJSON(w, http.StatusOK, respondeIdAtualizado)
}

// ResponderJSON envia qualquer dado formatado como JSON com o status HTTP desejado
func responderJSON(w http.ResponseWriter, status int, dados any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			responderErro(w, http.StatusInternalServerError, "Erro ao gerar o json do usuário:", erro.Error())
		}
	}
}

// ResponderErro envia uma mensagem de erro padronizada em formato JSON
func responderErro(w http.ResponseWriter, status int, mensagem string, erro string) {
	mensagem = fmt.Sprintf(mensagem+"%s", erro)
	responderJSON(w, status, map[string]string{
		"erro": mensagem,
	})
}
