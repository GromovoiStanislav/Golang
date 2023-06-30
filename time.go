package main

import (
	"log"
	"time"
)

func main() {
	log.Println("start") // 2009/11/10 23:00:00 start
	time.Sleep(5 * time.Second)
	log.Println("end") // 2009/11/10 23:00:05 end
}