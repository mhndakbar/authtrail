package main

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/mhndakbar/authtrails/internal/database"
)

type AuthTrail struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	TrailType string    `json:"trail_type"`
}

var (
	trailTypes = []string{"login", "logout", "signup"}
)

func (apiCfg *apiConfig) CreateAuthTrail(userID uuid.UUID, trailType string, ctx context.Context) (AuthTrail, error) {
	if !slices.Contains(trailTypes, trailType) {
		return AuthTrail{}, errors.New("invalid trail type")
	}

	id := uuid.New()
	err := apiCfg.DB.CreateAuthTrail(ctx, database.CreateAuthTrailParams{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userID,
		Type:      trailType,
	})
	if err != nil {
		return AuthTrail{}, errors.New("error creating auth trail")
	}

	return AuthTrail{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}

func (apiCfg *apiConfig) handlerGetAuthTrail(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	authTrails, err := apiCfg.DB.GetAuthTrailsForUser(r.Context(), uuid.MustParse(userID))
	if err != nil {
		http.Error(w, "Error getting auth trails", http.StatusInternalServerError)
		return
	}

	authTrailResponse := []AuthTrail{}
	for _, authTrail := range authTrails {
		authTrailResponse = append(authTrailResponse, AuthTrail{
			ID:        authTrail.ID,
			CreatedAt: authTrail.CreatedAt,
			UpdatedAt: authTrail.UpdatedAt,
			UserID:    authTrail.UserID,
			TrailType: authTrail.Type,
		})
	}

	responedWithJson(w, http.StatusOK, authTrailResponse)
}
