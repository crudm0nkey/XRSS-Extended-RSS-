package main

import (
	db "del/internal"
	luafunctions "del/internal/luaFunctions"
	"fmt"
)

func main() {
	db := db.SetupDb()
	fmt.Println("DB IUP")
	db.Close()

	luafunctions.LoadPlugins()
}
