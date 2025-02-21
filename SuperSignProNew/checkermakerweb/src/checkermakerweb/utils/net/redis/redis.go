/*Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
 */
package redis

import (
	"errors"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	SavePath string
	Poolsize int
	Password string
	DbNum    int
	Poollist *redis.Pool
}

var MAX_POOL_SIZE = 100

func (rp *Redis) Connect() error {
	configs := strings.Split(rp.SavePath, ",")
	if len(configs) > 0 {
		rp.SavePath = configs[0]
	}
	if len(configs) > 1 {
		poolsize, err := strconv.Atoi(configs[1])
		if err != nil || poolsize <= 0 {
			rp.Poolsize = MAX_POOL_SIZE
		} else {
			rp.Poolsize = poolsize
		}
	} else {
		rp.Poolsize = MAX_POOL_SIZE
	}
	if len(configs) > 2 {
		rp.Password = configs[2]
	}
	if len(configs) > 3 {

		dbnum, err := strconv.Atoi(configs[3])
		if err != nil || dbnum < 0 {
			rp.DbNum = 0
		} else {
			rp.DbNum = dbnum
		}

	} else {
		rp.DbNum = 0
	}
	rp.Poollist = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", rp.SavePath)
		if err != nil {
			return nil, err
		}
		if rp.Password != "" {
			if _, err := c.Do("AUTH", rp.Password); err != nil {
				c.Close()
				return nil, err
			}
		}
		_, err = c.Do("SELECT", rp.DbNum)
		if err != nil {
			c.Close()
			return nil, err
		}
		return c, err
	}, rp.Poolsize)

	return rp.Poollist.Get().Err()
}

func (rp *Redis) Exists(key string) bool {
	c := rp.Poollist.Get()
	defer c.Close()

	if existed, err := redis.Int(c.Do("EXISTS", key)); err != nil || existed == 0 {
		return false
	} else {
		return true
	}
}

func (rp *Redis) Get(key string) (val string, err error) {
	c := rp.Poollist.Get()
	defer c.Close()
	val, err = redis.String(c.Do("GET", key))
	if err != nil {
		return
	}
	return
}
func (rp *Redis) Set(key string, val string) (err error) {
	c := rp.Poollist.Get()
	defer c.Close()

	data, err := c.Do("SET", key, val)
	if err != nil {
		return
	}

	if data != "OK" {
		err = errors.New("Data Not Set")
		return
	}
	return
}
