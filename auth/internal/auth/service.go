package auth

import (
	"errors"

	"example.com/auth/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userRepo   *user.Repository
	jwtService *JWTService
}

func NewService(userRepo *user.Repository, jwtService *JWTService) *Service {
	return &Service{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *Service) Register(input RegisterRequest) (*user.User, error) {
	existingUser, err := s.userRepo.FindByEmail(input.Email)

	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(newUser)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) Login(input LoginRequest) (string, *user.User, error) {
	existingUser, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.Password))
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}
	token, err := s.jwtService.Generate(existingUser.ID)
	if err != nil {
		return "", nil, err
	}
	return token, existingUser, nil
}

func (s *Service) GetUserByID(id uint) (*user.User, error) {
	return s.userRepo.FindByID(id)
}
