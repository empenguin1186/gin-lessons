package main

type ReservableRoomRepository interface {
	FindByReservableRoomIdReservedDateOrderByReservableRoomIdRoomIdAsc() []ReservableRoom
}

type ReservableRoomRepositoryImpl struct{}

func (r ReservableRoomRepositoryImpl) FindByReservableRoomIdReservedDateOrderByReservableRoomIdRoomIdAsc() []ReservableRoom {
	var result []ReservableRoom
	var foo = ReservableRoomId{RoomId: 1, ReservedDate: 10000}
	var bar = MeetingRoom{RoomId: 1, RoomName: "hoge"}

	var baz = ReservableRoom{ReservableRoomId: foo, MeetingRoom: bar}
	result = append(result, baz)
	return result
}
