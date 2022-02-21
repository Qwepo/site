package utilty

import (
	"time"

	"gitlab.com/knopkalab/go/utils"
	"golang.org/x/crypto/bcrypt"
)

//primory key
type PK = int64

// Unix time
type Unix int64

func (t Unix) Int64() int64 { return int64(t) }

// SetNow time
func (t *Unix) SetNow() {
	*t = Unix(time.Now().Unix())
}

func NewPassword(pass string) string {
	newPass := setPasswordHash(pass)
	return newPass
}

func Match(passHash, pass string) bool {
	return bcrypt.CompareHashAndPassword(utils.ZAtob(passHash), utils.ZAtob(pass)) == nil

}

func setPasswordHash(pass string) string {
	if len(pass) == 0 {
		return ""	
	}
	result, err := bcrypt.GenerateFromPassword(utils.ZAtob(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(result)
}
