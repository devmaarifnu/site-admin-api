package routes

import (
	"site-admin-api/config"
	"site-admin-api/internal/handlers"
	"site-admin-api/internal/middlewares"
	"site-admin-api/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(
	router *gin.Engine,
	cfg *config.Config,
	authService services.AuthService,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	newsHandler *handlers.NewsHandler,
	opinionHandler *handlers.OpinionHandler,
	documentHandler *handlers.DocumentHandler,
	heroSlideHandler *handlers.HeroSlideHandler,
	organizationHandler *handlers.OrganizationHandler,
	pageHandler *handlers.PageHandler,
	eventFlyerHandler *handlers.EventFlyerHandler,
	mediaHandler *handlers.MediaHandler,
	categoryHandler *handlers.CategoryHandler,
	tagHandler *handlers.TagHandler,
	contactMessageHandler *handlers.ContactMessageHandler,
	settingHandler *handlers.SettingHandler,
	activityLogHandler *handlers.ActivityLogHandler,
	notificationHandler *handlers.NotificationHandler,
	cdnHandler *handlers.CDNHandler,
) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "LP Ma'arif NU Admin API is running",
			"version": cfg.App.APIVersion,
		})
	})

	// API routes
	api := router.Group("/api/" + cfg.App.APIVersion)
	{
		setupAuthRoutes(api, authHandler)
		setupPublicRoutes(api, mediaHandler)
		setupProtectedRoutes(api, cfg, authService, authHandler, userHandler, newsHandler, opinionHandler,
			documentHandler, heroSlideHandler, organizationHandler, pageHandler,
			eventFlyerHandler, mediaHandler, categoryHandler, tagHandler,
			contactMessageHandler, settingHandler, activityLogHandler, notificationHandler, cdnHandler)
	}
}

// setupAuthRoutes configures public authentication routes
func setupAuthRoutes(api *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := api.Group("/admin/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.RefreshToken)
		auth.POST("/forgot-password", authHandler.ForgotPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)
	}
}

// setupPublicRoutes configures public routes (CDN client)
func setupPublicRoutes(api *gin.RouterGroup, mediaHandler *handlers.MediaHandler) {
	cdn := api.Group("/cdn")
	{
		cdn.GET("/media", mediaHandler.GetAll)
		cdn.GET("/media/:id", mediaHandler.GetByID)
	}
}

// setupProtectedRoutes configures protected routes that require authentication
func setupProtectedRoutes(
	api *gin.RouterGroup,
	cfg *config.Config,
	authService services.AuthService,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	newsHandler *handlers.NewsHandler,
	opinionHandler *handlers.OpinionHandler,
	documentHandler *handlers.DocumentHandler,
	heroSlideHandler *handlers.HeroSlideHandler,
	organizationHandler *handlers.OrganizationHandler,
	pageHandler *handlers.PageHandler,
	eventFlyerHandler *handlers.EventFlyerHandler,
	mediaHandler *handlers.MediaHandler,
	categoryHandler *handlers.CategoryHandler,
	tagHandler *handlers.TagHandler,
	contactMessageHandler *handlers.ContactMessageHandler,
	settingHandler *handlers.SettingHandler,
	activityLogHandler *handlers.ActivityLogHandler,
	notificationHandler *handlers.NotificationHandler,
	cdnHandler *handlers.CDNHandler,
) {
	admin := api.Group("/admin")
	admin.Use(middlewares.AuthMiddleware(cfg, authService))
	{
		// Authenticated user routes
		admin.GET("/auth/me", authHandler.GetCurrentUser)
		admin.POST("/auth/logout", authHandler.Logout)
		admin.PUT("/auth/change-password", authHandler.ChangePassword)

		// User Management (Super Admin only)
		setupUserRoutes(admin, userHandler)

		// News Articles
		setupNewsRoutes(admin, newsHandler)

		// Opinion Articles
		setupOpinionRoutes(admin, opinionHandler)

		// Documents
		setupDocumentRoutes(admin, documentHandler)

		// Hero Slides
		setupHeroSlideRoutes(admin, heroSlideHandler)

		// Organization
		setupOrganizationRoutes(admin, organizationHandler)

		// Pages
		setupPageRoutes(admin, pageHandler)

		// Event Flyers
		setupEventFlyerRoutes(admin, eventFlyerHandler)

		// Media Library
		setupMediaRoutes(admin, mediaHandler)

		// Categories
		setupCategoryRoutes(admin, categoryHandler)

		// Tags
		setupTagRoutes(admin, tagHandler)

		// Contact Messages
		setupContactMessageRoutes(admin, contactMessageHandler)

		// Settings
		setupSettingRoutes(admin, settingHandler)

		// Activity Logs
		setupActivityLogRoutes(admin, activityLogHandler)

		// Notifications
		setupNotificationRoutes(admin, notificationHandler)

		// CDN Upload
		setupCDNRoutes(admin, cdnHandler)
	}
}

