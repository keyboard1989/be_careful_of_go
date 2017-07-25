package main

//锁用的不正确经常会出现如下两个问题
//panic1: concurrent map iteration and map write
//panic2: Unlock of unlocked RWMutex

//正确的做法
//写入修改的时候
// u.lock.Lock
// defer u.lock.Unlock()
//读取的时候
// u.lock.RLock
// defer u.lock.RUnlock()

//错误的做法1
//读写都不加锁
//concurrent map iteration and map write

//错误做法2
//读加锁，写不加锁，或反过来
//concurrent map iteration and map write

//错误做法3
// u.lock.Rlock()
// defer u.lock.Unlock()
// Unlock of unlocked RWMutex

import (
	"sync"
	"time"
)

type User struct {
	data map[string]string
	lock sync.RWMutex
}

func (u *User) Set(key string, val string) {
	u.lock.Lock()
	defer u.lock.Unlock()

	u.data[key] = val
}

func (u *User) Get(key string) map[string]string {
	u.lock.RLock()
	defer u.lock.RUnlock()

	d := map[string]string{}
	for key, val := range u.data {
		d[key] = val
	}
	return d
}

func main() {
	u := &User{
		data: map[string]string{},
		lock: sync.RWMutex{},
	}

	go func() {
		for {
			u.Set("hello", "world")
		}
	}()

	go func() {
		for {
			u.Get("hello,world")
		}
	}()

	for {
		time.Sleep(time.Hour)
	}
}
