package redirect

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

//go:generate go run github.com/vektra/mockery/v2@v2.40.1 --name=UrlGetter
type UrlGetter interface {
	GetUrl(alias string) (string, error)
}

func New(log *slog.Logger, getter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.redirect.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("not found"))
			return
		}

		url, err := getter.GetUrl(alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				w.WriteHeader(http.StatusNotFound)
				log.Info("url with alias not found", "alias", alias)
				render.JSON(w, r, resp.Error("not found"))
				return
			}

			log.Info("error")
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, resp.Error("internal error"))
			return

		}

		log.Info("got url", slog.String("url", url))

		http.Redirect(w, r, url, http.StatusFound)

	}
}
