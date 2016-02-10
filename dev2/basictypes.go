package main

import (
	"net/url"
	"net/http"
	"github.com/nsavageJVM/rest"
)

type Item struct { }

func (item Item) Get(values url.Values, headers http.Header) (int, interface{}, http.Header) {
	items := []string{"item1", "item2"}
	data := map[string][]string{"items": items}
	return 200, data, http.Header{"Content-type": {"application/json"}}
}


func main() {

	item := new(Item)
	api := rest.NewAPI()
	api.AddResource(item, "/items")
	api.BootStrapServer(8080)

}
