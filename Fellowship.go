package main

import (
    "fmt"
    "time"
)

type Token struct {
	data string
	recipient int //получатель
}

var n = 5

var chann = make([]chan Token, n) //создаем массив из n каналов которые принмают тип Token
func main() {
	var i int 
	for i = range chann{
		chann[i] = make(chan Token) //создаем каналы
	}
	var token Token;
	token.recipient = n
	token.data = "token"
	go start(chann[0], token)//поток, который записывает в 0 канал token
	for i = 0; i<n-1; i++ {
		go send(chann[i], chann[i+1],i) //n-1 потоки для передачи token по каналам: i-ый передает i+1-му
	}
	go final(chann[n-1], n-1)//поток-чтение содержимого (token) из последнего канала
	time.Sleep(1 * 1e9) // на случай если главная функция закончит раньше др.горутинс
}

func start (channel chan Token, t Token){
	fmt.Println("Start: send token to 1 channel")
	channel <- t  //запись в канал
}

func send (a <-chan  Token, b chan <- Token, num int){
	fmt.Println("read from channel ", num+1)
	fmt.Println("write into channel ", num+2)
	for {
		t := <-a  //чтение из канала
		b <- t   //запись в следующий канал
	}
}

func final(last <-chan Token, num int){
	for {
		fmt.Println("Final: recipient get token ", <-last, " from the last channel ", num+1) //чтение из канала
	}
}

//https://play.golang.org/p/M0aPaqbtQp