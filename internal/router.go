package internal

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/satishcg12/donate-me/internal/handler"
	"github.com/satishcg12/donate-me/internal/repository"
	"github.com/satishcg12/donate-me/view/pages"
)

// {
// 	"transaction_code": "0LD5CEH",
// 	"status": "COMPLETE",
// 	"total_amount": "1,000.0",
// 	"transaction_uuid": "240613-134231",
// 	"product_code": "NP-ES-ABHISHEK-EPAY",
// 	"signed_field_names": "transaction_code,status,total_amount,transaction_uuid,product_code,signed_field_names",
// 	"signature": "Mpwy0TFlHqpJjFUDGic+22mdoenITT+Ccz1LC61qMAc="
//   }

func (a *App) LoadRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		pages.Index().Render(r.Context(), w)
	})
	r.Get("/success/{donation_id}", func(w http.ResponseWriter, r *http.Request) {
		repo := repository.New(a.db)
		donatioinId := chi.URLParam(r, "donation_id")
		donatioinIdInt, _ := strconv.Atoi(donatioinId)
		res, err := repo.GetDonationByID(r.Context(), int64(donatioinIdInt))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))

			return
		}

		if res.Status != "COMPLETE" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Donation not found"))

			return
		}

		pages.ThankYou(pages.ThankYouStruct{
			Amount:  res.Amount,
			Name:    res.FullName,
			Email:   res.Email,
			Message: res.Message.String,
		}).Render(r.Context(), w)
	})

	r.Get("/admin/dashboard", func(w http.ResponseWriter, r *http.Request) {
		pages.AdminDashboard().Render(r.Context(), w)
	})

	r.Route("/api/v1/donation", a.loadDonationRoutes)
	r.Route("/api/v1/admin", a.loadAdminRoutes)

	a.router = r
}

func (a *App) loadDonationRoutes(r chi.Router) {
	donationRepo := repository.New(a.db)
	donationHandler := handler.NewDonationHandler(donationRepo)
	r.Post("/esewaform", donationHandler.EsewaForm)
	r.Get("/success/{donation_id}", donationHandler.DonationSuccess)
	r.Get("/failure/{donation_id}", donationHandler.DonationFail)
	r.Get("/list", donationHandler.ListDonations)
	r.Get("/total", donationHandler.TotalDonations)
	r.Get("/totalamount", donationHandler.TotalDonationsAmount)
	r.Get("/recent", donationHandler.RecentDonations)

	// r.Get("/", donationHandler.Get)
	// r.Delete("/", donationHandler.Delete)
}

func (a *App) loadAdminRoutes(r chi.Router) {

	// adminHandler := handler.NewAdminHandler()
	// r.Get("/dashboard", adminHandler.GetAdminDashboard)

}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
