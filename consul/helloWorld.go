package consul

import (
	"GoUtils/utils"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
	"sync"
	"time"
)

func NewConsul() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:18500"
	config.Scheme = "http"
	consul, err := api.NewClient(config)
	utils.CheckNoError(err)
	kv := consul.KV()
	d := &api.KVPair{Key: "ian/hello/world", Value: []byte("hello consul")}
	_, err = kv.Put(d, nil)
	utils.CheckNoError(err)
	value, _, err := kv.Get("ian/hello/world", nil)
	if value != nil {
		fmt.Println(string(value.Value))
		fmt.Println(string(value.Session))
		fmt.Println(string(value.Key))
	}

}



func NewConsul2() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go tryLock("mylock", "session"+strconv.Itoa(i), wg)
	}
	wg.Wait()

}

func tryLock(key string, sessionName string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Get a new client
	config := &api.Config{
		Address: "localhost:8500",
		Scheme:  "http",
	}

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	opts := &api.LockOptions{
		Key:         key,
		SessionName: sessionName,
	}

	//a,b,c := client.KV().Acquire(&api.KVPair{Key: "ian/hello/world", Value: []byte("hello consul")}, nil)
	lock, err := client.LockOpts(opts)
	log.Println(sessionName, "try to get lock obj")
	for i := 0; i < 3; i++ {
		leaderCh, err := lock.Lock(nil)
		if err != nil {
			log.Println("err", err, sessionName)
		}
		if leaderCh == nil {
			log.Println("err", err, sessionName)
			continue
		}
		log.Println(sessionName, "lock and sleep")
		time.Sleep(5 * time.Second)
		err = lock.Unlock()
		if err != nil {
			log.Fatal("err", err)
		}
		log.Println(sessionName, "unlock")
		time.Sleep(5 * time.Second)
	}
}
