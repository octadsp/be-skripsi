package routes

import (
	"be-skripsi/handlers"
	"be-skripsi/pkg/pg"
	"be-skripsi/repositories"

	"github.com/labstack/echo/v4"
)

func PriceListRoutes(e *echo.Group) {
	priceListRepository := repositories.RepositoryPriceList(pg.DB)
	h := handlers.HandlerPriceList(priceListRepository)

	e.GET("/price-lists", h.FindPriceLists)
	e.GET("/price-all-lists", h.FindAllPriceLists)
	e.GET("/price-list/:id", h.GetPriceList)
	e.POST("/price-list", h.AddPriceList)
	e.PATCH("/price-list/:id", h.UpdatePriceList)
	e.DELETE("/price-list/:id", h.DeletePriceList)
}
