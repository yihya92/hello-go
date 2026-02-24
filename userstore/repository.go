package userstore

type UserRepository interface {
	Add(User) error
	Get(string) (User, error)
	List() []User
	Delete(string) error
}
