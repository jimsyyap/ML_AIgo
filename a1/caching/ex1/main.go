package main

import (
    "fmt"
    "time"
    cache "github.com/patrickmn/go-cache"
)

func main() {
    c := cache.New(5*time.Minute, 10*time.Minute)
    c.Set("mykey", "myvalue", cache.DefaultExpiration)
    v, found := c.Get("mykey")
    if found {
        fmt.Println(v)
    }
}
