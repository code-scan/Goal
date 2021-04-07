package Gsensor

type Result map[string]string
type ISensor interface {
	GetInfo() string
	SetDomain(string)
	SetAccount(string)
	SetPassword(string)
	SetType(string)
	GetResult() Result
	Login(bool) bool
}