func setupUserRoutes(admin *gin.RouterGroup, handler *handlers.UserHandler) {
	users := admin.Group("/users")
	users.Use(middlewares.PermissionMiddleware("users.view"))
	{
		users.GET("", handler.GetAll)
		users.GET("/:id", handler.GetByID)
		users.POST("", middlewares.PermissionMiddleware("users.create"), handler.Create)
		users.PUT("/:id", middlewares.PermissionMiddleware("users.update"), handler.Update)
		users.DELETE("/:id", middlewares.PermissionMiddleware("users.delete"), handler.Delete)
		users.PATCH("/:id/status", middlewares.PermissionMiddleware("users.update"), handler.UpdateStatus)
	}
}

func setupNewsRoutes(admin *gin.RouterGroup, handler *handlers.NewsHandler) {
	news := admin.Group("/news")
	news.Use(middlewares.PermissionMiddleware("news.view"))
	{
		news.GET("", handler.GetAll)
		news.GET("/:id", handler.GetByID)
		news.POST("", middlewares.PermissionMiddleware("news.create"), handler.Create)
		news.PUT("/:id", middlewares.PermissionMiddleware("news.update"), handler.Update)
		news.DELETE("/:id", middlewares.PermissionMiddleware("news.delete"), handler.Delete)
		news.PATCH("/:id/publish", middlewares.PermissionMiddleware("news.update"), handler.Publish)
		news.PATCH("/:id/archive", middlewares.PermissionMiddleware("news.update"), handler.Archive)
		news.PATCH("/:id/featured", middlewares.PermissionMiddleware("news.update"), handler.ToggleFeatured)
	}
}

func setupOpinionRoutes(admin *gin.RouterGroup, handler *handlers.OpinionHandler) {
	opinions := admin.Group("/opinions")
	opinions.Use(middlewares.PermissionMiddleware("opinions.view"))
	{
		opinions.GET("", handler.GetAll)
		opinions.GET("/:id", handler.GetByID)
		opinions.POST("", middlewares.PermissionMiddleware("opinions.create"), handler.Create)
		opinions.PUT("/:id", middlewares.PermissionMiddleware("opinions.update"), handler.Update)
		opinions.DELETE("/:id", middlewares.PermissionMiddleware("opinions.delete"), handler.Delete)
		opinions.PATCH("/:id/publish", middlewares.PermissionMiddleware("opinions.update"), handler.Publish)
	}
}

func setupDocumentRoutes(admin *gin.RouterGroup, handler *handlers.DocumentHandler) {
	documents := admin.Group("/documents")
	documents.Use(middlewares.PermissionMiddleware("documents.view"))
	{
		documents.GET("", handler.GetAll)
		documents.GET("/:id", handler.GetByID)
		documents.POST("", middlewares.PermissionMiddleware("documents.create"), handler.Create)
		documents.PUT("/:id/file", middlewares.PermissionMiddleware("documents.update"), handler.ReplaceFile)
		documents.DELETE("/:id", middlewares.PermissionMiddleware("documents.delete"), handler.Delete)
		documents.GET("/:id/stats", handler.GetStats)
	}
}

