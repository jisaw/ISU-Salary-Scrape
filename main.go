package main 

import "fmt"
import "net/http"

import "io/ioutil"

func main() {
	fmt.Printf("[INFO] Getting Web Page\n")
	resp, err := http.Get("http://db.desmoinesregister.com/state-salaries-for-iowa/page=1")
	if err != nil {
		fmt.Println(err)
		return
	} else {
	defer resp.Body.Close()
	doc, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("%s\n", string(doc))
	}
}