package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "114.67.80.197:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// _, err = c.Do("SET", "mykey", "superWang")
	_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	// 判断key是否存在
    is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey1"))
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Printf("exists or not: %v \n", is_key_exit)
    }

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
