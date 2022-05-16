package server

type ErrorCode int

const (
	SomethingWentWrong = iota
	InvalidData
	AccessDenied
)
