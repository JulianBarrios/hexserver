package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mooc "github.com/julianbarrios/hexserver/infrastructure/data"
)

type CreateCourseRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateCourse(r mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateCourseRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := mooc.NewCourse(req.ID, req.Name, req.Duration)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := r.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
