package model

import "github.com/google/uuid"

const (
	Corporations       CompanyType = "Corporations"
	NonProfit          CompanyType = "Non Profit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"

	CompanyTable string = "company"
)

type Company struct {
	Id                uuid.UUID   `json:"id"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	AmountOfEmployees int         `json:"amountOfEmployees"`
	Registered        bool        `json:"registered"`
	CompanyType       CompanyType `json:"companyType"`
}

type CompanyType string
