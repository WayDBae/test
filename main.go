package main

import (
	"context"
	"dockedis/entities"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Go Redis!")
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//TESTING CONNECTION...
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	//ДОБАВЛЕНИЕ ЗНАЧЕНИЯ В СЛОВАРЬ NAME
	err = client.Set(ctx, "name", "Daler", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	//ЗАБИРАЕМ ИЗ СЛОВАРЯ NAME ЗНАЧЕНИЕ
	value, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

	json, err := json.Marshal(entities.User{
		Name: "Daler",
		Age:  1,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(json)

	err = client.Set(ctx, "id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Set(ctx, "id1234", json, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
