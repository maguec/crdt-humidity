package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {

	// flags
	redisHost := flag.String("host", "localhost", "Redis Host")
	redisPort := flag.String("port", "6379", "Redis Port")
	redisPassword := flag.String("password", "", "RedisPassword")
	keyName := flag.String("key-name", "humidity", "Redis Key Name")
	flag.Parse()

	rClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", *redisHost, *redisPort),
		Password: *redisPassword,
	})

	_, err := rClient.HMSet(
		*keyName,
		map[string]interface{}{
			"watcher": 1,
		}).Result()

	if err != nil {
		fmt.Println("error setting key", err)
		os.Exit(1)
	}

	r, _ := rClient.HMGet(
		*keyName,
		"winner",
		"net_humidity",
		"humidifier_count",
		"dehumidifier_count",
	).Result()

	// On the first run these will all be nil so fill in n/a for easier viewing
	for q := range r {
		if r[q] == nil {
			r[q] = "n/a"
		}
	}

	fmt.Printf("%-16s %-16s\n", "key", "value")
	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("%-16s %-16s\n", "lastwriter", r[0])
	fmt.Printf("%-16s %-16s\n", "net humidity", r[1])
	fmt.Printf("---------------------------------------------\n")
	fmt.Printf("%-16s %-16s\n", "humidifier", r[2])
	fmt.Printf("%-16s %-16s\n", "dehumidifier", r[3])

}
