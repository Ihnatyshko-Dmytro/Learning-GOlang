package main

import (
	"context"
	"fmt"
	"golang-ninja/basic/shape"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/scheduler"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	square := shape.NewSquare(5)
	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, parseSubscriptionData, time.Second*5)
	worker.Add(ctx, sendStatistics, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()

	PrintShapeArea(square)
	// PrintShapeArea(circle)

	// printInterface(square)
	// printInterface(circle)
	// printInterface("sada")
	// printInterface(77)
	// printInterface("ggtdvdv")
	// printInterface(true)
}

func parseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf("subscription parsed successfuly at %s\n", time.Now().String())
}

func sendStatistics(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("statistics sent at %s\n", time.Now().String())
}

func PrintShapeArea(s shape.Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimetr())
}

func printInterface(i interface{}) {
	// switch value := i.(type) {
	// case int:
	//
	//	fmt.Println("int", value)
	//
	// case bool:
	//
	//	fmt.Println("bool", value)
	//
	// default:
	//
	//		fmt.Println("unknown type", value)
	//	}
	//
	// fmt.Printf("%#v\n", i)
}
