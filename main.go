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
