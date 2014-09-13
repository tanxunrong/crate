package main

import (
	"fmt"
	"flag"
	sql "github.com/tanxunrong/crate/mysql"
)

func main() {
	pass := flag.String("p","","password")
	flag.Parse()
	param := new(sql.ConnectionParams)
	param.Host = "127.0.0.1"
	param.Port = 3306
	param.Uname = "root"
	param.Pass = *pass
	param.DbName = "test"
	param.Charset = "utf8"
	c,err := sql.Connect(*param)
	if err != nil {
		fmt.Errorf("conn err %s",err)
	}
	if ! c.IsClosed() {
		fmt.Println("conn alive")
		c.Close()
	}
}
