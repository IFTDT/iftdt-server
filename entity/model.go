package entity

type Model[T interface{}] interface {
	Create(*T) (err error)
	FindOne(id string) (t *T, err error)
	FindAll() (list []T, err error)
	Update(*T) (err error)
}
