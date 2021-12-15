package serviceauth

import (
	"errors"
	"go-mysql-api/dto"
	"go-mysql-api/entity"
	"go-mysql-api/usecase/auth"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
	repo auth.Repository
}

func NewAuthService(auth auth.Repository) *Service {
	return &Service{auth}
}

func (auth *Service) DoLogin(dto dto.GetAuthUserRequest) (dto.GetUserResponse, error) {
	user, err := auth.repo.AuthLogin(dto.Email, dto.Password)
	usr := auth.mapUserEntityToGetUserByAuthDTO(user)
	return usr, err
}

func (auth *Service) CreateToken(userId uint64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET_KEY")))
	if err != nil {
		return "", err
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
