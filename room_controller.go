package main

type RoomController interface {
	ListRooms() []ReservableRoom
}

func NewRoomController(r RoomService) RoomController {
	return RoomControllerImpl{RoomService: r}
}

type RoomControllerImpl struct {
	RoomService RoomService
}

func (r RoomControllerImpl) ListRooms() []ReservableRoom {
	return r.RoomService.FindReservableRooms()
}
