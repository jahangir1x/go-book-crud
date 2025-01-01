package routes

import "github.com/labstack/echo/v4"

type PingRoutes struct {
	echo *echo.Echo
}

func NewPingRoutes(echo *echo.Echo) *PingRoutes {
	return &PingRoutes{
		echo: echo,
	}
}

func (pingRoutes *PingRoutes) RegisterPingRoutes(e *echo.Echo) {
	group := e.Group("/v1")
	group.GET("/server/ping", PongServer)
	group.GET("/database/ping", PongDatabase)
}

func PongServer(ctx echo.Context) error {
	return ctx.JSON(200, "pong [Server is running]")
}

func PongDatabase(ctx echo.Context) error {
	return nil
}