func setupHeroSlideRoutes(admin *gin.RouterGroup, handler *handlers.HeroSlideHandler) {
	heroSlides := admin.Group("/hero-slides")
	heroSlides.Use(middlewares.PermissionMiddleware("hero_slides.view"))
	{
		heroSlides.GET("", handler.GetAll)
		heroSlides.GET("/:id", handler.GetByID)
		heroSlides.POST("", middlewares.PermissionMiddleware("hero_slides.create"), handler.Create)
		heroSlides.PUT("/:id", middlewares.PermissionMiddleware("hero_slides.update"), handler.Update)
		heroSlides.DELETE("/:id", middlewares.PermissionMiddleware("hero_slides.delete"), handler.Delete)
		heroSlides.PUT("/reorder", middlewares.PermissionMiddleware("hero_slides.update"), handler.Reorder)
	}
}

func setupOrganizationRoutes(admin *gin.RouterGroup, handler *handlers.OrganizationHandler) {
	org := admin.Group("/organization")
	org.Use(middlewares.PermissionMiddleware("organization.view"))
	{
		org.GET("/positions", handler.GetPositions)
		org.GET("/board-members", handler.GetBoardMembers)
		org.POST("/board-members", middlewares.PermissionMiddleware("organization.create"), handler.CreateBoardMember)
		org.PUT("/board-members/:id", middlewares.PermissionMiddleware("organization.update"), handler.UpdateBoardMember)
		org.DELETE("/board-members/:id", middlewares.PermissionMiddleware("organization.delete"), handler.DeleteBoardMember)
		org.GET("/pengurus", handler.GetPengurus)
		org.GET("/pengurus/:id", handler.GetPengurusByID)
		org.POST("/pengurus", middlewares.PermissionMiddleware("organization.create"), handler.CreatePengurus)
		org.PUT("/pengurus/:id", middlewares.PermissionMiddleware("organization.update"), handler.UpdatePengurus)
		org.DELETE("/pengurus/:id", middlewares.PermissionMiddleware("organization.delete"), handler.DeletePengurus)
		org.PUT("/pengurus/reorder", middlewares.PermissionMiddleware("organization.update"), handler.ReorderPengurus)
		org.GET("/departments", handler.GetDepartments)
		org.PUT("/departments/:id", middlewares.PermissionMiddleware("organization.update"), handler.UpdateDepartment)
		org.GET("/editorial-team", handler.GetEditorialTeam)
		org.PUT("/editorial-team/:id", middlewares.PermissionMiddleware("organization.update"), handler.UpdateEditorialTeam)
		org.GET("/editorial-council", handler.GetEditorialCouncil)
		org.PUT("/editorial-council/:id", middlewares.PermissionMiddleware("organization.update"), handler.UpdateEditorialCouncil)
	}
}

func setupPageRoutes(admin *gin.RouterGroup, handler *handlers.PageHandler) {
	pages := admin.Group("/pages")
	pages.Use(middlewares.PermissionMiddleware("pages.view"))
	{
		pages.GET("", handler.GetAll)
		pages.GET("/:slug", handler.GetBySlug)
		pages.PUT("/:slug", middlewares.PermissionMiddleware("pages.update"), handler.Update)
	}
}

func setupEventFlyerRoutes(admin *gin.RouterGroup, handler *handlers.EventFlyerHandler) {
	eventFlyers := admin.Group("/event-flyers")
	eventFlyers.Use(middlewares.PermissionMiddleware("events.view"))
	{
		eventFlyers.GET("", handler.GetAll)
		eventFlyers.GET("/:id", handler.GetByID)
		eventFlyers.POST("", middlewares.PermissionMiddleware("events.create"), handler.Create)
		eventFlyers.PUT("/:id", middlewares.PermissionMiddleware("events.update"), handler.Update)
		eventFlyers.DELETE("/:id", middlewares.PermissionMiddleware("events.delete"), handler.Delete)
	}
}

