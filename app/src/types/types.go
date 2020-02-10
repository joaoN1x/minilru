package types

import (
	"github.com/joaoN1x/minilru/src/db"
)

type MessageOut struct {
	Status string         `json:"status,omitempty"`
	Code   int            `json:"code,omitempty"`
	Data   MessageOutData `json:"data,omitempty"`
}

type MessageOutData struct {
	Detail   string        `json:"detail,omitempty"`
	Affected int           `json:"affected,omitempty"`
	Url      []db.Url      `json:"url,omitempty"`
	UrlStats []db.UrlStats `json:"url_stats,omitempty"`
}
