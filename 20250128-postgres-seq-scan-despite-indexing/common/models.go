package common

type User struct {
	Id       string `json:"id"`
	UserInfo string `json:"user_info"`
}

type NewUserEntity struct {
	Id       string
	Ssn      string
	UserInfo string
}

type ThirdPartyApiUserInfo struct {
	UserInfo string `json:"userInfo"`
}
