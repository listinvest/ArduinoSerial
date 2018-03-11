package GoArduinoSerial

import (
	"log"
	"Arduino/ArduinoSerial/Serial"
)

type Connection struct{
	mode  *serial.Mode
	port  serial.Port
}

func NewConnection (portName string, mode *serial.Mode) (*Connection){
	connection := &Connection{}
	port, err := serial.Open(portName, mode)

	if err == nil{
		connection.mode = mode
		connection.port = port

		connection.Write("Подключено")
		log.Println ("Подключено")
	}else{
		log.Fatal(err)
	}

	return connection
}

func (connection *Connection) Write (text string) bool{
	_, err := connection.port.Write([]byte(text))

	return err == nil
}

func (connection *Connection) Read (buffer *[]byte) (string, bool){
	n, err := connection.port.Read(*buffer)

	if err != nil{
		log.Fatal(err)
	}

	if n == 0{
		log.Println("\nEOF")
	}

	return string(*buffer), err == nil
}

func (connection *Connection) Close (){
	connection.port.Close()
}