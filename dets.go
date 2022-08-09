package dets

import (
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"time"
)

var db *badger.DB

// SetDB 用以和dgraph-io/badge配合使用
func SetDB(d *badger.DB) {
	db = d
}

// Start 启动数据库
func Start(path string, l ...LogInterface) {
	if len(l) == 0 || l[0] == nil {
		logger = &defaultLogger{}
	} else {
		logger = l[0]
	}
	var err error
	db, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	go gc()
}

func gc() {
	ticker := time.NewTicker(time.Hour)
	for range ticker.C {
	again:
		err := db.RunValueLogGC(0.5)
		if err == nil {
			goto again
		}
	}
}

// Stop 关闭数据库并保存数据
func Stop() {
	_ = db.Close()
}

// Get returns the value associated with the key as a byte slice.
func Get(key []byte) []byte {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == badger.ErrKeyNotFound {
			return nil
		} else if err != nil {
			return errors.WithStack(err)
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		logger.Error(fmt.Sprintf("get failed, key: %s, error:\n%+v", string(key), err))
		return nil
	}
	return value
}

// Put 设置键值。你可以指定一个超时时间。
func Put(key []byte, value interface{}, ttl ...time.Duration) {
	var v []byte
	switch e := value.(type) {
	case []byte:
		v = e
	case []int:
		v, _ = json.Marshal(e)
	case []string:
		v, _ = json.Marshal(e)
	case map[string]map[string]string:
		v, _ = json.Marshal(e)
	case map[string]string:
		v, _ = json.Marshal(e)
	case map[string][]string:
		v, _ = json.Marshal(e)
	default:
		s, err := cast.ToStringE(value)
		if err != nil {
			panic(err)
		}
		v = []byte(s)
	}
	err := db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, v)
		if len(ttl) > 0 {
			e = e.WithTTL(ttl[0])
		}
		err := txn.SetEntry(e)
		return err
	})
	if err != nil {
		logger.Error(fmt.Sprintf("put failed, key: %s, value: %+v, error:\n%+v", string(key), string(v), err))
	}
}

// Del 删除键值
func Del(key []byte) {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	})
	if err != nil {
		logger.Error(fmt.Sprintf("delete failed, key: %s, error:\n%+v", string(key), err))
	}
}

// GetString returns the value associated with the key as a string.
func GetString(key []byte) string {
	if value := Get(key); value != nil {
		return string(value)
	}
	return ""
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key []byte) bool {
	return cast.ToBool(GetString(key))
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key []byte) int {
	return cast.ToInt(GetString(key))
}

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key []byte) int32 {
	return cast.ToInt32(GetString(key))
}

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key []byte) int64 {
	return cast.ToInt64(GetString(key))
}

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key []byte) uint {
	return cast.ToUint(GetString(key))
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key []byte) uint32 {
	return cast.ToUint32(GetString(key))
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key []byte) uint64 {
	return cast.ToUint64(GetString(key))
}

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key []byte) float64 {
	return cast.ToFloat64(GetString(key))
}

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key []byte) time.Duration {
	return cast.ToDuration(GetString(key))
}

// GetIntSlice returns the value associated with the key as a slice of int values.
func GetIntSlice(key []byte) []int {
	var s []int
	_ = json.Unmarshal(Get(key), &s)
	return s
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key []byte) []string {
	var s []string
	_ = json.Unmarshal(Get(key), &s)
	return s
}

// GetStringMap returns the value associated with the key as a map of strings.
func GetStringMap(key []byte) map[string]string {
	return cast.ToStringMapString(GetString(key))
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key []byte) map[string][]string {
	return cast.ToStringMapStringSlice(GetString(key))
}
