package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go_clean/domain"
	"net/http"
	"strconv"

	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TodolistHandler struct {
	TdUseCase domain.TodoListUsecase
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewTodolistHandler(c *gin.Engine, us domain.TodoListUsecase) {
	handler := &TodolistHandler{
		TdUseCase: us,
	}
	c.GET("/todo/:id", handler.GetByID)
	c.GET("/todo", handler.Get)
	c.POST("/todo", handler.Insert)
	c.PUT("/todo/:id", handler.Update)
	c.DELETE("/todo/:id", handler.Delete)
}

func (t *TodolistHandler) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	todo := map[string]interface{}{
		"id": idP,
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	err = t.TdUseCase.Delete(c, idP)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
	return
}
func (t *TodolistHandler) Update(c *gin.Context) {
	var todo domain.TodoList
	err := c.ShouldBind(&todo)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	todo.ID = int64(idP)
	var ok bool
	if ok, err = isRequestValid(&todo); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = t.TdUseCase.Update(c, &todo)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
	return

}

func (t *TodolistHandler) Insert(c *gin.Context) {
	var todo domain.TodoList
	err := c.ShouldBind(&todo)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var ok bool
	if ok, err = isRequestValid(&todo); !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = t.TdUseCase.Insert(c, &todo)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": todo,
	})
	return
}
func isRequestValid(m *domain.TodoList) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil

}
func (t *TodolistHandler) Get(c *gin.Context) {
	td, err := t.TdUseCase.Get(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": td,
	})
	return
}
func (t *TodolistHandler) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{
			"message": err.Error(),
		})
		return
	}
	id := int(idP)

	td, errMsg := t.TdUseCase.GetByID(c, id)
	if errMsg != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errMsg.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": td,
	})
	return
}
