package process

import (
	"os"

	"github.com/speps/go-hashids"
)

func GenerateShortUrl(urlId int64) (string, bool) {

	hd := hashids.NewData()

	serviceSalt, ok := os.LookupEnv("SERVICE_SALT")
	if !ok {
		return "", false
	}
	hd.Salt = serviceSalt

	h, _ := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{int(urlId)})

	return id, true
}
