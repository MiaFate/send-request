package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"send-request/models"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	godotenv.Load(".env")
	bearer := os.Getenv("BEARER")
	url := os.Getenv("URL")

	var messages = &models.Messages{}

	for range 20 {
		var msg models.Message
		product := readFirstLine()

		// product := readFirstLine()
		if product == "" {
			//crea una archivo con el texto "No hay productos"
			// Open file
			f, err := os.Create("finished.txt")
			if err != nil {
				log.Fatalln(err)
			}
			f.WriteString("exit")
			defer f.Close()

		} else {
			msg.Msg = &models.Msg{}
			msg.Msg.ProductId = &product
			msg.Id = &product
			messages.Messages = append(messages.Messages, &msg)
		}
	}
	if messages.Messages == nil {
		return
	}

	payloadBytes, err := json.Marshal(messages)
	if err != nil {
		log.Fatalln(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// read response body
	respBody, error := io.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	var responses models.Responses
	json.Unmarshal(respBody, &responses)
	for _, response := range responses.Responses {
		appendProduct(response.Id, response.Code)
	}

	defer resp.Body.Close()

}

// funcion que lee el archivo, guarde la primera linea en una variable y la elimine del archivo
// luego que guarde el archivo en la misma ubicacion
// y por ultimo que retorne la variable con la primera linea
func readFirstLine() string {
	// Open file
	f, err := os.Open("products.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Create a new scanner and read the file line by line
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	firstLine := scanner.Text()

	// Create a new file and write the rest of the lines
	f, err = os.Create("products.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for scanner.Scan() {
		_, err := f.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}

	return firstLine
}

// funcion que appendea el product en un archivo responsesOK.txt o responsesError.txt dependiendo del status code
// y nada mas
func appendProduct(product string, statusCode int) {
	var f *os.File
	var err error
	if statusCode == 200 {
		f, err = os.OpenFile("responsesOK.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		f, err = os.OpenFile("responsesError.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.WriteString(product + "\n")
	if err != nil {
		log.Fatalln(err)
	}
}
