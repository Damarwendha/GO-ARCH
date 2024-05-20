package controller

import (
	"go-arch/middleware"
	"go-arch/model"
	"go-arch/model/dto"
	"go-arch/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService service.TaskServiceI
	router      *gin.RouterGroup
	authMiddle  middleware.AuthMiddlewareI
}

func (t *taskController) listHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	size, err2 := strconv.Atoi(c.Query("size"))
	if err != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err2 != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listData, paging, err := t.taskService.FindAll(page, size)
	if err != nil {
		dto.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// convert to interfaces
	listDataI := make([]interface{}, len(listData))
	for i, v := range listData {
		listDataI[i] = v
	}

	dto.SendManyResponse(c, listDataI, paging, http.StatusOK, "OK")
}

func (t *taskController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := t.taskService.FindById(id)

	if err != nil {
		dto.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	dto.SendSingleResponse(c, data, http.StatusOK, "OK")
}

func (t *taskController) updateByIdHandler(c *gin.Context) {
	var payload model.Task

	if err := c.ShouldBindJSON(&payload); err != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Param("id")

	if err := t.taskService.UpdateById(id, payload); err != nil {
		dto.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	dto.SendSingleResponse(c, nil, http.StatusOK, "OK")
}

func (t *taskController) createHandler(c *gin.Context) {
	var payload model.Task

	if err := c.ShouldBindJSON(&payload); err != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := t.taskService.Create(payload); err != nil {
		dto.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	dto.SendSingleResponse(c, nil, http.StatusCreated, "Created")
}

func (t *taskController) Routing() {
	t.router.GET("/tasks", t.authMiddle.VerifyTokenAndRole("admin"), t.listHandler)
	t.router.GET("/tasks/:id", t.authMiddle.VerifyTokenAndRole("admin"), t.getByIdHandler)
	t.router.PUT("/tasks/:id", t.authMiddle.VerifyTokenAndRole("admin"), t.updateByIdHandler)
	t.router.POST("/tasks", t.authMiddle.VerifyTokenAndRole("admin"), t.createHandler)
}

func NewTaskController(taskService service.TaskServiceI, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddlewareI) *taskController {
	return &taskController{taskService: taskService, router: rg, authMiddle: authMiddleware}
}
