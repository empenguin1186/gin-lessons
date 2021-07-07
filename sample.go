package main

import "fmt"

func main2() {
	var RoomService = NewRoomService(ReservableRoomRepositoryImpl{})
	var RoomController = NewRoomController(RoomService)
	fmt.Printf(RoomController.ListRooms()[0].MeetingRoom.RoomName)
}

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
