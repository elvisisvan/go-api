import (
	"encoding/json"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Face string
	Length int
}

func GetCatFact() {
	url := "https://catfact.ninja/fact"
	var catFat CatFact
	err := GetJson(url, catFact)
	if err != nil {
		fmt.Printf("error getting cat fact: %s\n", err.Error())
		return
	}
}

func GetJson (url string, target interface{} error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
})

func main() {
	client = &http.Client{Timeout: 10 * time.Second})
}