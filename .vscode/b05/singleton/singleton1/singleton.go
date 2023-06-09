package singleton

import (
	"fmt"
	"sync"
	"time"
)

type singleton struct{
	number int 
}

var instance *singleton

/* Một số cách cài đặt Singleton Pattern*/

// Sử dụng hàm khởi tạo
func GetInstanceFunc() *singleton {
	if instance == nil { // Zero value
		time.Sleep(time.Second)
		instance = &singleton{ number : 1 }
	}
	return instance
}

// Sử dụng sync
var once sync.Once

func GetInstanceSync() *singleton {
	//
	once.Do(func() {
		instance = &singleton{ number : 1 }
	})
	return instance
}

// Sử dụng mutex lock
var mu sync.Mutex

func GetInstanceMutex() *singleton {
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		instance = &singleton{ number : 1 }
	}
	return instance
}

// Sử dụng Init của package
func init() {
	instance = &singleton{ number: 1}
	fmt.Println("Singleton init")
}

func GetInstanceInit() *singleton {
	return instance
}

// DEMO

func DemoGetInstanceFunc() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", GetInstanceFunc())
		}()
	}
}

func DemoGetInstanceSync() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", GetInstanceSync())
		}()
	}
}

func DemoGetInstanceMutex() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", GetInstanceMutex())
		}()
	}
}

func DemoGetInstanceInit() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", GetInstanceInit())
		}()
	}
}