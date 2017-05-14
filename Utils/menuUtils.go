package Utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Menu struct {
	Key string
	Url string
}

func initMenu() []Menu {

	bytes, err := ioutil.ReadFile("menu.json")
	if err != nil {
		log.Fatal(err)
	}

	var menus []Menu
	if err := json.Unmarshal(bytes, &menus); err != nil {
		log.Fatal(err)
	}

	return menus

}

func GetMenuKeys() []string {
	menus := initMenu()

	var menuKeys []string

	for _, item := range menus {
		menuKeys = append(menuKeys, item.Key)
	}

	return menuKeys
}

func GetCodeUrl(key string) string {

	menus := initMenu()

	var url string

	for _, item := range menus {
		if item.Key == key {
			url = item.Url
		}
	}

	return url
}