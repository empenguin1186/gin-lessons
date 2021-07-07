package main

type MeetingRoom struct {
	RoomId   int
	RoomName string

	// setter, getter は必要になったら実装
}

type ReservableRoom struct {
	ReservableRoomId ReservableRoomId
	MeetingRoom      MeetingRoom
}

type ReservableRoomId struct {
	RoomId       int
	ReservedDate float32
}

type Reservation struct {
	ReservationId  int
	StartTime      float32
	EndTime        float32
	ReservableRoom ReservableRoom
	User           User
}
