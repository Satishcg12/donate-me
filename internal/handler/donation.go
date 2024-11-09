package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/satishcg12/dotnet-me/internal/repository"
	"github.com/satishcg12/dotnet-me/utils"
	"github.com/satishcg12/dotnet-me/view/components"
	"github.com/satishcg12/dotnet-me/view/pages"
)

type DonationHandler struct {
	repo *repository.Queries
}

type DonationHanderInterface interface {
	// will response html components
	EsewaForm(w http.ResponseWriter, r *http.Request)
	DonationSuccess(w http.ResponseWriter, r *http.Request)
	DonationFail(w http.ResponseWriter, r *http.Request)
}
type FormError struct {
	Amount  string `json:"amount"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type EsewaResponse struct {
	TransactionCode  string `json:"transaction_code"`
	Status           string `json:"status"`
	TotalAmount      string `json:"total_amount"`
	TransactionUUID  string `json:"transaction_uuid"`
	ProductCode      string `json:"product_code"`
	SignedFieldNames string `json:"signed_field_names"`
	Signature        string `json:"signature"`
}

func NewDonationHandler(repo *repository.Queries) *DonationHandler {
	return &DonationHandler{
		repo: repo,
	}
}

func (d *DonationHandler) EsewaForm(w http.ResponseWriter, r *http.Request) {

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
	errMsg := FormError{}
	// check the length of the name, email and message
	if len(name) > 100 {
		errMsg.Name = "Name is too long"
	}
	if len(email) > 100 {
		errMsg.Email = "Email is too long"
	}
	if len(message) > 1000 {
		errMsg.Message = "Message is too long"
	}

	// check if name is all alphabets
	if !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(name) {
		errMsg.Name = "Name should only contain alphabets"
	}
	// check if name has space
	if !regexp.MustCompile(`^[a-zA-Z]+ [a-zA-Z]+$`).MatchString(name) {
		errMsg.Name = "Name should contain first name and last name"
	}
	// check valid email
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		errMsg.Email = "Invalid email"
	}

	// check if the name, email and message is empty
	if name == "" {
		errMsg.Name = "Name is required"
	}
	if email == "" {
		errMsg.Email = "Email is required"
	}
	if message == "" {
		errMsg.Message = "Message is required"
	}

	log.Println(errMsg)
	// convert to json and strignify
	if errMsg.Name != "" || errMsg.Email != "" || errMsg.Message != "" {
		res, _ := json.Marshal(errMsg)
		//convert to string
		strRes := string(res)
		log.Println(strRes)
		components.Mainform(name, email, message, uint32(amountUint), strRes).Render(r.Context(), w)
		return
	}

	res, err := d.repo.CreateDonation(r.Context(), repository.CreateDonationParams{
		FullName: name,
		Email:    email,
		Message: sql.NullString{
			String: message,
			Valid:  true,
		},
		Amount: int64(amountUint * 100),
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Println(res)

	components.EsewaForm(strconv.Itoa(int(res.ID)), name, email, message, uint32(amountUint*100)).Render(r.Context(), w)
}

func (d *DonationHandler) DonationSuccess(w http.ResponseWriter, r *http.Request) {
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
	// transaction_code,status,total_amount,transaction_uuid,product_code,signed_field_names
	signedFieldNames := fmt.Sprintf("transaction_code=%s,status=%s,total_amount=%s,transaction_uuid=%s,product_code=%s,signed_field_names=%s", dataStruct.TransactionCode, dataStruct.Status, dataStruct.TotalAmount, dataStruct.TransactionUUID, dataStruct.ProductCode, dataStruct.SignedFieldNames)

	log.Println(signature)
	log.Println(signedFieldNames)
	log.Println(utils.VerifySignature(secretKey, signedFieldNames, signature))
	if !utils.VerifySignature(secretKey, signedFieldNames, signature) {
		http.Error(w, "Invalid signature", http.StatusBadRequest)
		return
	}

	// GET DONATION ID
	donationId := chi.URLParam(r, "donation_id")
	if donationId == "" {
		http.Error(w, "Invalid donation id", http.StatusBadRequest)
		return
	}
	donationIdInt, err := strconv.Atoi(donationId)
	if err != nil {
		http.Error(w, "Invalid donation id", http.StatusBadRequest)
		return
	}

	// UPDATE DONATION STATUS
	_, err = d.repo.UpdateStatus(r.Context(), repository.UpdateStatusParams{
		ID:     int64(donationIdInt),
		Status: dataStruct.Status,
	})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// SEND EMAIL

	// Redirect to success page
	pages.RedirectPage("http://localhost:3000/success/"+donationId, "Donation Success").Render(r.Context(), w)
}

func (d *DonationHandler) DonationFail(w http.ResponseWriter, r *http.Request) {

}
