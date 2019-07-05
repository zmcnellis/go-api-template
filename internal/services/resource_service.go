package services

import (
	"github.com/jinzhu/gorm"
	"github.com/zmcnellis/go-api-template/internal/models"
)

// ResourceService ...
type ResourceService struct {
	DB *gorm.DB
}

// NewResourceService ...
func NewResourceService(db *gorm.DB) *ResourceService {
	return &ResourceService{DB: db}
}

// GetResources ...
func (s *ResourceService) GetResources() *[]models.Resource {
	var resources []models.Resource
	s.DB.Find(&resources)
	return &resources
}

// GetResource ...
func (s *ResourceService) GetResource(id int) *models.Resource {
	var resource models.Resource
	s.DB.First(&resource, id)
	return &resource
}

// CreateResource ...
func (s *ResourceService) CreateResource(resource *models.Resource) {
	s.DB.Create(&resource)
}

// DeleteResource ...
func (s *ResourceService) DeleteResource(id int) {
	var resource models.Resource
	s.DB.First(&resource, id)
	s.DB.Delete(&resource)
}
