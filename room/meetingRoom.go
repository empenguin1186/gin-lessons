package main

type MeetingRoom struct {
	roomId   int
	roonName string

	// setter, getter は必要になったら実装
}

type ReservableRoom struct {
	reservableRoomId ReservableRoomId
	meetingRoom      MeetingRoom
}

type ReservableRoomId struct {
	roomId       int
	reservedDate float32
}

type Reservation struct {
	reservationId  int
	startTime      float32
	endTime        float32
	reservableRoom ReservableRoom
	user           User
}
