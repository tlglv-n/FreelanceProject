package market

import "github.com/redis/go-redis/v9"

// redis://username:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2

type Redis struct {
	Connection *redis.Client
}

func NewRedis(url string) {

}
