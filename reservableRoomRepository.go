package main

type ReservableRoomRepository interface {
	findByReservableRoomIdReservedDateOrderByReservableRoomIdRoomIdAsc()
}

type ReservableRoomRepositoryImpl struct{}

func (r ReservableRoomRepositoryImpl) findByReservableRoomIdReservedDateOrderByReservableRoomIdRoomIdAsc() []ReservableRoom {
	var result []ReservableRoom
	var foo = ReservableRoomId{RoomId: 1, ReservedDate: 10000}
	var bar = MeetingRoom{RoomId: 1, RoonName: "hoge"}

	var baz = ReservableRoom{ReservableRoomId: foo, MeetingRoom: bar}
	result[0] = baz
	return result
}
