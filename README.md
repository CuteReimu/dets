# dets

![](https://img.shields.io/github/languages/top/CuteReimu/dets "语言")
[![](https://img.shields.io/github/workflow/status/CuteReimu/dets/Go)](https://github.com/CuteReimu/dets/actions/workflows/golangci-lint.yml "代码分析")
[![](https://img.shields.io/github/license/CuteReimu/dets)](https://github.com/CuteReimu/dets/blob/master/LICENSE "许可协议")

基于 [dgraph-io/badger](https://github.com/dgraph-io/badger) 做的内嵌式key-value型数据库的**接口封装**。

因为仅仅是**接口封装**，所以**对于同一个key，不支持并发**。想要支持并发，请直接使用[dgraph-io/badger](https://github.com/dgraph-io/badger)

## 安装

```bash
go get github.com/CuteReimu/dets
```

## 使用方法举例

```go
package main

import (
    "fmt"
    "github.com/CuteReimu/dets"
)

func main() {
    dets.Start("temp")
	key := []byte("aaa")
	dets.Put(key, "vvv")
    s := dets.GetString(key)
	dets.Del(key)
    fmt.Println(s)
    dets.Stop()
}
```

## 函数一览

`Put`和`Del`函数统一使用，`Get`用了不同的函数名

| 支持的value类型               | 对应`Get`函数名                | 
|--------------------------|---------------------------|
| `[]byte`                 | `Get`                     |
| `string`                 | `GetString`               |
| `bool`                   | `GetBool`                 |
| `int`                    | `GetInt`                  |
| `int32`                  | `GetInt32`                |
| `int64`                  | `GetInt64`                |
| `uint`                   | `GetUint`                 |
| `uint32`                 | `GetUint32`               |
| `uint64`                 | `GetUint64`               |
| `float64`                | `GetFloat64`              |
| `time.Time`              | `GetTime`                 |
| `time.Duration`          | `GetDuration`             |
| `[]int`                  | `GetIntSlice`             |
| `[]string`               | `GetStringSlice`          |
| `map[string]interface{}` | `GetStringMap`            |
| `map[string]string`      | `GetStringMapString`      |
| `map[string][]string`    | `GetStringMapStringSlice` |
