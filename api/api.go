package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	// 1. Формуємо URL (зверніть увагу: без фігурних дужок)
	url := "https://rest.coincap.io/v3/price/bysymbol/BTC"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Додаємо заголовки ТОЧНО як у документації
	// Важливо: Bearer [ПРОБІЛ] Ключ
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer 1qaz@WSX3edc")

	// 3. Виконуємо запит
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}