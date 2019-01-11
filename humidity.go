package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {

	// flags
	redisHost := flag.String("host", "localhost", "Redis Host")
	redisPort := flag.String("port", "6379", "Redis Port")
	redisPassword := flag.String("password", "", "RedisPassword")
	redisCount := flag.Int("count", 100000, "run this man times")
	redisJitter := flag.Int("jitter", 5000, "run jitter")
	keyName := flag.String("key-name", "humidity", "Redis Key Name")
	role := flag.String("role", "", "humidifier or dehumidifier")
	keywait := flag.Bool("wait", false, "if true wait for key to be created")
	flag.Parse()

	if *role == "" {
		fmt.Println("Please set role to either humidifier or dehumidifier")
		os.Exit(1)
	}

	rClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", *redisHost, *redisPort),
		Password: *redisPassword,
	})

	if *keywait {
		fmt.Println("Waiting for key", *keyName, "before starting")
		for {
			r, _ := rClient.HGetAll(
				*keyName,
			).Result()
			if len(r) > 0 {
				break
			}
			time.Sleep(100 * time.Millisecond)

		}
	}

	for i := 0; i < *redisCount; i++ {

		// use sleep to add some jitter
		if i%*redisJitter == 0 {
			r := rand.Int63n(100)
			time.Sleep(time.Duration(r) * time.Millisecond)
			fmt.Println("Number of iterations:", i)
		}

		// set the role
		_, err := rClient.HMSet(
			*keyName,
			map[string]interface{}{
				"role": *role,
			}).Result()

		if err != nil {
			fmt.Println("error setting key")
			os.Exit(1)
		}

		rClient.HIncrBy(*keyName, fmt.Sprintf("%s_count", *role), 1)

	}
	fmt.Println("Number of iterations:", *redisCount)
}
