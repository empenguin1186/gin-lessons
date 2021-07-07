package main

type RoomService interface {
	FindReservableRooms() []ReservableRoom
}

type RoomServiceImpl struct {
	ReservableRoomRepository ReservableRoomRepository
}

func NewRoomService(r ReservableRoomRepository) RoomService {
	return RoomServiceImpl{ReservableRoomRepository: r}
}

func (r RoomServiceImpl) FindReservableRooms() []ReservableRoom {
	return r.ReservableRoomRepository.FindByReservableRoomIdReservedDateOrderByReservableRoomIdRoomIdAsc()
}
