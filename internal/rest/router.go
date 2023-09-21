package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Router struct {
	rg gin.RouterGroup
	db *gorm.DB
	cc CompaniesController
}

func NewRouter(rg gin.RouterGroup, db *gorm.DB) *Router {
	return &Router{rg: rg, db: db}
}

func (r *Router) RouteControllers() {
	r.cc = &CompaniesRestController{r.db}
	r.cc.Routes(r.rg)
}
