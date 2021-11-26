package login

type Repository interface {
	LoginRepo(email string, password string) error
}
type Service interface {
}
