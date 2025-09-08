package main

import (
	"fmt"
	"time"
)

// Реализация через канал и таймер
func sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C // Блокируемся до получения сигнала из канала
}

// Альтернативная реализация через цикл проверки времени
func sleepLoop(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		// Просто ждем в цикле
	}
}

func main() {
	fmt.Println("Начало:", time.Now())

	sleep(2 * time.Second)
	fmt.Println("Прошло 2 секунды:", time.Now())

	sleepLoop(1 * time.Second)
	fmt.Println("Прошло еще 1 секунда:", time.Now())
}
