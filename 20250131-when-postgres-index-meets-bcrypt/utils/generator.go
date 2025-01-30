package utils

import (
	"20250128-postgres-seq-scan-despite-indexing/common"
	"fmt"
	"github.com/google/uuid"
	"math/rand/v2"
	"strconv"
	"time"
)

func GenRandomUsers(num int) []common.NewUserEntity {
	users := make([]common.NewUserEntity, num)
	for i := 0; i < num; i++ {
		users[i] = GenRandomUser()
	}
	return users
}

func GenRandomUser() common.NewUserEntity {
	ssn := genRandomSsn()
	fmt.Println("Generated SSN:", ssn)
	return common.NewUserEntity{
		Id:       uuid.New().String(),
		Ssn:      ssn,
		UserInfo: GenRandomUserInfo(),
	}
}

func GenRandomUserInfo() string {
	return fmt.Sprintf("Some info %s %s", strconv.FormatInt(time.Now().UnixMilli(), 10), uuid.New().String())
}

func genRandomSsn() string {
	ssn := ""
	for i := 0; i < 10; i++ {
		ssn += strconv.Itoa(rand.IntN(10))
	}
	return ssn
}
