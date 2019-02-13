package getme

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://SECRETS.SHH"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Vault-Token", "NOTYOURAPIKEY")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
