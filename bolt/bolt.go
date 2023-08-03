package main

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
)

func main() {
	path := "/tmp/bolt.db"

	db, err := bolt.Open(path, 0666, nil) // ①
	if err != nil {
		fmt.Println("open db failed:", path, err)
		return
	}
	defer db.Close()

	tx, err := db.Begin(true) // ②
	if err != nil {
		fmt.Println("begin trans failed:", err)
		return
	}
	defer tx.Rollback()

	_ = tx.DeleteBucket([]byte("numbers"))

	bucket, err := tx.CreateBucketIfNotExists([]byte("numbers")) // ③
	if err != nil {
		fmt.Println("create buckets failed")
		return
	}

	for i := 0; i < 256; i++ {
		_ = bucket.Put([]byte(strconv.Itoa(i)), []byte(strconv.Itoa(i))) // ④
	}

	val := bucket.Get([]byte("128")) // ⑤
	fmt.Println(string(val))

	tx.Commit() // ⑥
	return
}
