package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	cancel()
	// }()

	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Time out, can't book hotel")
	case <-time.After(time.Second * 5):
		fmt.Println("hotel booked")
	}
}
