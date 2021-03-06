# Simple wrapper functions based go-i18n package

How to usage:

```
package main

import (
	"encoding/json"
	"github.com/dualface/go-i18n-helper/i18h"
	"log"
)

func main() {
	langs := map[string]string{
		"zh-cn": "tests/zh-cn.json",
		"en":    "tests/en.json",
	}
	err := i18h.Load(langs, "json", json.Unmarshal)
	if err != nil {
		log.Panic(err)
	}

	T := i18h.Lang("zh-cn")
	println(T("hello", "World"))
	println(T("not_found_dbtable", 3099, "admin_db", "users"))
	println()
	println()

	T = i18h.Lang("en")
	println(T("hello", "World"))
	println(T("not_found_dbtable", 3099, "admin_db", "users"))
	println()
	println()

	T = i18h.Lang("invalid-language")
	println(T("hello", "World"))
	println(T("not_found_dbtable", 3099, "admin_db", "users"))
	println()
	println()
}
```

Output:

```
你好，World
错误代码 [3099] 在数据库 'admin_db' 中没有找到表格 'users'


Hello, World
ERR_CODE 3099: NOT FOUND TABLE 'users' IN DB 'admin_db'


#I18N MISS# hello []interface {}{"World"}
#I18N MISS# not_found_dbtable []interface {}{3099, "admin_db", "users"}
```

\- EOF \-
