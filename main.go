package main 

import "fmt"
import "net/http"
import "golang.org/x/net/html"
import "strings"
import "io/ioutil"

func main() {
	fmt.Printf("[INFO] Getting Web Page\n")
	resp, err := http.Get("http://db.desmoinesregister.com/state-salaries-for-iowa/page=1")
	if err != nil {
		fmt.Println(err)
		return
	} else {
	defer resp.Body.Close()
	//doc, err := ioutil.ReadAll(resp.Body)
	data, err := html.Parse(resp)
	if err != nil {
		fmt.Println(err)
	} 
	fmt.Println("[INFO] Parsing out data")
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "td" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(data)
	}
}