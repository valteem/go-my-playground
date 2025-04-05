package api

import (
	ac "webapi/gin-router-groups/pkg/accesscontrol"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")

	apiAdminUsers := v1.Group("/api/admin/users")

	apiAdminUsers.POST("/", authorize(ac.EvalPermissions(ac.ActionUserCreate)), AdminCreateUser)

}

func authorize(gin.HandlerFunc) gin.HandlerFunc {
	// stub
	return func(c *gin.Context) {
		//stub
	}
}
