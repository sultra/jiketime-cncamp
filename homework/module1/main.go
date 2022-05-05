package main

import (
	"fmt"
	"time"
)

func main() {
	// subject1_1()
	subject1_2()
}

/*
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
*/
func subject1_1() {
	msgs := []string{"I", "am", "stupid", "and", "weak"}

	for i, _ := range msgs {
		switch i {
		case 2:
			msgs[i] = "smart"
		case 4:
			msgs[i] = "strong"
		default:
		}
	}

	fmt.Printf("msgs %+v\n", msgs)
}

/*

队列：
队列长度 10，队列元素类型为 int
生产者：
每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
消费者：
每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞

*/
func subject1_2() {
	ch := make(chan int, 10)
	defer close(ch)
	go prod(ch)
	go consumer(ch)
	time.Sleep(13 * time.Second)
}

// 生产者,只发送通道
func prod(ch chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("fa fa")
		ch <- i
	}
}

func consumer(ch <-chan int) {
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		v := <-ch
		fmt.Println(v)
	}
}
