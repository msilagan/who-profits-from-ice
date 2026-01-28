package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/msilagan/who-profits-from-ice/backend/internal/db"
	"github.com/msilagan/who-profits-from-ice/backend/internal/models"
)

func GetEntityByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid entity ID", http.StatusBadRequest)
		return
	}

	var entity models.Entity
	query := "SELECT id, name, type FROM entities WHERE id=$1"
	err = db.Pool.QueryRow(context.Background(), query, id).Scan(&entity.ID, &entity.Name, &entity.Type)
	if err != nil {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity)
}
