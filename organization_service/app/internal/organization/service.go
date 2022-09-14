package organization

import (
	"context"
	"errors"
	"fmt"

	"github.com/EDLadder/go-munsell/organization_service/internal/apperror"
	"github.com/EDLadder/go-munsell/organization_service/pkg/logging"
)

var _ Service = &service{}

type service struct {
	storage Storage
	logger  logging.Logger
}

func NewService(organizationStorage Storage, logger logging.Logger) (Service, error) {
	return &service{
		storage: organizationStorage,
		logger:  logger,
	}, nil
}

type Service interface {
	Create(ctx context.Context, dto CreateOrganizationDTO) (string, error)
}

func (s service) Create(ctx context.Context, dto CreateOrganizationDTO) (organizationUUID string, err error) {
	organization := NewOrganization(dto)
	organizationUUID, err = s.storage.Create(ctx, organization)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			return organizationUUID, err
		}
		return organizationUUID, fmt.Errorf("failed to create organization. error: %w", err)
	}

	return organizationUUID, nil
}
