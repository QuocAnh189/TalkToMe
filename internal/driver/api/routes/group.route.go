package routes

import (
	"gochat/internal/application/service"
	"gochat/internal/driver/api/handler"
	"gochat/internal/driver/api/middlewares"
	"gochat/internal/infrashstructrure/cache"
	"gochat/internal/infrashstructrure/persistence/db"
	"gochat/internal/infrashstructrure/persistence/repository"
	"gochat/pkg/storage"
	"gochat/pkg/token"
	"gochat/pkg/validation"

	"github.com/gin-gonic/gin"
)

func NewGroupRoutes(
	r *gin.RouterGroup,
	sqlDB db.IDatabase,
	validator validation.Validation,
	storage storage.IUploadService,
	cache cache.IRedis,
	token token.IMarker,
) {
	groupRepository := repository.NewGroupRepository(sqlDB)
	groupUserRepository := repository.NewGroupUserRepository(sqlDB)
	groupService := service.NewGroupService(validator, groupRepository, groupUserRepository, storage)
	groupHandler := handler.NewGroupHandler(groupService)

	authMiddleware := middlewares.NewAuthMiddleware(token, cache).TokenAuth()

	groupRoutes := r.Group("/groups").Use(authMiddleware)
	{
		groupRoutes.POST("", groupHandler.CreateGroup)
		groupRoutes.GET("", groupHandler.ListUserGroups)
		groupRoutes.GET("/:groupId", groupHandler.GetGroupDetails)
		groupRoutes.PUT("/:groupId", groupHandler.UpdateGroup)
		groupRoutes.DELETE("/:groupId", groupHandler.DeleteGroup)

		// Member management
		groupRoutes.POST("/:groupId/members", groupHandler.AddMember)
		groupRoutes.DELETE("/:groupId/members/:userId", groupHandler.RemoveMember)
		groupRoutes.GET("/:groupId/members", groupHandler.ListMembers)
	}
}
