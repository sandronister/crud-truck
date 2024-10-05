package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sandronister/crud-truck/internal/dto"
	"github.com/sandronister/crud-truck/internal/usecase"
)

type DriverHandler struct {
	usecase *usecase.DriverUsecase
}

func NewDriverHandler(usecase *usecase.DriverUsecase) *DriverHandler {
	return &DriverHandler{
		usecase: usecase,
	}
}

func (h *DriverHandler) Delete(c echo.Context) error {
	driverID := c.Param("id")

	err := h.usecase.Delete(driverID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *DriverHandler) FindAll(c echo.Context) error {
	drivers, err := h.usecase.FindAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, drivers)
}

func (h *DriverHandler) FindByID(c echo.Context) error {
	driverID := c.Param("id")

	driver, err := h.usecase.FindByID(driverID)

	if driver == nil {
		return c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})

	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	if driver == nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, driver)
}

func (h *DriverHandler) Save(c echo.Context) error {
	var driverDTO dto.Driver

	if err := c.Bind(&driverDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	driver, err := h.usecase.Save(driverDTO.Name, driverDTO.LicenseID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, driver)
}

func (h DriverHandler) Update(c echo.Context) error {
	var driverDTO dto.Driver

	if err := c.Bind(&driverDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")

	driverDTO.ID = id

	err := h.usecase.Update(&driverDTO)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
