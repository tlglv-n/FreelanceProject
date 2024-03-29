package http

import (
	"errors"
	"exchanger/internal/domain/customer"
	"exchanger/internal/service/hiring"
	"exchanger/pkg/market"
	"exchanger/pkg/server/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type CustomerHandler struct {
	hiringService *hiring.Service
}

func NewCustomerHandler(s *hiring.Service) *CustomerHandler {
	return &CustomerHandler{hiringService: s}
}

func (h *CustomerHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.list)
	r.Post("/", h.add)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})

	return r
}

// @Summary	list of customers from the repository
// @Tags		customers
// @Accept		json
// @Produce	json
// @Success	200			{array}		customer.Response
// @Failure	500			{object}	response.Object
// @Router		/customers 	[get]
func (h *CustomerHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.hiringService.ListCustomers(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// @Summary	add a new customer to the repository
// @Tags		customers
// @Accept		json
// @Produce	json
// @Param		request	body		customer.Request	true	"body param"
// @Success	200		{object}	customer.Response
// @Failure	400		{object}	response.Object
// @Failure	500		{object}	response.Object
// @Router		/customers [post]
func (h *CustomerHandler) add(w http.ResponseWriter, r *http.Request) {
	req := customer.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	res, err := h.hiringService.AddCustomer(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}
	response.OK(w, r, res)
}

// @Summary	get the customer from the repository
// @Tags		customers
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"path param"
// @Success	200	{object}	customer.Response
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/customers/{id} [get]
func (h *CustomerHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.hiringService.GetCustomer(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, market.ErrorNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}
		return
	}

	response.OK(w, r, res)
}

// @Summary	update the customer in the repository
// @Tags		customers
// @Accept		json
// @Produce	json
// @Param		id		path	int				true	"path param"
// @Param		request	body	customer.Request	true	"body param"
// @Success	200
// @Failure	400	{object}	response.Object
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/customers/{id} [put]
func (h *CustomerHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := customer.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	if err := h.hiringService.UpdateCustomer(r.Context(), id, req); err != nil {
		switch {
		case errors.Is(err, market.ErrorNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}
		return
	}
}

// @Summary	delete the customer from the repository
// @Tags		customers
// @Accept		json
// @Produce	json
// @Param		id	path	int	true	"path param"
// @Success	200
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/customers/{id} [delete]
func (h *CustomerHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.hiringService.DeleteCustomer(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, market.ErrorNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}
		return
	}
}
