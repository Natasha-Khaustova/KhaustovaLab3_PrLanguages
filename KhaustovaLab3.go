package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Token struct {
	data      string
	recipient int
	ttl       int
}

func main() {
	var N int
	fmt.Print("Enter the number of channels -> ")
	fmt.Fscan(os.Stdin, &N)
	ch := make([]chan Token, N)
	var idRecipient int
	fmt.Print("Enter the recipient channel -> ")
	fmt.Fscan(os.Stdin, &idRecipient)

	for i := 0; i < N; i++ {
		ch[i] = make(chan Token)
	}

	for i := 0; i < N; i++ {
		go func(i int, chIn chan Token, chOut chan Token) {
			for {
				token := <-chIn
				if token.ttl <= 0 {
					fmt.Println("Lifetime is over")
					return
				}
				if i == token.recipient {
					fmt.Println("Received:", token)
				} else {
					token.ttl--
					chOut <- token
				}
			}
		}(i, ch[i], ch[(i+1)%N])
	}

	randomNumber := rand.Intn(N + 1)
	fmt.Println("Sender number ", randomNumber)

	ch[randomNumber] <- Token{data: "Here can be your message", recipient: idRecipient, ttl: 10000} //значение ttl указано для эмулирования
	// стандартной задержки 10мс TokenRing
	time.Sleep(time.Second * 10)
}
