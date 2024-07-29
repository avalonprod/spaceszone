package users

type store interface {
	Save()
	Get()
	Update()
	SoftDelete()
}
