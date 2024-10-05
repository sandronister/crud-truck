package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/usecase"
)

type LinkHandler struct {
	usecase *usecase.LinkUseCase
}

func NewLinkHandler(usecase *usecase.LinkUseCase) *LinkHandler {
	return &LinkHandler{
		usecase: usecase,
	}
}

func (h *LinkHandler) Delete(c echo.Context) error {
	driverId := c.Param("driver_id")
	truckId := c.Param("truck_id")
	err := h.usecase.Delete(driverId, truckId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *LinkHandler) ListByDriver(c echo.Context) error {
	driverId := c.Param("driver_id")
	list, err := h.usecase.ListByDriver(driverId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, list)
}

func (h *LinkHandler) Save(c echo.Context) error {
	var linkDTO dto.Link

	if err := c.Bind(&linkDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.usecase.Save(&linkDTO)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, linkDTO)
}
