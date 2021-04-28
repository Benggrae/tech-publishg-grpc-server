package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetHtml(url string) {
	res, _ := http.Get(url)

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}
