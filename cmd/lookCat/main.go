package main

import (
	other "LookCat/internal/linkSql"
	"fmt"
)

func main() {
	fmt.Println("hello")
	other.SpliceDsn()
	other.LinkSql()
}
