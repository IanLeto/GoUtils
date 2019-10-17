package main

import (
	"fmt"
	"goUtils/cmd"
	"io/ioutil"
	"os"
	"time"
	"transfer/types"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}

	t := time.Now()
	//a := t
	//b := t
	fmt.Println(types.NewTimeStamp(t))
	fmt.Println(t.Zone())
	fmt.Println(t.UTC())
	fmt.Println(types.NewTimeStamp(t.In(time.FixedZone("UTC", 2*60*60))))
	fmt.Println(t.Zone())
	fmt.Println(t.UTC())

	loc := time.FixedZone("CST", 3600*4)
	t.In(loc)
	fmt.Println(time.ParseInLocation("2006-01-02 03:04:05", "2019-10-15 01:01:24", loc))
	filePath := "D:\\GoPath\\src\\transfer\\xxx.txt"
	data := []byte("xx")
	err := ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		panic(err)
	}

}

func CheckFileExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
