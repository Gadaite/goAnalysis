package DataSource

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/patrickmn/go-cache"
	"log"
	"os"
	"time"
)

var MemoryCacheObject *cache.Cache = nil
var BoltCacheDB *bolt.DB = nil

func Cache2Memory(key string, value any) {

	MemoryCacheObject.Set(key, value, cache.DefaultExpiration)
}

func LoadMemoryCache(key string) any {
	res, _ := MemoryCacheObject.Get(key)
	return res
}

func DeleteMemoryCache(key string) {
	MemoryCacheObject.Delete(key)
}

func InitCacheMemory(expiration time.Duration, expired time.Duration) {
	c := cache.New(expiration, expired)
	MemoryCacheObject = c
}

func LoadBucketCache(bucketName string) {
	if err := BoltCacheDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("key:%s ,value:%s\n", k, v)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func Cache2Bucket(key string, value string, bucketName string) {
	if err := BoltCacheDB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			log.Fatal(err)
		}
		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			log.Fatal(err)
		}
		return err
	}); err != nil {
		log.Fatal(err)
	}
}

func CreateBoltBucket(bucketName string) {
	if err := BoltCacheDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			panic(err.Error())
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func InitCacheBolt(path string, mode os.FileMode, options *bolt.Options) {
	db, err := bolt.Open(path, mode, options)
	if err != nil {
		log.Fatal(err)
	}
	BoltCacheDB = db
	defer db.Close()
}
