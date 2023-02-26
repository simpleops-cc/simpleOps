package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/datameta/apis"
	"go-admin/common/actions"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerDatametaRegionRouter)
}

// registerDatametaRegionRouter
func registerDatametaRegionRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.DatametaRegion{}
	// r := v1.Group("/datameta-region").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	r := v1.Group("/datameta-region")
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}
