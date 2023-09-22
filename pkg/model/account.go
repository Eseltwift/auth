package model

import (
	"cli-go/internal/api"
	"fmt"
)

type Account struct {
	Email     string    `json:"email" binding:"required"`
	UserID    int64     `json:"userID" binding:"required"`
	App       api.App   `json:"app" binding:"required"`
	MasterKey EncString `json:"masterKey" binding:"required"`
	SecretKey EncString `json:"secretKey" binding:"required"`
	Token     EncString `json:"token" binding:"required"`
}

func (a Account) AccountKey() string {
	return fmt.Sprintf("%s-%d", a.App, a.UserID)
}

func (a Account) DataBucket() string {
	return fmt.Sprintf("%s-%d-data", a.App, a.UserID)
}
