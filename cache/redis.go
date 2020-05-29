package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var redisCli *redis.Pool

func Init() {
	redisCli = createPool()
	return
}

func createPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 50,
		// max number of connections
		MaxActive: 1000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// ping tests connectivity for redis (PONG should be returned)
func Ping() error {
	// Send PING command to Redis
	// PING command returns a Redis "Simple String"
	// Use redis.String to convert the interface type to string

	// get a connection from the pool (redis.Conn)
	// use defer to close the connection when the function completes
	conn := redisCli.Get()
	defer conn.Close()

	s, err := redis.String(conn.Do("PING"))
	if err != nil {
		return err
	}

	fmt.Printf("PING Response = %s\n", s)
	// Output: PONG

	return nil
}

// set executes the redis SET command
func Set() error {

	conn := redisCli.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "Favorite Movie", "Repo Man")
	if err != nil {
		return err
	}

	return nil
}

// get executes the redis GET command
func Get() error {
	// Simple GET example with String helper
	conn := redisCli.Get()

	key := "Favorite Movie"
	s, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return (err)
	}
	fmt.Printf("%s = %s\n", key, s)

	return nil
}
