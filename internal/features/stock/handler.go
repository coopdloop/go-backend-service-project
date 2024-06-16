package stock

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Handler interface {
		// Search : GET /stock
		Search(w http.ResponseWriter, r *http.Request)
		// Create : POST /stock
		Create(w http.ResponseWriter, r *http.Request)
		// Update : PATCH /todos/{todoId}
		// Update : POST /todos/{todoId}/edit
		// Update(w http.ResponseWriter, r *http.Request)
		// Get : GET /todos/{todoId}
		// Get(w http.ResponseWriter, r *http.Request)
		// Delete : DELETE /todos/{todoId}
		// Delete : POST /todos/{todoId}/delete
		// Delete(w http.ResponseWriter, r *http.Request)
		// Sort : POST /todos/sort
		// Sort(w http.ResponseWriter, r *http.Request)
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(r chi.Router, h Handler) {
	r.Route("/stock", func(r chi.Router) {
		r.Get("/", h.Search)
		r.Post("/", h.Create)
	})
}

func (h handler) Search(w http.ResponseWriter, r *http.Request) {
	var search = r.URL.Query().Get("search")
	stock, err := h.service.Search(r.Context(), search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(stock)
	stockBytes, err := json.Marshal(stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(stockBytes)
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	// if err := r.ParseForm(); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// var description = r.Form.Get("description")
	var ticker = r.URL.Query().Get("ticker")

	msg, err := h.service.Add(r.Context(), ticker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(msg))

	// switch isHTMX(r) {
	// case true:
	// 	err = partials.RenderTodo(todo).Render(r.Context(), w)
	// default:
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// }
	//
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
