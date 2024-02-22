package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ernest_k/defi_stats/lib"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

const LAST_PROCESSED_BLOCK_KEY = "last_processed_block"

func WriteLastProcessedBlock(chainName string, lastBlockNumber int64) {
	err := godotenv.Load()
	lib.Check(err)

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"), // Replace with your Redis server address
		Password: "",                         // No password for local Redis, provide if needed
		DB:       0,                          // Default DB
	})

	defer rdb.Close()

	err = rdb.Set(context.Background(), collectionName(chainName), lastBlockNumber, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetLastProcessedBlock(latestBlockInChain int64, chainName string) int64 {
	err := godotenv.Load()
	lib.Check(err)

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"), // Replace with your Redis server address
		Password: "",                         // No password for local Redis, provide if needed
		DB:       0,                          // Default DB
	})

	val, err := rdb.Get(context.Background(), collectionName(chainName)).Result()
	if err == redis.Nil {
		// Key does not exist
		WriteLastProcessedBlock(chainName, latestBlockInChain)
		return latestBlockInChain - 10_000
	} else if err != nil {
		// Some other error occurred
		lib.CheckWithMessage(err, "cannot connect to redis")
		return 0
	} else {
		// Key exists, and val contains the value
		intNumber, err := strconv.ParseInt(val, 10, 64)
		lib.CheckWithMessage(err, "cannot convert to int64")

		return intNumber
	}
}

func collectionName(chainName string) string {
	return fmt.Sprintf("chain_name:%s:%s", chainName, LAST_PROCESSED_BLOCK_KEY)
}
