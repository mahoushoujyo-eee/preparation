package dao

import (
	"errors"
	"math/rand"
	"time"
	"w2/weekend/entity"
)

var fakeUsers = []entity.MockUser{
	{ID: 1, Name: "alice"},
	{ID: 2, Name: "bob"},
	{ID: 3, Name: "charlie"},
}

func QueryAllUsers() ([]entity.MockUser, error) {
	time.Sleep(10 * time.Millisecond)
	return fakeUsers, nil
}

// 模拟一次探查，比如SELECT 1
func CheckDBCon() error {
	time.Sleep(5 * time.Millisecond) // 模拟一点 IO 耗时
	if rand.Intn(100) < 30 {
		return errors.New("db connection lost")
	}
	return nil
}
