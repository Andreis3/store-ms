package idatabase

type IDatabase interface {
	InstanceDB() any
	Close()
}
