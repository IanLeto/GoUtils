package main

import (
	"goUtils/cmd"
	"os"
)

func main() {
	cmd.Execute()
	//if err := cmd.RootCmd.Execute(); err != nil {
	//	panic(err)
	//}
	//
	//t := time.Now()
	////a := t
	////b := t
	//fmt.Println(types.NewTimeStamp(t))
	//fmt.Println(t.Zone())
	//fmt.Println(t.UTC())
	//fmt.Println(types.NewTimeStamp(t.In(time.FixedZone("UTC", 2*60*60))))
	//fmt.Println(t.Zone())
	//fmt.Println(t.UTC())
	//
	//loc := time.FixedZone("UTC", 3600*2)
	//t.In(loc)
	//fmt.Println(time.ParseInLocation("2006-01-02 03:04:05", "2019-10-15 01:01:24", loc))
	//loc = time.FixedZone("UTC+2", 3600*2)
	//t.In(loc)
	//fmt.Println(time.ParseInLocation("2006-01-02 03:04:05", "2019-10-15 01:01:24", loc))
	//fmt.Println(time.Parse("2006-01-02 03:04:05", "2019-10-15 01:01:24"))
	//fmt.Println(time.Parse("2006-01-02 03:04:05", "2019-10-15 01:01:24"))
	//filePath := "D:\\GoPath\\src\\transfer\\xxx.txt"
	//data := []byte("xx")
	//err := ioutil.WriteFile(filePath, data, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//
	//viper.SetConfigName("config")
	//viper.AddConfigPath(".")
	//if err := viper.ReadInConfig(); err != nil {
	//	panic(err)
	//}
	//var redis ConfigRedis
	//err = viper.Unmarshal(&redis)
	//if err != nil {
	//	panic(err)
	//}

}

type ConfigRedis struct {
	IP string
	DB int
}

func CheckFileExist(filePath string) bool {
	var exist = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
