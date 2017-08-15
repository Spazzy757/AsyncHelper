package main

import (
	"fmt"
	"gopkg.in/redis.v4"
	"log"
	"net/http"
	"os"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	http.HandleFunc("/", foo)
	fmt.Println("Server Started on Port :8600")
	log.Fatal(http.ListenAndServe(":8600", nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