func setupMediaRoutes(admin *gin.RouterGroup, handler *handlers.MediaHandler) {
	media := admin.Group("/media")
	media.Use(middlewares.PermissionMiddleware("media.view"))
	{
		media.GET("", handler.GetAll)
		media.POST("/upload", middlewares.PermissionMiddleware("media.upload"), handler.Upload)
		media.DELETE("/:id", middlewares.PermissionMiddleware("media.delete"), handler.Delete)
	}
}

func setupCategoryRoutes(admin *gin.RouterGroup, handler *handlers.CategoryHandler) {
	categories := admin.Group("/categories")
	categories.Use(middlewares.PermissionMiddleware("categories.view"))
	{
		categories.GET("", handler.GetAll)
		categories.POST("", middlewares.PermissionMiddleware("categories.create"), handler.Create)
		categories.PUT("/:id", middlewares.PermissionMiddleware("categories.update"), handler.Update)
		categories.DELETE("/:id", middlewares.PermissionMiddleware("categories.delete"), handler.Delete)
	}
}

func setupTagRoutes(admin *gin.RouterGroup, handler *handlers.TagHandler) {
	tags := admin.Group("/tags")
	tags.Use(middlewares.PermissionMiddleware("tags.view"))
	{
		tags.GET("", handler.GetAll)
		tags.POST("", middlewares.PermissionMiddleware("tags.create"), handler.Create)
		tags.PUT("/:id", middlewares.PermissionMiddleware("tags.update"), handler.Update)
		tags.DELETE("/:id", middlewares.PermissionMiddleware("tags.delete"), handler.Delete)
	}
}

func setupContactMessageRoutes(admin *gin.RouterGroup, handler *handlers.ContactMessageHandler) {
	contactMessages := admin.Group("/contact-messages")
	contactMessages.Use(middlewares.PermissionMiddleware("contact_messages.view"))
	{
		contactMessages.GET("", handler.GetAll)
		contactMessages.GET("/:id", handler.GetByID)
		contactMessages.PATCH("/:id/status", middlewares.PermissionMiddleware("contact_messages.update"), handler.UpdateStatus)
		contactMessages.DELETE("/:id", middlewares.PermissionMiddleware("contact_messages.delete"), handler.Delete)
	}
}

func setupSettingRoutes(admin *gin.RouterGroup, handler *handlers.SettingHandler) {
	settings := admin.Group("/settings")
	settings.Use(middlewares.PermissionMiddleware("settings.view"))
	{
		settings.GET("", handler.GetAll)
		settings.PUT("", middlewares.PermissionMiddleware("settings.update"), handler.Update)
	}

}

func setupActivityLogRoutes(admin *gin.RouterGroup, handler *handlers.ActivityLogHandler) {
	activityLogs := admin.Group("/activity-logs")
	activityLogs.Use(middlewares.PermissionMiddleware("activity_logs.view"))
	{
		activityLogs.GET("", handler.GetAll)
	}
}

func setupNotificationRoutes(admin *gin.RouterGroup, handler *handlers.NotificationHandler) {
	notifications := admin.Group("/notifications")
	{
		notifications.GET("", handler.GetAll)
		notifications.PATCH("/:id/read", handler.MarkAsRead)
		notifications.PATCH("/read-all", handler.MarkAllAsRead)
		notifications.DELETE("/:id", handler.Delete)
	}
}

func setupCDNRoutes(admin *gin.RouterGroup, handler *handlers.CDNHandler) {
	cdn := admin.Group("/cdn")
	cdn.Use(middlewares.PermissionMiddleware("media.upload"))
	{
		cdn.POST("/upload", handler.UploadToCDN)
		cdn.DELETE("/:tag/:filename", handler.DeleteFromCDN)
		cdn.GET("/url", handler.GetCDNFileURL)
	}
}
