package room

// ログインユーザオブジェクト
type User struct {
	userId    string
	password  string
	firstName string
	lastName  string
	roleName  RoleName
}

// ユーザー権限
type RoleName int

const (
	ADMIN = iota
	USER
)
