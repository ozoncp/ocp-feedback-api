package entity

// Entity is the interface that represents basic object
type Entity interface {
	Id() uint64
	UserId() uint64
}
