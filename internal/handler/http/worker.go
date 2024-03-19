package http

import (
	"errors"
	"exchanger/internal/domain/worker"
	"exchanger/internal/service/hiring"
	"exchanger/pkg/market"
	"exchanger/pkg/server/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type WorkerHandler struct {
	hiringService *hiring.Service
}

func NewWorkerService(s *hiring.Service) *WorkerHandler {
	return &WorkerHandler{hiringService: s}
}

func (h *WorkerHandler) Routes() chi.Router {
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

// @Summary	list of workers from the repository
// @Tags		workers
// @Accept		json
// @Produce	json
// @Success	200			{array}		worker.Response
// @Failure	500			{object}	response.Object
// @Router		/customers 	[get]
func (h *WorkerHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.hiringService.ListWorkers(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// @Summary	add a new worker to the repository
// @Tags		workers
// @Accept		json
// @Produce	json
// @Param		request	body		worker.Request	true	"body param"
// @Success	200		{object}	worker.Response
// @Failure	400		{object}	response.Object
// @Failure	500		{object}	response.Object
// @Router		/customers [post]
func (h *WorkerHandler) add(w http.ResponseWriter, r *http.Request) {
	req := worker.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	res, err := h.hiringService.AddWorker(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}
	response.OK(w, r, res)
}

// @Summary	get the worker from the repository
// @Tags		workers
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"path param"
// @Success	200	{object}	worker.Response
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/workers/{id} [get]
func (h *WorkerHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.hiringService.GetWorker(r.Context(), id)
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

// @Summary	update the worker in the repository
// @Tags		workers
// @Accept		json
// @Produce	json
// @Param		id		path	int				true	"path param"
// @Param		request	body	worker.Request	true	"body param"
// @Success	200
// @Failure	400	{object}	response.Object
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/workers/{id} [put]
func (h *WorkerHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := worker.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	if err := h.hiringService.UpdateWorker(r.Context(), id, req); err != nil {
		switch {
		case errors.Is(err, market.ErrorNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}
		return
	}
}

// @Summary	delete the worker from the repository
// @Tags		workers
// @Accept		json
// @Produce	json
// @Param		id	path	int	true	"path param"
// @Success	200
// @Failure	404	{object}	response.Object
// @Failure	500	{object}	response.Object
// @Router		/customers/{id} [delete]
func (h *WorkerHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.hiringService.DeleteWorker(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, market.ErrorNotFound):
			response.NotFound(w, r, err)
		default:
			response.InternalServerError(w, r, err)
		}
		return
	}
}
