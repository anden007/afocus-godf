package interface_core

type IViewModel interface {
	FromView() error
	FromDB() error
}