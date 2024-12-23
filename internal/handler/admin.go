package handler

import (
	"net/http"

	"github.com/satishcg12/donate-me/view/pages"
)

type AdminHandler struct {
}

type AdminHandlerInterface interface {
	GetAdminDashboard(w http.ResponseWriter, r *http.Request)
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (a *AdminHandler) GetAdminDashboard(w http.ResponseWriter, r *http.Request) {
	pages.AdminDashboard().Render(r.Context(), w)
}
