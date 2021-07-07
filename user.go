package main

// ログインユーザオブジェクト
type User struct {
	UserId    string
	Password  string
	FirstName string
	LastName  string
	RoleName  RoleName
}

// ユーザー権限
type RoleName int

const (
	ADMIN = iota
	USER
)
