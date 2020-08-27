package main

import (
	"bobac/cache"
	"bobac/http"
)

func main() {
	c := cache.New("inMemoryCache")
    s := http.New(c)
    s.Listen()
}
