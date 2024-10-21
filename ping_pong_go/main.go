package main

import (
	"time"
	"fmt"
	)

type Ball struct{ hits int }

func main() {
    table := make(chan *Ball) // create channel

    go player("ping", table)
    go player("pong", table)
    go player("year", table)

	fmt.Println("начало игры")
	time.Sleep(1 * time.Second)

    table <- new(Ball) // game on; toss the ball
    time.Sleep(1 * time.Second)
    <-table // game over;
	
}

func player(name string, table chan *Ball) {
    for {
        ball := <-table // блокируется пока в канал не передатут мяч
        ball.hits++
        fmt.Println(name, ball.hits)
        time.Sleep(100 * time.Millisecond)
        table <- ball
    }
}