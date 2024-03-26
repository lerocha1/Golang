package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Cliente representa a estrutura de dados de um cliente
type Cliente struct {
	Nome      string `json:"nome"`
	Endereco  string `json:"endereco"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Apelido   string `json:"apelido"`
	Descricao string `json:"descricao"`
}

var clientes []Cliente
var arquivoClientes = "clientes.json"

// Handler para lidar com requisições POST para adicionar um cliente
func adicionarCliente(w http.ResponseWriter, r *http.Request) {
	var novoCliente Cliente
	err := json.NewDecoder(r.Body).Decode(&novoCliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clientes = append(clientes, novoCliente)
	salvarClientes()

	log.Printf("Cliente adicionado: %+v\n", novoCliente)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoCliente)
}

// Handler para lidar com requisições GET para obter todos os clientes
func obterClientes(w http.ResponseWriter, r *http.Request) {
	log.Println("Solicitado GET para /clientes")

	json.NewEncoder(w).Encode(clientes)
}

// Handler para lidar com requisições DELETE para remover um cliente
func removerCliente(w http.ResponseWriter, r *http.Request) {
	var clienteRemover Cliente
	err := json.NewDecoder(r.Body).Decode(&clienteRemover)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, cliente := range clientes {
		if cliente.Nome == clienteRemover.Nome {
			clientes = append(clientes[:i], clientes[i+1:]...)
			salvarClientes()

			log.Printf("Cliente removido: %+v\n", clienteRemover)
			break
		}
	}

	// Envie uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Cliente removido: %s\n", clienteRemover.Nome)
}

// Salva os clientes no arquivo JSON
func salvarClientes() {
	dadosJSON, err := json.MarshalIndent(clientes, "", "  ")
	if err != nil {
		log.Fatal("Erro ao converter dados para JSON: ", err)
	}

	err = ioutil.WriteFile(arquivoClientes, dadosJSON, 0644)
	if err != nil {
		log.Fatal("Erro ao escrever no arquivo JSON: ", err)
	}
}

// Carrega os clientes do arquivo JSON
func carregarClientes() {
	arquivo, err := os.Open(arquivoClientes)
	if err != nil {
		log.Println("Arquivo de clientes não encontrado. Será criado um novo.")
		return
	}
	defer arquivo.Close()

	decoder := json.NewDecoder(arquivo)
	if err := decoder.Decode(&clientes); err != nil {
		log.Fatal("Erro ao decodificar arquivo JSON: ", err)
	}
}

func main() {
	carregarClientes()

	// Rota para a API de clientes
	http.HandleFunc("/clientes", obterClientes)
	http.HandleFunc("/adicionar-cliente", adicionarCliente)
	http.HandleFunc("/remover-cliente", removerCliente)

	// Rota para servir o arquivo HTML
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
