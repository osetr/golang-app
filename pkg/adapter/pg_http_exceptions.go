package adapter

import "github.com/go-pg/pg/v10"

func NewHttpExcAdapter() *HttpExcAdapter {
	return &HttpExcAdapter{}
}

type HttpExcAdapter struct {
}

func (*HttpExcAdapter) Transform(e error) (error, string, int) {
	switch e {
	case pg.ErrNoRows:
		return e, "not found", 404
	default:
		return e, "internal server error", 500
	}
}
