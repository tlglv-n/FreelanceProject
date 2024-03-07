package customer

import (
	"errors"
	"net/http"
)

type Request struct {
	FullName  string `json:"fullName"`
	Pseudonym string `json:"pseudonym"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.FullName == "" {
		return errors.New("fullname: cannot be blank")
	}

	if s.Pseudonym == "" {
		return errors.New("pseudonym: cannot be blank")
	}

	return nil
}

type Response struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	Pseudonym string `json:"pseudonym"`
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:        data.ID,
		FullName:  *data.FullName,
		Pseudonym: *data.Pseudonym,
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
