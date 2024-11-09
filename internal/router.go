package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/satishcg12/dotnet-me/utils"
	"github.com/satishcg12/dotnet-me/view/components"
	"github.com/satishcg12/dotnet-me/view/pages"
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

type EsewaResponse struct {
	TransactionCode  string `json:"transaction_code"`
	Status           string `json:"status"`
	TotalAmount      string `json:"total_amount"`
	TransactionUUID  string `json:"transaction_uuid"`
	ProductCode      string `json:"product_code"`
	SignedFieldNames string `json:"signed_field_names"`
	Signature        string `json:"signature"`
}

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
	r.Post("/api/coffee", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			return
		}
		amount := r.FormValue("amount")
		custom := r.FormValue("custom")
		if custom != "0" {
			amount = custom
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")
		amountUint, err := strconv.ParseUint(amount, 10, 32)
		if err != nil {
			http.Error(w, "Invalid amount", http.StatusBadRequest)
			return
		}
		components.EsewaForm(name, email, message, uint32(amountUint*100)).Render(r.Context(), w)

	})
	r.Get("/api/v1/payment/success", func(w http.ResponseWriter, r *http.Request) {
		// get data from esewa in url params
		urlParams := r.URL.Query()
		data := urlParams.Get("data")
		if data == "" {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		//DECRYPT DATA
		data = utils.DecodeBase64(data)
		// JSON DATA TO STRUCT
		dataStruct := EsewaResponse{}
		err := json.Unmarshal([]byte(data), &dataStruct)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		// VERIFY SIGNATURE
		secretKey := "8gBm/:&EnhH.1/q"
		signature := dataStruct.Signature
		signedFieldNames := dataStruct.SignedFieldNames

		fmt.Println(signature)
		fmt.Println(signedFieldNames)
		if !utils.VerifySignature(secretKey, signedFieldNames, signature) {
			http.Error(w, "Invalid signature", http.StatusBadRequest)
			return
		}

		// SAVE TO DATABASE
		// SEND EMAIL
		// RENDER THANK YOU PAGE
		pages.ThankYou().Render(r.Context(), w)
	})

	// frontend
	r.Get("/frontend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend/index.html")
	})

	r.Route("/api/v1/donation", a.loadDonationRoutes)

	// r.Route("/url", a.LoadURLRoutes)

	a.router = r
}

func (a *App) loadDonationRoutes(r chi.Router) {
	// donationRepo := repository.New(a.db)
	// donationHandler := handlers.NewDonationHandler(donationRepo)
	// r.Get("/", donationHandler.Get)
	// r.Post("/", donationHandler.Create)
	// r.Delete("/", donationHandler.Delete)
}

// func (a *App) LoadURLRoutes(r chi.Router) {
// 	urlRepo := repository.New(a.db)
// 	urlHandler := handlers.NewURLHandler(urlRepo)
// 	r.Get("/", urlHandler.Get)
// 	r.Post("/", urlHandler.Create)
// 	r.Delete("/", urlHandler.Delete)
// 	r.Get("/{shortUrl}", urlHandler.Redirect)
// }

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
