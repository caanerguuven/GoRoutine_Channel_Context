package main

import (
	"context"
	"fmt"
	"time"
)

var intList = []int{1}
var duration = 15 * time.Second

func runAccount(id int, index int, acc string, accountChan chan string) {
	time.Sleep(time.Second * 5)
	fmt.Printf(" index : %d --> Bank No : %d  Account No : %s Completed \n", index, id, acc)
	accountChan <- acc
}

func runSQL(ctx context.Context, id int, accChan chan string) {

	go func() {
		i := 0
		for acc := range accChan {
			i = i + 1
			go runAccount(id, i, acc, accChan)

			select {
			case <-ctx.Done():
				return
			default:
				continue
			}
		}
	}()
}

func runAutomation(index int, id int, idChan chan int) {
	fmt.Printf("Bank %d starting. Channel Len : %d \n", index, id)

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	accChan := make(chan string)
	accList := []string{"A", "B", "C"} //This is where I got datas from sql.

	go func() {
		for _, acc := range accList {
			accChan <- acc
		}
	}()

	go runSQL(ctx, id, accChan)

	time.Sleep(duration)

	idChan <- id
	fmt.Printf("Worker %d done\n", id)

}

func main() {
	automationChannel := make(chan int)
	go func() {
		for _, automation := range intList {
			fmt.Printf("Channel %d was added. \n", automation)
			automationChannel <- automation
		}
	}()

	i := 0
	for chnnl := range automationChannel {
		i = i + 1
		go runAutomation(i, chnnl, automationChannel)
	}
}
