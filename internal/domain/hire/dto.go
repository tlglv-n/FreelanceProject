package hire

import (
	"errors"
	"net/http"
)

type Request struct {
	JobName     string `json:"jobname"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Position    string `json:"position"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.JobName == "" {
		return errors.New("jobname: cannot be blank")
	}

	if s.Amount == 0 {
		return errors.New("amount: cannot be blank")
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
	JobName     string `json:"jobname"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Position    string `json:"position"`
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:          data.ID,
		JobName:     *data.JobName,
		Amount:      *data.Amount,
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
