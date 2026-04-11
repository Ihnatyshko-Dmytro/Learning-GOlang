package coincap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	url := "https://rest.coincap.io/v3/assets/bitcoin"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

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

	var r AssetResponse
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.Asset.Info())
	// for _, asset := range r.Data {
	// 	fmt.Println(asset.Info())
	// }
}
