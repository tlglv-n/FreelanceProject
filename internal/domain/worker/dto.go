package worker

import (
	"errors"
	"net/http"
)

type Request struct {
	FullName    string `json:"fullname"`
	Pseudonym   string `json:"pseudonym"`
	Description string `json:"description"`
	Position    string `json:"position"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.FullName == "" {
		return errors.New("fullname: cannot be blank")
	}

	if s.Pseudonym == "" {
		return errors.New("pseudonym: cannot be blank")
	}

	if s.Description == "" {
		return errors.New("description: cannot be blank")
	}

	if s.Position == "" {
		return errors.New("position: cannot be blank")
	}

	return nil
}

type Response struct {
	ID          string `json:"id"`
	FullName    string `json:"fullname"`
	Pseudonym   string `json:"pseudonym"`
	Description string `json:"description"`
	Position    string `json:"position"`
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:          data.ID,
		FullName:    *data.FullName,
		Pseudonym:   *data.Pseudonym,
		Description: *data.Position,
		Position:    *data.Position,
	}
	return
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, object := range data {
		res = append(res, ParseFromEntity(object))
	}
	return
}
