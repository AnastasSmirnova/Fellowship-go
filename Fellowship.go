package main

import (
    "fmt"
    "time"
)

type Token struct {
	data string
	recipient int //����������
}

var n = 5

var chann = make([]chan Token, n) //������� ������ �� n ������� ������� �������� ��� Token
func main() {
	var i int 
	for i = range chann{
		chann[i] = make(chan Token) //������� ������
	}
	var token Token;
	token.recipient = n
	token.data = "token"
	go start(chann[0], token)//�����, ������� ���������� � 0 ����� token
	for i = 0; i<n-1; i++ {
		go send(chann[i], chann[i+1],i) //n-1 ������ ��� �������� token �� �������: i-�� �������� i+1-��
	}
	go final(chann[n-1], n-1)//�����-������ ����������� (token) �� ���������� ������
	time.Sleep(1 * 1e9) // �� ������ ���� ������� ������� �������� ������ ��.��������
}

func start (channel chan Token, t Token){
	fmt.Println("Start: send token to 1 channel")
	channel <- t  //������ � �����
}

func send (a <-chan  Token, b chan <- Token, num int){
	fmt.Println("read from channel ", num+1)
	fmt.Println("write into channel ", num+2)
	for {
		t := <-a  //������ �� ������
		b <- t   //������ � ��������� �����
	}
}

func final(last <-chan Token, num int){
	for {
		fmt.Println("Final: recipient get token ", <-last, " from the last channel ", num+1) //������ �� ������
	}
}

//https://play.golang.org/p/M0aPaqbtQp