package main

import (
	"idis/cache"
	"idis/http"
)

func main() {
	c := cache.New("inMemoryCache")
    s := http.New(c)
    s.Listen()
}
