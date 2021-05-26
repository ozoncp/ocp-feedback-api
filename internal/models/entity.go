package models

// Entity is an interface that represents an object with an id
type Entity interface {
	ObjectId() uint64
}
