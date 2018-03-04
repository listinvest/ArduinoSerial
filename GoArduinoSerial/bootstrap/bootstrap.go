package main

import (
	"log"
	"Arduino/ArduinoSerial/GoArduinoSerial"
	"Arduino/ArduinoSerial/Serial"
	"fmt"
)

func main (){
	ports, err := serial.GetPortsList()

	if err != nil{
		log.Fatal(err)
	}

	log.Println("Какой порт Вы хотите использовать?")

	for _, port := range ports{
		log.Println(port)
	}

	var port string

	fmt.Scanln(&port)

	log.Println("Порт:", port)

	mode := &serial.Mode{
		BaudRate: 9600,
	}

	connection := &GoArduinoSerial.Connection{}
	connection.NewConnection(port, mode)

	defer connection.Close()

	go func (){
		buffer := make([]byte, 100)
		for {
			text, isOk := connection.Read(&buffer)

			if isOk{
				fmt.Println(text)
			}
		}
	} ()

	var text string
	for {
		fmt.Scanln(&text)

		connection.Write(text)
	}
}