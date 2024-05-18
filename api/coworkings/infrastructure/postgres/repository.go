package postgres

import (
	domain "api/coworkings/domain"
	"api/utils/logger"
	r "api/utils/result"
	"errors"

	"gorm.io/gorm"
)

type CoworkingRepositoryPostgres struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CoworkingRepositoryPostgres {
	return &CoworkingRepositoryPostgres{
		db: db,
	}
}

func (repository *CoworkingRepositoryPostgres) SaveCoworking(coworking *domain.Coworking) r.Result[domain.Coworking] {
	postgresCoworking := MapCoworkingToPostgres(*coworking)

	var existingCoworking PostgresCoworking
	result := repository.db.Where("uuid = ?", postgresCoworking.UUID).First(&existingCoworking)
	if result.Error == nil {
		postgresCoworking.ID = existingCoworking.ID
		postgresCoworking.CreatedAt = existingCoworking.CreatedAt
		logger.GetLogger().Info("UPDATING", len(postgresCoworking.Images))
		result = repository.db.Save(&postgresCoworking)
	} else {
		logger.GetLogger().Info("RECREATING", nil)
		result = repository.db.Create(&postgresCoworking)
	}

	if result.Error != nil {
		return r.NewResult(domain.Coworking{}, errors.New(result.Error.Error()))
	}

	return r.NewResult(*coworking, nil)
}

func (repository *CoworkingRepositoryPostgres) GetCoworkings() r.Result[[]domain.Coworking] {
	var postgresCoworkings []PostgresCoworking

	switch err := repository.db.Preload("Workspaces").Preload("Images").Find(&postgresCoworkings).Error; {
	case err != nil:
		return r.NewResult(make([]domain.Coworking, 0), errors.New(err.Error()))

	default:
		domainCoworkings := MapPostgresCoworkingListToDomain(postgresCoworkings)
		return r.NewResult(domainCoworkings, nil)
	}
}

func (repository *CoworkingRepositoryPostgres) GetCoworkingByUuid(coworkingUuid *domain.CoworkingUUID) r.Result[domain.Coworking] {
	var postgresCoworking PostgresCoworking

	result := repository.db.Where("UUID = ?", coworkingUuid.Value).First(&postgresCoworking)

	if result.Error != nil {
		return r.NewResult(domain.Coworking{}, result.Error)
	}

	domainCoworking := MapPostgresCoworkingToDomain(postgresCoworking)
	return r.NewResult(domainCoworking, nil)
}
