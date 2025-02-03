package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user Users) (Users, error) {
	return s.repo.CreateUser(user)
}
func (s *UserService) GetAllUsers() ([]Users, error) {
	return s.repo.GetAllUsers()
}
func (s *UserService) UpdateUserByID(id uint, user Users) (Users, error) {
	return s.repo.UpdateUserByID(id, user)
}
func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
