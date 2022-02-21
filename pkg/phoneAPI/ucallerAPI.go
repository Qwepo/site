package phoneapi

import (
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/knopkalab/go/logger"

	"math/rand"
	"net/http"
	"strconv"
)

type Ucaller struct {
	ServiceID string
	SecretKey string
	Debug     bool
}

type UcallerResponse struct {
	Status     bool   `json:"status"`
	UcallerId  int64  `json:"uncaller_id"`
	Phone      int    `json:"phone"`
	Code       int    `json:"code"`
	Client     string `json:"client"`
	UniqueRtId string `json:"nique_request_id"`
	Exists     bool   `json:"exists"`
}

var errUcaller = errors.New("error send code uCaller")

func (u *Ucaller) Send(code string, phone string, log logger.Logger) {
	fmt.Println(u.Debug)
	if !u.Debug {
		request := fmt.Sprintf("https://api.ucaller.ru/v1.0/initCall?service_id=%s&key=%s&phone=%s&code=%s", u.ServiceID, u.SecretKey, phone, code)
		resp, err := http.Get(request)
		if err != nil {
			log.Fatal().Err(err)

		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Fatal().Err(err)

		}
		var respJSON UcallerResponse
		if err := json.NewDecoder(resp.Body).Decode(&respJSON); err != nil {
			log.Fatal().Err(err)

		}

		if !respJSON.Status {
			log.Fatal().Err(errUcaller)

		}
	}

}

func (u *Ucaller) GenerateCode() string {
	code := rand.Intn(9000) + 1000
	return strconv.Itoa(code)
}
