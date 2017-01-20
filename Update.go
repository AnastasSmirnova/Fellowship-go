package main // ������

import (   // ����������� ����������
    "fmt"  //����/�����
)

type Token struct {//=�����
	data string  //���������
	recipient int //����� ����������
}

var count_gorout = 10 //���.�� �������  
var out_chan chan Token // ���������(����������� ��� �����������) ����� ��� ����������
func main() {
	var token Token;
	token.recipient = 4  //����� ����������
	token.data = "token"  //���������
	if token.recipient >= count_gorout { // ���.�� ������� ������ ������ ����������...
		count_gorout = token.recipient//...�������������� ������ ���.�� �������=������ ����������
	}
	out_chan = make(chan Token)  // ��������� ����� � ������� ����� ��������
	var chanmass = make([]chan Token, count_gorout+1)// ������� ������ ������� � ������� �������� �����,
	for i:= range chanmass{    //���� �� �������
		chanmass[i] = make(chan Token)//�������� ������ ������
	}	
	for i := count_gorout-1; i >=0; i-- {    //������� ��������������� ,� �������� �� ����������..
		go send(chanmass[i],chanmass[i+1],i)//������ n-����  �������
	}
	chanmass[0] <- token //������ ������ � 0�� �����
	fmt.Println(<-out_chan)//���������� �� out_chan
}

func send(current_chan chan Token,next_chan chan Token, n int) {// �������� ������� ����� ����� �������...����� ����� ��������
	fmt.Println("�������� ����� �����: ", n+1)
	token := <- current_chan //������ �� �������� ������ ����
	if n+1 == token.recipient {//���� ����� ������==������ ����������
		fmt.Println("����� ",n+1," Token ������ ������ ���������� ",token.recipient)
		out_chan <- token // �������� � ��������� �����
	} else {
		fmt.Println("����� ",n+1,"�������� ���������,�.�. ����������=", token.recipient)
		next_chan <- token
	}
}


//https://play.golang.org/p/iVeNJGhhSP