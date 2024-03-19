package http

import (
	"errors"
	"exchanger/internal/domain/hire"
	"exchanger/internal/service/hiring"
	"exchanger/pkg/market"
	"exchanger/pkg/server/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type HireHandler struct {
	hiringService *hiring.Service
}

func NewHireHandler(s *hiring.Service) *HireHandler {
	return &HireHandler{hiringService: s}
}

func (h *HireHandler) Routes() chi.Router {
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

// @Summary	list of hires from the repository
// @Tags		hires
// @Accept		json
// @Produce	json
// @Success	200			{array}		hire.Response
// @Failure	500			{object}	response.Object
// @Router		/hires 	[get]
func (h *HireHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.hiringService.ListHires(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// @Summary	add a new hire to the repository
// @Tags		hires
// @Accept		json
// @Produce	json
// @Param		request	body		hire.Request	true	"body param"
// @Success	200		{object}	hire.Response
// @Failure	400		{object}	response.Object
// @Failure	500		{object}	response.Object
// @Router		/hires [post]
func (h *HireHandler) add(w http.ResponseWriter, r *http.Request) {
	req := hire.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	res, err := h.hiringService.AddHire(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}
	response.OK(w, r, res)
}

// @Summary	get the hire from the repository
// @Tags		hires
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"path param"
// @Success	200	{object}	hire.Response
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/hires/{id} [get]
func (h *HireHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.hiringService.GetHire(r.Context(), id)
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

// @Summary	update the hire in the repository
// @Tags		hires
// @Accept		json
// @Produce	json
// @Param		id		path	int				true	"path param"
// @Param		request	body	hire.Request	true	"body param"
// @Success	200
// @Failure	400	{object}	response.Object
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/hires/{id} [put]
func (h *HireHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := hire.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	if err := h.hiringService.UpdateHire(r.Context(), id, req); err != nil {
		switch {
		case errors.Is(err, market.ErrorNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}
		return
	}
}

// @Summary	delete the hire from the repository
// @Tags		hires
// @Accept		json
// @Produce	json
// @Param		id	path	int	true	"path param"
// @Success	200
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/hires/{id} [delete]
func (h *HireHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.hiringService.DeleteHire(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, market.ErrorNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}
		return
	}
}
