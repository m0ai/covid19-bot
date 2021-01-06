package covid19info

import (
	"context"
	"gorm.io/gorm/clause"
	entity "scrapper/internal/entity"
	"scrapper/pkg/dbcontext"
	"time"
)

type Repository interface {
	Get(ctx context.Context, stateDt time.Time) (entity.Covid19InfoEntity, error)
	UpdateOrCreate(ctx context.Context, covid19infoArr entity.Covid19InfoEntity) error
}

// repository persists covid19info in database
type repository struct {
	db *dbcontext.DB
}

// NewRepository creates a new covid19info repository
func NewRepository(db *dbcontext.DB) Repository {
	return repository{db: db}
}

func (r repository) Get(ctx context.Context, stateDt time.Time) (entity.Covid19InfoEntity, error) {
	var obj entity.Covid19InfoEntity
	r.db.With(ctx).First(&obj, "state_dt <= ?", stateDt).Order("state_dt desc")
	return obj, nil
}

func (r repository) UpdateOrCreate(ctx context.Context, covid19infoArr entity.Covid19InfoEntity) error {
	r.db.With(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&covid19infoArr)
	return nil
}
