//@Description Redis服务器地址consistent
//@Contact czw@outlook.com

package XYLibs

import (
	"stathat.com/c/consistent"
	"github.com/garyburd/redigo/redis"
	"strings"
	"errors"
	"time"
	"strconv"
	"sync"
)

const ERROR_REDIS_NOT_CONFIG = "redis not config"


func NewRedisHash() *RedisHash {
	rdb := new(RedisHash)
	rdb.redisIPList = consistent.New()
	rdb.exeLock = new(sync.Mutex)
	return rdb
}

type RedisHash struct {
	
	redisIPList *consistent.Consistent
 	redisConnPool map[string]*redis.Pool 
	arrIP []string
	pool   *redis.Pool
	conn   redis.Conn
	HashID 	string
	exeLock  *sync.Mutex
}


func (c *RedisHash)ConnectRedis(ipaddr string) {
	c.arrIP = strings.Split(ipaddr,",")
	l := len(c.arrIP)
	if l == 0 {
		panic(ERROR_REDIS_NOT_CONFIG)
	}
	c.redisConnPool = make(map[string]*redis.Pool,l)
	for i,v := range c.arrIP {
		index := strconv.Itoa(i)
		c.redisIPList.Add(index)
		c.redisConnPool[index] = c.newRedisConnPool(v)
	}
}

func (c *RedisHash) newRedisConnPool(server string) *redis.Pool {
    return &redis.Pool{
        MaxIdle: 3,
		MaxActive: 10000,
        IdleTimeout: 180 * time.Second,
        Dial: func () (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
			println("conn redis...")
            if err != nil {
                return nil, err
            }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
}

func (c *RedisHash) connRedis() error {
	
	c.getRedisIpByHashID()
	if c.pool == nil {
		return errors.New("redis conn pool is null")
	}
	c.conn = c.pool.Get()
	return nil
}
//放回连接池
func (c *RedisHash) closeRedis() error {
	
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *RedisHash) INCR(key string) error {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return  res
	}
	defer c.closeRedis()
	_, err := c.conn.Do("INCR", key)
	return err
}

func (c *RedisHash) PipeliningINCR(key []interface{}) error {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return  res
	}
	defer c.closeRedis()
	for _,v := range key {
		 c.conn.Send("INCR", v)
	}
	err :=	c.conn.Flush()

	return err
}

func (c *RedisHash) SETEX(key string, expire int , val interface{}) error {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return  res
	}
	defer c.closeRedis()
	_, err := c.conn.Do("SETEX", key,expire, val)
	return err
}

func (c *RedisHash) Get(key string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("GET",  key)

}

func (c *RedisHash) Del(key string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("DEL",  key)

}


func (c *RedisHash) HMSET(key ,field , value string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("HMSET", key, field , value)
}

func (c *RedisHash) MSET(key ,  value string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("MSET", key, value)

}

func (c *RedisHash) MSETByte(key string,  value []byte) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("MSET", key, value)

}

func (c *RedisHash) MGET(key []interface{}) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("MGET", key...)

}

func (c *RedisHash) SADD(key ,  value string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SADD", key, value)

}

func (c *RedisHash) SADDByte(key string,  value []byte) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SADD", key, value)

}



func (c *RedisHash) SREM(key ,  value string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SREM", key, value)

}

func (c *RedisHash) SREMByte(key string,  value []byte) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SREM", key, value)

}

func (c *RedisHash) SMEMBERS(key string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SMEMBERS", key)

}

func (c *RedisHash) SISMEMBER(key , member string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SISMEMBER", key , member)

}

func (c *RedisHash) SDIFF(key , diffKey string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SDIFF", key , diffKey)

}


func (c *RedisHash) ZADD(key , score, member string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZADD", key, score, member)

}


func (c *RedisHash) ZREVRANGE(key ,  start, stop string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZREVRANGE",key,start,stop)

}

func (c *RedisHash) ZREVRANK(key ,  member string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
		return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZREVRANK",key,member)

}

func (c *RedisHash) ZCOUNT(key ,  min , max string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
			return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZCOUNT",key,min,max)

}


func (c *RedisHash) ZCARD(key string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
			return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZCARD",key)

}


func (c *RedisHash) SRANDMEMBER(key , count string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
			return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("SRANDMEMBER",key,count)

}

func (c *RedisHash) ZSCORE(key ,  member  string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
			return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZSCORE",key,member)

}

func (c *RedisHash) ZREM(key ,  member  string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
			return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZREM",key,member)

}

func (c *RedisHash) ZINCRBY(key , increment, member  string) (interface{}, error) {
	c.exeLock.Lock()
	defer c.exeLock.Unlock()
	res := c.connRedis()
	if res != nil {
			return nil, res
	}
	defer c.closeRedis()
	return c.conn.Do("ZINCRBY",key,increment,member)

}


func (c *RedisHash) getRedisIpByHashID() {	
	index, err := c.redisIPList.Get(c.HashID)	
	if err != nil {
		c.pool =  nil
	}
	if p,ok := c.redisConnPool[index]; ok {
		c.pool = p
	}else {
		c.pool = nil
	}
	
	
}

func (c *RedisHash) GetRedisIpAll() []string {
	return c.arrIP
}

