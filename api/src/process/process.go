package process

import (
	"fmt"
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
	fmt.Printf("\n\n\n cal= %v", hd.Salt)

	h, _ := hashids.NewWithData(hd)
	id, _ := h.Encode([]int{int(urlId)})

	fmt.Printf("\n\n\n id= %v", id)

	return id, true
}
