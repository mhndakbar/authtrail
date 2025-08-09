package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mhndakbar/authtrails/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, "Error decoding parameters", http.StatusBadRequest)
		return
	}

	apiKey, err := generateRandomSHA256Hash()
	if err != nil {
		http.Error(w, "Error generating API key", http.StatusInternalServerError)
		return
	}

	_, err = apiCfg.DB.GetUserByName(r.Context(), params.Name)
	if err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	err = apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		ApiKey:    apiKey,
		Password:  params.Password,
	})
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), apiKey)
	if err != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	_, err = apiCfg.CreateAuthTrail(user.ID, "signup", r.Context())
	if err != nil {
		http.Error(w, "Error creating auth trail", http.StatusInternalServerError)
		return
	}

	userResponse := databaseUserToUser(user)

	responedWithJson(w, http.StatusCreated, userResponse)
}

func (apiCfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request, user database.User) {
	_, err := apiCfg.CreateAuthTrail(user.ID, "login", r.Context())
	if err != nil {
		http.Error(w, "Error creating auth trail", http.StatusInternalServerError)
		return
	}

	userResponse := databaseUserToUser(user)

	responedWithJson(w, http.StatusOK, userResponse)
}

func generateRandomSHA256Hash() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(randomBytes)
	hashString := hex.EncodeToString(hash[:])
	return hashString, nil
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	userResponse := databaseUserToUser(user)

	responedWithJson(w, http.StatusOK, userResponse)
}
