// +build wireinject
package main

import "github.com/google/wire"

// このコードはgo buildには含まない

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
