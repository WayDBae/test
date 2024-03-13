package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go Redis!")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//TESTING CONNECTION...
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	//ДОБАВЛЕНИЕ ЗНАЧЕНИЯ В СЛОВАРЬ NAME
	err = client.Set(context.Background(), "name", "Daler", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	//ЗАБИРАЕМ ИЗ СЛОВАРЯ NAME ЗНАЧЕНИЕ
	value, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}
