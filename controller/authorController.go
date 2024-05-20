package controller

import (
	"net/http"
	"strconv"

	"go-arch/middleware"
	"go-arch/model/dto"
	"go-arch/service"

	"github.com/gin-gonic/gin"
)

type authorController struct {
	authorService service.AuthorServiceI
	router        *gin.RouterGroup
	authMiddle    middleware.AuthMiddlewareI
}

func (a *authorController) listHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	size, err2 := strconv.Atoi(c.Query("size"))
	if err != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err2 != nil {
		dto.SendErrorResponse(c, http.StatusBadRequest, err2.Error())
		return
	}

	listData, paging, err := a.authorService.FindAll(page, size)
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

func (a *authorController) getByIdHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := a.authorService.FindById(id)

	if err != nil {
		dto.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	dto.SendSingleResponse(c, data, http.StatusOK, "OK")
}

func (a *authorController) Routing() {
	a.router.GET("/authors", a.authMiddle.VerifyTokenAndRole("admin", "user"), a.listHandler)
	a.router.GET("/authors/:id", a.authMiddle.VerifyTokenAndRole("admin", "user"), a.getByIdHandler)
}

func NewAuthorController(authorService service.AuthorServiceI, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddlewareI) *authorController {
	return &authorController{authorService: authorService, router: rg, authMiddle: authMiddleware}
}
