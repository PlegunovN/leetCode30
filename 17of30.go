// Write a class that allows getting and setting key-value pairs, however a time until expiration is associated with each key.
//
// The class has three public methods:
//
// set(key, value, duration): accepts an integer key, an integer value, and a duration in milliseconds. Once the duration has elapsed, the key should be inaccessible. The method should return true if the same un-expired key already exists and false otherwise. Both the value and duration should be overwritten if the key already exists.
//
// get(key): if an un-expired key exists, it should return the associated value. Otherwise it should return -1.
//
// count(): returns the count of un-expired keys.

package main

import (
	"fmt"
	"time"
)

var mP map[int]int

func init() {
	mP = make(map[int]int)
}

type ExpirableMap struct {
	mP map[int]int
}

func NewExpirableMap() *ExpirableMap {
	return &ExpirableMap{
		mP: make(map[int]int),
	}
}

var expirationTime time.Time

func (m *ExpirableMap) Set(key int, value int, period int) bool {
	if _, ok := m.mP[key]; ok {
		m.mP[key] = value
		return true
	} else {
		expirationTime = time.Now().Add(time.Duration(period) * time.Millisecond)
		m.mP[key] = value
		return false
	}
}

func (m *ExpirableMap) Get(key int) int {
	value, ok := m.mP[key]
	if !ok || time.Since(m.GetExpirationTime(key)) > 0 {
		return value
	} else {
		return -1
	}
}

func (m *ExpirableMap) Count() int {
	count := 0
	for _, value := range m.mP {
		if time.Since(m.GetExpirationTime(value)) > 0 {
			count++
		}
	}
	return count
}

func (m *ExpirableMap) GetExpirationTime(key int) time.Time {
	_, ok := m.mP[key]
	if ok {
		value, _ := m.mP[key]
		expirationTime = time.Unix(0, int64(value)).Add(time.Duration(-1) * time.Second)
		return expirationTime
	} else {
		return time.Time{}
	}
}

func main() {
	m := NewExpirableMap()

	m.Set(1, 100, 500)
	fmt.Println(m.Get(1))
	time.Sleep(100 * time.Millisecond)
	fmt.Println(m.Count())
	m.Set(2, 30, 500)
	fmt.Println(m.Get(2))
	fmt.Println(m.Get(3))
	fmt.Println(m.Count())

	time.Sleep(1000 * time.Millisecond)

	fmt.Println(m.Count())
	fmt.Println(m.Get(1))

}
