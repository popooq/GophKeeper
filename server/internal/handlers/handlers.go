package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gtihub.com/popooq/Gophkeeper/server/internal/services"
	"gtihub.com/popooq/Gophkeeper/server/internal/storage"
)

type Handlers struct {
	keeper storage.Keeper
	//secret string
}

func New(keeper storage.Keeper) *Handlers {
	return &Handlers{
		keeper: keeper,
	}
}

func (h *Handlers) Route() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/server/registration", h.Registration)
	e.POST("/server/login", h.Login)

	logged := e.Group("/entry") //, echojwt.WithConfig(echojwt.Config{SigningKey: []byte(h.secret)}))

	logged.POST("/new", h.newEntry)
	logged.POST("/update", h.updateEntry)
	logged.GET("/get/:username/:servicename", h.getEntry)
	return e
}

func (h *Handlers) Registration(c echo.Context) error {
	return nil
}

func (h *Handlers) Login(c echo.Context) error {
	return nil
}

func (h *Handlers) newEntry(c echo.Context) error {
	status, err := services.NewEntry(c.Request().Body, h.keeper)
	if status != http.StatusOK {
		c.Response().Writer.WriteHeader(status)
		err := fmt.Errorf("http status is not OK %s", err)
		return err
	}
	c.Response().Writer.WriteHeader(status)
	c.Response().Writer.Write([]byte("entry added"))

	return nil
}

func (h *Handlers) updateEntry(c echo.Context) error {
	resp, status := services.UpdateEntry(c.Request().Body, h.keeper)
	if status != http.StatusOK {
		c.Response().Writer.WriteHeader(status)
		c.Response().Writer.Write(resp)

		err := fmt.Errorf("http status is not OK")
		return err
	}
	c.Response().Writer.WriteHeader(status)
	c.Response().Writer.Write(resp)

	return nil
}

func (h *Handlers) getEntry(c echo.Context) error {
	entry, status := services.GetEntry(c.Param("username"), c.Param("servicename"), h.keeper)
	if status != http.StatusOK {
		c.Response().Writer.WriteHeader(status)
		err := fmt.Errorf("http status is not OK")
		return err
	}
	c.Response().Writer.WriteHeader(status)
	c.Response().Writer.Write(entry)

	return nil
}
