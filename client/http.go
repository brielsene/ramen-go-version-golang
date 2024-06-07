package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"ramen-go/dto"

	"github.com/joho/godotenv"
)

func CreateReqWithExternalAPI() dto.OrderId {
	//CARREGAR ARQUIVOS DO .ENV
	godotenv.Load()
	url := "https://api.tech.redventures.com.br/orders/generate-id"

	client := http.Client{}

	var dto dto.OrderId
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("Erro na requisição: ", err)
	}
	apiKey := os.Getenv("API_KEY")
	req.Header.Set("x-api-key", apiKey)
	fmt.Println("API KEY: ", apiKey)

	// Executa a requisição
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erro ao executar requisição: %v", err)
	}
	defer resp.Body.Close()

	// Lê a resposta
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Erro ao ler resposta: %v", err)
	}

	err = json.Unmarshal(body, &dto)

	// Retorna a resposta ao cliente
	return dto
}
