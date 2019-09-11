package handler

import (
	"context"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"goNam/contract"
	"goNam/models"
    "fmt"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	AnUC contract.UseCase
}

func (a *UserHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AnUC.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, art)
}
func (a *UserHandler) Add(c echo.Context) error {
	var user models.UserModel
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()

	fmt.Println(c.Request().Context())
	fmt.Println(context.Background())
	if ctx == nil {
		ctx = context.Background()
	}

	err = a.AnUC.Add(ctx,&user)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}
func (a *UserHandler) DeleteById(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	  err = a.AnUC.DeleteById(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (a *UserHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listAr, err := a.AnUC.Fetch(ctx)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	//c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}
func NewUserHandler(e *echo.Echo, us contract.UseCase) {
	handler := &UserHandler{
		AnUC: us,
	}
	e.GET("/users/:id", handler.GetByID)
	e.GET("/users", handler.Fetch)
	e.DELETE("/users/:id", handler.DeleteById)
	//e.PUT("/users/:id", handler.UpdateById)
	e.POST("/users",handler.Add)

}


func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
