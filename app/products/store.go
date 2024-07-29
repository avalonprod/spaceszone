package products

type store interface {
	Save()
	Get()
	GetAll()
	Update()
	Delete()
}
