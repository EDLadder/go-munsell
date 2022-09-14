package organization

import (
	"net/http"

	"github.com/EDLadder/go-munsell/organization_service/internal/apperror"
	"github.com/EDLadder/go-munsell/organization_service/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	organizationsURL = "/api/organizations"
)

type Handler struct {
	Logger              logging.Logger
	OrganizationService Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, organizationsURL, apperror.Middleware(h.CreateOrganization))
}

func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) error {
	return nil
}
