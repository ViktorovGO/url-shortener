package deleter

import (
	"errors"
	"log/slog"
	"net/http"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UrlDeleter interface {
	DeleteUrl(alias string) error
}

func New(log *slog.Logger, deleter UrlDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.delete.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("alias is empty"))
			return
		}

		err := deleter.DeleteUrl(alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Info("url not found")
				render.JSON(w, r, resp.Error("url not found"))
				return
			}
			log.Info("failed to delete url")
			render.JSON(w, r, resp.Error("failed to delete url"))
			return

		}
		render.JSON(w, r, resp.OK())

	}

}
