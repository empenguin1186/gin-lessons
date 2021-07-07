package room

type ReservableRoomRepository interface {
	findByReservableRoomIdReservedDateOrderByReservableRoomIdRoomIdAsc()
}

type ReservableRoomRepositoryImpl struct{}

func (r ReservableRoomRepositoryImpl) findByReservableRoomIdReservedDateOrderByReservableRoomIdRoomIdAsc() {
	var list = ReservableRoom{}
}
