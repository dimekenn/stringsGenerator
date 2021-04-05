package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"main/config"
	"main/model"
	"math/rand"
	"net/http"
	"time"
)

//Handler
type Handler struct {
	cfg config.Config
}

func NewHandler(cfg config.Config) *Handler{
	return &Handler{
		cfg: cfg,
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (h *Handler) GeneratorHandler(w http.ResponseWriter, r *http.Request)  {
	var requestModel model.NumberRequest
	var stringsRequest model.StringsRequest

	err := json.NewDecoder(r.Body).Decode(&requestModel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var randomStrings []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= requestModel.NumberOfStrings; i++ {
		randomStrings = append(randomStrings, randSeq(rand.Intn(requestModel.NumberOfStrings)))
	}

	stringsRequest.Strings = randomStrings
	b, err := json.Marshal(stringsRequest)

	request, err := http.NewRequest("POST", h.cfg.Host, bytes.NewBuffer(b))
	if err != nil {
		log.Fatal("error with new request")
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error with client request: s%", err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	err =json.Unmarshal(data, &stringsRequest)

	err = Response(w, stringsRequest)
}

func Response(res http.ResponseWriter, response interface{}) error {
	res.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(res).Encode(response)
	return nil
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}