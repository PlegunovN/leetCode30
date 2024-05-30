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

type valuePeriod struct {
	value     int
	period    time.Duration
	startTime time.Time
}

type ExpirableMap struct {
	target map[int]valuePeriod
}

func NewExpirableMap() *ExpirableMap {
	return &ExpirableMap{
		target: make(map[int]valuePeriod),
	}
}

func (m *ExpirableMap) Set(key int, value int, period time.Duration, start time.Time) bool {
	start = time.Now()
	if _, ok := m.target[key]; ok {
		m.target[key] = valuePeriod{value: value, period: period, startTime: start}
		return true
	} else {

		m.target[key] = valuePeriod{value: value, period: period, startTime: start}
		return false
	}
}

func (m *ExpirableMap) Get(key int) int {
	period := m.target[key].period
	start := m.target[key].startTime
	value := m.target[key].value
	_, ok := m.target[key]
	if !ok || time.Now().Sub(start) < period {
		fmt.Println(time.Now().Sub(start))
		return value
	} else {
		return -1
	}
}

func (m *ExpirableMap) Count() int {
	count := 0

	for i := 0; i < len(m.target); i++ {
		if m.GetExpirationTime(i) == true {
			count++
		}
	}
	return count
}

func (m *ExpirableMap) GetExpirationTime(key int) bool {
	period := m.target[key].period
	start := m.target[key].startTime
	_, ok := m.target[key]
	if !ok || time.Now().Sub(start) < period {

		return true
	} else {
		return false
	}
}

func main() {
	m := NewExpirableMap()

	m.Set(1, 100, 1*time.Second, time.Now())
	fmt.Println("1 = ", m.Get(1))
	time.Sleep(100 * time.Millisecond)
	fmt.Println("count:", m.Count())
	m.Set(2, 30, 1*time.Second, time.Now())
	time.Sleep(100 * time.Millisecond)
	fmt.Println("2 = ", m.Get(2))
	time.Sleep(100 * time.Millisecond)
	fmt.Println("3 = ", m.Get(3))
	fmt.Println("count:", m.Count())

	time.Sleep(1000 * time.Millisecond)

	fmt.Println("count:", m.Count())
	fmt.Println("1 = ", m.Get(1))

}
