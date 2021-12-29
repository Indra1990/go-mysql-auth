package serviceauth

import (
	"errors"
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/auth"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo auth.Repository
}

func NewAuthService(auth auth.Repository) *Service {
	return &Service{auth}
}
func (auth *Service) RegisterUserInput(dto dto.UserCreateRequest) (dto.GetUserResponse, error) {
	user := entity.User{}
	user.Name = dto.Name
	user.Email = dto.Email
	password, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.MinCost)
	user.Password = string(password)
	usrRegis, err := auth.repo.UserRegister(user)
	usr := auth.mapUserEntityToGetUserByAuthDTO(usrRegis)
	if err != nil {
		return usr, err
	}

	return usr, nil
}

func (auth *Service) DoLogin(dto dto.GetAuthUserRequest) (dto.GetUserResponse, error) {
	user, err := auth.repo.FindByID(dto.Email)
	usr := auth.mapUserEntityToGetUserByAuthDTO(user)

	if err != nil {
		return usr, err
	}

	if user.ID == 0 {
		return usr, errors.New("user not Found")
	}

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if comparePasswordErr != nil {
		return usr, comparePasswordErr
	}
	return usr, nil
}

func (auth *Service) CreateToken(userId uint64) (string, error) {
	atClaims := jwt.MapClaims{}
	// atClaims["authorized"] = true
	atClaims["user_id"] = userId
	// atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))
	if err != nil {
		return token, err
	}
	return token, nil
}

func (auth *Service) mapUserEntityToGetUserByAuthDTO(ent entity.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:    ent.ID,
		Name:  ent.Name,
		Email: ent.Email,
	}
}

func (auth *Service) ValidateToken(tokenEncoded string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenEncoded, func(tokenEncoded *jwt.Token) (interface{}, error) {
		_, okay := tokenEncoded.Method.(*jwt.SigningMethodHMAC)
		if !okay {
			return nil, errors.New("invalid token")
		}

		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
