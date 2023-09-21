package rest

import (
	"companies/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"net/http"
)

type CompaniesController interface {
	Routes(rg gin.RouterGroup)
	GetCompany(c *gin.Context)
	CreateCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
}

type CompaniesRestController struct {
	*gorm.DB
}

func (crc *CompaniesRestController) Routes(rg gin.RouterGroup) {
	rg.GET("/company/:id", crc.GetCompany)
	rg.POST("/company", crc.CreateCompany)
	rg.PATCH("/company/", crc.UpdateCompany)
	rg.DELETE("/company/:id", crc.DeleteCompany)
}

func (crc *CompaniesRestController) GetCompany(c *gin.Context) {
	id := c.Param("id")
	var uid uuid.UUID
	var err error

	if uid, err = uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id could not be parsed into uuid"})
		return
	}

	company := model.Company{}

	crc.Table(model.CompanyTable).First(&company, "id = ?", uid)

	if company.Id != uid {
		c.JSON(http.StatusNotFound, gin.H{"message": "could not be found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (crc *CompaniesRestController) CreateCompany(c *gin.Context) {
	company := model.Company{}

	err := c.BindJSON(&company)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "body could not be parsed"})
		return
	}

	companyToCheck := model.Company{}

	crc.Table(model.CompanyTable).Find(&companyToCheck, "name = ?", company.Name)
	if companyToCheck.Name != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "company with this name already created"})
		return
	}

	crc.Table(model.CompanyTable).Create(&company)

	c.JSON(http.StatusCreated, company)
}

func (crc *CompaniesRestController) UpdateCompany(c *gin.Context) {
	company := model.Company{}

	err := c.BindJSON(&company)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "body could not be parsed"})
		return
	}

	companyToCheck := model.Company{}

	crc.Table(model.CompanyTable).Find(&companyToCheck, "id = ?", company.Id)
	if companyToCheck.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there is no company with such id"})
		return
	}

	crc.Table(model.CompanyTable).Save(&company)

	c.JSON(http.StatusOK, company)
}

func (crc *CompaniesRestController) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	var uid uuid.UUID
	var err error

	if uid, err = uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id could not be parsed into uuid"})
		return
	}

	company := model.Company{}

	crc.Table(model.CompanyTable).First(&company, "id = ?", uid)

	if company.Id != uid {
		c.JSON(http.StatusNotFound, gin.H{"message": "could not be found"})
		return
	}

	crc.Table(model.CompanyTable).Delete(&company)

	c.JSON(http.StatusNoContent, "")
}
