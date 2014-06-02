package main

import (
	"fmt"
	"menteslibres.net/gosexy/redis"
	"os"
	"strconv"
)

var redisCli *redis.Client

func main() {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	portStr := os.Getenv("REDIS_PORT")
	if portStr == "" {
		portStr = "6379"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		os.Exit(1)
	}
	redisCli = redis.New()
	err = redisCli.Connect(host, uint(port))
	if err != nil {
		os.Exit(1)
	}
	defer redisCli.Quit()

	cmd := os.Args[1]
	switch cmd {
	case "get":
		value, err := redisCli.Get(os.Args[2])
		if err == nil {
			fmt.Println(value)
		}
	case "set":
		redisCli.Set(os.Args[2], os.Args[3])
	case "setex":
		args := os.Args[3]
		seconds, err := strconv.Atoi(args)
		if err == nil {
			redisCli.SetEx(os.Args[2], int64(seconds), os.Args[4])
		}
	case "del":
		redisCli.Del(os.Args[2])
	case "incr":
		value, err := redisCli.Incr(os.Args[2])
		if err == nil {
			fmt.Println(value)
		}
	case "decr":
		value, err := redisCli.Decr(os.Args[2])
		if err == nil {
			fmt.Println(value)
		}
	case "publish":
		redisCli.Publish(os.Args[2], os.Args[3])
	}
}
