package dets

import (
	"bytes"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	Start("temp")
	defer Stop()
	testByteSlice(t)
	testString(t)
	testBool(t)
	testInt(t)
	testInt32(t)
	testInt64(t)
	testUint(t)
	testUint32(t)
	testUint64(t)
	testFloat64(t)
	testTime(t)
	testDuration(t)
	testIntSlice(t)
	testStringSlice(t)
	testStringMap(t)
	testStringMapStringSlice(t)
}

func testByteSlice(t *testing.T) {
	s := []byte("aaa bbb ccc")
	Put([]byte("byte slice"), s)
	b := Get([]byte("byte slice"))
	if !bytes.Equal(s, b) {
		t.Errorf("s: %s, b: %s", string(s), string(b))
	}
}

func testString(t *testing.T) {
	s := "aaa bbb ccc"
	Put([]byte("byte slice"), s)
	s2 := GetString([]byte("byte slice"))
	if s != s2 {
		t.Errorf("s: %s, s2: %s", s, s2)
	}
}

func testBool(t *testing.T) {
	Put([]byte("false"), false)
	Put([]byte("true"), true)
	b := GetBool([]byte("false"))
	b2 := GetBool([]byte("true"))
	if b || !b2 {
		t.Errorf("false: %v, true: %v", b, b2)
	}
}

func testInt(t *testing.T) {
	i := -1234567890
	Put([]byte("int"), i)
	i2 := GetInt([]byte("int"))
	if i != i2 {
		t.Errorf("i: %d, i2: %d", i, i2)
	}
}

func testInt32(t *testing.T) {
	i := int32(-1234567890)
	Put([]byte("int"), i)
	i2 := GetInt32([]byte("int"))
	if i != i2 {
		t.Errorf("i: %d, i2: %d", i, i2)
	}
}

func testInt64(t *testing.T) {
	i := int64(-1234567890123456)
	Put([]byte("int"), i)
	i2 := GetInt64([]byte("int"))
	if i != i2 {
		t.Errorf("i: %d, i2: %d", i, i2)
	}
}

func testUint(t *testing.T) {
	i := uint(1234567890)
	Put([]byte("int"), i)
	i2 := GetUint([]byte("int"))
	if i != i2 {
		t.Errorf("i: %d, i2: %d", i, i2)
	}
}

func testUint32(t *testing.T) {
	i := uint32(1234567890)
	Put([]byte("int"), i)
	i2 := GetUint32([]byte("int"))
	if i != i2 {
		t.Errorf("i: %d, i2: %d", i, i2)
	}
}

func testUint64(t *testing.T) {
	i := uint64(1234567890123456)
	Put([]byte("int"), i)
	i2 := GetUint64([]byte("int"))
	if i != i2 {
		t.Errorf("i: %d, i2: %d", i, i2)
	}
}

func testFloat64(t *testing.T) {
	f := 123.4
	Put([]byte("float"), f)
	f2 := GetFloat64([]byte("float"))
	if f != f2 {
		t.Errorf("f: %f, f2: %f", f, f2)
	}
}

func testTime(t *testing.T) {
	time1, _ := time.Parse("2006-01-02 15:04:05.999999", "2022-08-01 01:02:03.456789")
	Put([]byte("time"), time1)
	time2 := GetTime([]byte("time"))
	if !time1.Equal(time2) {
		t.Errorf("time1: %v, time2: %v", time1, time2)
	}
}

func testDuration(t *testing.T) {
	d := 7 * 24 * time.Hour
	Put([]byte("duration"), d)
	d2 := GetDuration([]byte("duration"))
	if d != d2 {
		t.Errorf("d: %v, d2: %v", d, d2)
	}
}

func testIntSlice(t *testing.T) {
	s := []int{111, 222}
	Put([]byte("slice"), s)
	s2 := GetIntSlice([]byte("slice"))
	if len(s) != len(s2) || s[0] != s2[0] || s[1] != s2[1] {
		t.Errorf("m: %v, m2: %v", s, s2)
	}
}

func testStringSlice(t *testing.T) {
	s := []string{"a", "b"}
	Put([]byte("slice"), s)
	s2 := GetStringSlice([]byte("slice"))
	if len(s) != len(s2) || s[0] != s2[0] || s[1] != s2[1] {
		t.Errorf("m: %v, m2: %v", s, s2)
	}
}

func testStringMap(t *testing.T) {
	m := map[string]string{"a": "aaa", "b": "bbb"}
	Put([]byte("map"), m)
	m2 := GetStringMap([]byte("map"))
	if len(m) != len(m2) || m["a"] != m2["a"] || m["b"] != m2["b"] {
		t.Errorf("m: %v, m2: %v", m, m2)
	}
}

func testStringMapStringSlice(t *testing.T) {
	m := map[string][]string{"a": {"aa"}}
	Put([]byte("map"), m)
	m2 := GetStringMapStringSlice([]byte("map"))
	if len(m) != len(m2) || len(m["a"]) != len(m2["a"]) || m["a"][0] != m2["a"][0] {
		t.Errorf("m: %v, m2: %v", m, m2)
	}
}
