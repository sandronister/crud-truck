package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/usecase"
)

type TruckHandler struct {
	usecase *usecase.TruckUseCase
}

func NewTruckHandler(usecase *usecase.TruckUseCase) *TruckHandler {
	return &TruckHandler{usecase}
}

func (h *TruckHandler) Delete(c echo.Context) error {
	truckID := c.Param("id")

	err := h.usecase.Delete(truckID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TruckHandler) FindAll(c echo.Context) error {
	trucks, err := h.usecase.FindAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, trucks)
}

func (h *TruckHandler) FindByID(c echo.Context) error {
	truckID := c.Param("id")

	truck, err := h.usecase.FindByID(truckID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, truck)
}

func (h *TruckHandler) Save(c echo.Context) error {
	truck := new(dto.Truck)
	if err := c.Bind(truck); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	err := h.usecase.Save(truck)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, truck)
}

func (h *TruckHandler) Update(c echo.Context) error {
	var truck dto.Truck
	if err := c.Bind(&truck); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := c.Param("id")

	truck.ID = id
	if err := h.usecase.Update(&truck); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
