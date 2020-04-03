package currencyInGo

import (
	"GoUtils/utils"
	"context"
	"fmt"
	"sync"
	"time"
)

/*
 */

func CtxDemo() {
	var (
		wg   sync.WaitGroup
		done = make(chan interface{})
	)
	defer close(done)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			utils.NoErr(err)
			return
		}
	}()
}

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world\n", greeting)
	return nil
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
	case err != nil:
		return "", err
	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")

}

func locale(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("canceled")
	case <-time.After(1 * time.Minute):

	}
	return "EN/US", nil
}

func GrandFather() {
	ticker := time.NewTicker(5 * time.Second)
	ticker2 := time.NewTicker(10 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	go Father(ctx)
	for {
		select {
		case <-ticker.C:
			fmt.Println("已到")
		case <-ticker2.C:
			cancel()
		}
	}
}

func Father(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	ticker2 := time.NewTicker(15 * time.Second)
	// 关键在这里，我们要创建一个新ctx
	ctx2, cancel := context.WithCancel(context.Background())
	go Son(ctx2)
	for {
		select {
		case <-ticker.C:
			fmt.Println("晨时已到")
		case <-ticker2.C:
			cancel()
		case <-ctx.Done():
			fmt.Println("father done")
			cancel()
			return
		}
	}

}
func Son(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("son done")
			return
		case <-ticker.C:
			fmt.Println("午时已到")

		}
	}

}
