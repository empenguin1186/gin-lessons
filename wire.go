package main

import "github.com/google/wire"

// +build wireinject
// このコードはgo buildには含まない

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
