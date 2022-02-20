/**
Definindo e utilizando structs
**/

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Definições das structs
type Pais struct {
	Nome   string `json:"nome"`   // tags para mudar o nome da chave no json
	Codigo int32  `json:"codigo"` // tags para mudar o nome da chave no json
}

type Cidade struct {
	Nome   string `json:"nome"`   // tags para mudar o nome da chave no json
	Codigo int32  `json:"codigo"` // tags para mudar o nome da chave no json
	Pais   Pais
}

type Cliente struct {
	Nome   string `json:"nome"`   // tags para mudar o nome da chave no json
	Email  string `json:"email"`  // tags para mudar o nome da chave no json
	Cidade Cidade `json:"cidade"` // tags para mudar o nome da chave no json
}

func (c Cliente) toString() {
	fmt.Printf("Cliente: %s, cidade %s, pais: %s", c.Nome, c.Cidade.Nome, c.Nome)
	fmt.Println()
}

func main() {

	cliente := Cliente{
		Nome:  "Joaozinho",
		Email: "joaozinho@google.com",
		Cidade: Cidade{
			Nome:   "Cascavel",
			Codigo: 4104808,
			Pais:   Pais{Nome: "Brasil", Codigo: 55},
		},
	}

	//Gerar Json a partir de uma struct
	clienteJson, err := json.Marshal(cliente)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Imprimindo array de bytes
	fmt.Println("Impressão de bytes")
	fmt.Println(clienteJson)

	//Imprimindo json texto
	fmt.Println("Impressao do Json")
	fmt.Println(string(clienteJson))

	//Preenchar uma variavel de uma origem JSON
	textJson := `{"nome":"Joaozinho","email":"joaozinho@google.com","cidade":{"nome":"Cascavel","codigo":4104808,"Pais":{"nome":"Brasil","codigo":55}}}`
	clienteFromJson := Cliente{}

	json.Unmarshal([]byte(textJson), &clienteFromJson) // aqui a variavel clienteFromJson deve ser passada como referencia usando o &

	fmt.Println("Objeto preenchido a partir de um json")
	fmt.Println(clienteFromJson)
}
