package engine

import (
	"context"
	"database/sql"
	"github.com/Kandy-Hamisi/watch-ndinga/models"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) EngineById(ctx context.Context, id string) (models.Engine, error) {

}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) EngineDelete(ctx context.Context, carReq *models.CarRequest) (models.Engine, error) {

}
