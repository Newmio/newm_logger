package logger

import (
	newm "newm/internal"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	s ILoggerService
}

func NewLoggerHandler(s ILoggerService) *Handler {
	return &Handler{s: s}
}

func (h *Handler) InitLoggerRoutes(e *echo.Echo) *echo.Echo {

	create := e.Group("/create")
	{
		create.POST("/log", h.CreateLogRout)
		create.POST("/logs", h.CreateLogsRout)
	}

	return e
}

func (h *Handler) CreateLogsRout(c echo.Context) error {
	var log []Log

	if err := c.Bind(&log); err != nil {
		return c.JSON(400, errorRespnse(newm.Trace(err)))
	}

	if err := h.s.CreateArrayLog(log); err != nil {
		return c.JSON(500, errorRespnse(newm.Trace(err)))
	}

	return c.JSON(200, map[string]string{"status": "ok"})
}

func (h *Handler) CreateLogRout(c echo.Context) error {
	var log Log

	if err := c.Bind(&log); err != nil {
		return c.JSON(400, errorRespnse(newm.Trace(err)))
	}

	if err := h.s.CreateLog(&log); err != nil {
		return c.JSON(500, errorRespnse(newm.Trace(err)))
	}

	return c.JSON(200, map[string]string{"status": "ok"})
}
