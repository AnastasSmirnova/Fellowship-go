package main // начало

import (   // подключение библиотеки
    "fmt"  //ввод/вывод
)

type Token struct {//=класс
	data string  //сообщение
	recipient int //номер получателя
}

var count_gorout = 10 //кол.во горутин  
var out_chan chan Token // последний(блокирующий для последующих) канал для получателя
func main() {
	var token Token;
	token.recipient = 4  //номер получателя
	token.data = "token"  //сообщение
	if token.recipient >= count_gorout { // кол.во горутин меньше номера получателя...
		count_gorout = token.recipient//...переопределяем теперь кол.во горутин=номеру получателя
	}
	out_chan = make(chan Token)  // последний канал в который потом записали
	var chanmass = make([]chan Token, count_gorout+1)// создаем массив каналов в котором хранится токен,
	for i:= range chanmass{    //цикл по массиву
		chanmass[i] = make(chan Token)//создание одного канала
	}	
	for i := count_gorout-1; i >=0; i-- {    //создает последовательно ,а получает от последнего..
		go send(chanmass[i],chanmass[i+1],i)//запуск n-штук  потоков
	}
	chanmass[0] <- token //запись токена в 0ой канал
	fmt.Println(<-out_chan)//считывание из out_chan
}

func send(current_chan chan Token,next_chan chan Token, n int) {// получает текущий канал потом следующ...потом номер текущего
	fmt.Println("Создался поток номер: ", n+1)
	token := <- current_chan //чтение из текущего канала туда
	if n+1 == token.recipient {//если номер потока==номеру получателя
		fmt.Println("поток ",n+1," Token достиг своего получателя ",token.recipient)
		out_chan <- token // передача в последний канал
	} else {
		fmt.Println("поток ",n+1,"передает следущему,т.к. получатель=", token.recipient)
		next_chan <- token
	}
}


//https://play.golang.org/p/iVeNJGhhSP