package main

import (
	"fmt"
	"testing"
	"time"
)

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Truncate(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

func Test(t *testing.T) {
	today := time.Now()
	yesterday := time.Now().Add(-24 * time.Hour)
	jamTutupBuku := yesterday.Add(1439 * time.Minute)

	fmt.Println("Today : ", today)
	fmt.Println("yesterday :", yesterday)
	fmt.Println("Tutup Buku : ", jamTutupBuku)
}

func Test2(t *testing.T) {
	now := time.Now()
	fmt.Println(Bod(now))
	fmt.Println(Truncate(now))
	fmt.Println("Tutup Buku : ", Bod(now).Add(-1*time.Second))
}
