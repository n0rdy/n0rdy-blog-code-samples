package service

import (
	"20250131-when-postgres-index-meets-bcrypt/common"
	"20250131-when-postgres-index-meets-bcrypt/db"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type Service struct {
	repo       *db.Repo
	httpClient *http.Client
}

func NewService(repo *db.Repo) *Service {
	return &Service{
		repo:       repo,
		httpClient: &http.Client{},
	}
}

func (s *Service) GetUsers() ([]common.User, error) {
	return s.repo.SelectUsers()
}

func (s *Service) GetUser(ssn string) (*common.User, error) {
	user, err := s.repo.SelectUserBySsn(ssn)
	if err != nil {
		return nil, err
	}
	if user == nil || user.Id == "" {
		fmt.Println("User not found, making a request to the third-party API...")

		resp, err := s.httpClient.Get("http://localhost:8081/third-party/api/v1/user-info/" + ssn)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var userInfo common.ThirdPartyApiUserInfo
		err = json.NewDecoder(resp.Body).Decode(&userInfo)
		if err != nil {
			return nil, err
		}

		newUser := common.NewUserEntity{
			Id:       uuid.New().String(),
			Ssn:      ssn,
			UserInfo: userInfo.UserInfo,
		}
		err = s.repo.InsertUser(newUser)
		if err != nil {
			return nil, err
		}
		user = &common.User{
			Id:       newUser.Id,
			UserInfo: newUser.UserInfo,
		}
	}
	return user, nil
}
