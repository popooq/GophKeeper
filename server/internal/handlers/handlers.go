package handlers

import (
	"fmt"
	"log"
	"net/http"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gtihub.com/popooq/Gophkeeper/server/internal/services"
	"gtihub.com/popooq/Gophkeeper/server/internal/storage"
)

type Handlers struct {
	keeper storage.Keeper
	secret []byte
}

func New(keeper storage.Keeper, secret string) *Handlers {
	return &Handlers{
		keeper: keeper,
		secret: []byte(secret),
	}
}

func (h *Handlers) Route() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/server/registration", h.Registration)
	e.POST("/server/login", h.Login)

	logged := e.Group("/entry", echojwt.WithConfig(echojwt.Config{SigningKey: []byte(h.secret)}))

	logged.POST("/new", h.newEntry)
	logged.POST("/update", h.updateEntry)
	logged.GET("/get/:username/:servicename", h.getEntry)
	logged.DELETE("/delete/:username/:servicename", h.deleteEntry)
	return e
}

func (h *Handlers) Registration(c echo.Context) error {
	status, err := services.Registration(c.Request().Body, h.keeper)
	if err != nil {
		c.Response().Writer.WriteHeader(status)
		err := fmt.Errorf("http status is not OK %s", err)
		return err
	}
	c.Response().Writer.WriteHeader(status)
	c.Response().Writer.Write([]byte("registration complete"))

	return err
}

func (h *Handlers) Login(c echo.Context) error {
	jwt, status, err := services.Login(c.Request().Body, h.secret, h.keeper)
	if err != nil {
		log.Println(status, err)
		c.Response().Writer.WriteHeader(status)
		err := fmt.Errorf("http status is not OK %s", err)
		return err
	}
	c.Response().Writer.WriteHeader(status)
	c.Response().Writer.Write([]byte(jwt))

	return err
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

	return err
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

func (h *Handlers) deleteEntry(c echo.Context) error {
	status := services.DeleteEntry(c.Param("username"), c.Param("servicename"), h.keeper)
	c.Response().Writer.WriteHeader(status)

	return nil
}
