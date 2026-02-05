package main

import (
	"fmt"

	"site-admin-api/config"
	"site-admin-api/internal/handlers"
	"site-admin-api/internal/middlewares"
	"site-admin-api/internal/repositories"
	"site-admin-api/internal/routes"
	"site-admin-api/internal/services"
	"site-admin-api/pkg/database"
	"site-admin-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	logger.Info("Starting LP Ma'arif NU Admin API...")

	// Load configuration
	cfg := config.LoadConfig()
	logger.Info(fmt.Sprintf("Environment: %s", cfg.App.Env))

	// Initialize database
	db, err := database.InitDB(cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database: " + err.Error())
	}
	logger.Info("Database connected successfully")

	// Set Gin mode
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Gin router
	router := gin.New()

	// Global middlewares
	router.Use(middlewares.LoggerMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.CORSMiddleware(cfg))

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	newsRepo := repositories.NewNewsRepository(db)
	opinionRepo := repositories.NewOpinionRepository(db)
	documentRepo := repositories.NewDocumentRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	tagRepo := repositories.NewTagRepository(db)
	heroSlideRepo := repositories.NewHeroSlideRepository(db)
	orgPositionRepo := repositories.NewOrganizationPositionRepository(db)
	boardMemberRepo := repositories.NewBoardMemberRepository(db)
	pengurusRepo := repositories.NewPengurusRepository(db)
	departmentRepo := repositories.NewDepartmentRepository(db)
	editorialTeamRepo := repositories.NewEditorialTeamRepository(db)
	editorialCouncilRepo := repositories.NewEditorialCouncilRepository(db)
	pageRepo := repositories.NewPageRepository(db)
	eventFlyerRepo := repositories.NewEventFlyerRepository(db)
	mediaRepo := repositories.NewMediaRepository(db)
	contactRepo := repositories.NewContactMessageRepository(db)
	settingRepo := repositories.NewSettingRepository(db)
	activityLogRepo := repositories.NewActivityLogRepository(db)
	notificationRepo := repositories.NewNotificationRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg)
	userService := services.NewUserService(userRepo, cfg)
	newsService := services.NewNewsService(newsRepo, cfg)
	opinionService := services.NewOpinionService(opinionRepo, cfg)
	documentService := services.NewDocumentService(documentRepo, cfg)
	categoryService := services.NewCategoryService(categoryRepo, cfg)
	tagService := services.NewTagService(tagRepo, cfg)
	heroSlideService := services.NewHeroSlideService(heroSlideRepo, cfg)
	orgPositionService := services.NewOrganizationPositionService(orgPositionRepo, cfg)
	boardMemberService := services.NewBoardMemberService(boardMemberRepo, cfg)
	pengurusService := services.NewPengurusService(pengurusRepo, cfg)
	departmentService := services.NewDepartmentService(departmentRepo, cfg)
	editorialTeamService := services.NewEditorialTeamService(editorialTeamRepo, cfg)
	editorialCouncilService := services.NewEditorialCouncilService(editorialCouncilRepo, cfg)
	pageService := services.NewPageService(pageRepo, cfg)
	eventFlyerService := services.NewEventFlyerService(eventFlyerRepo, cfg)
	mediaService := services.NewMediaService(mediaRepo, cfg)
	contactService := services.NewContactMessageService(contactRepo, cfg)
	settingService := services.NewSettingService(settingRepo, cfg)
	activityLogService := services.NewActivityLogService(activityLogRepo, cfg)
	notificationService := services.NewNotificationService(notificationRepo, cfg)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, userService)
	userHandler := handlers.NewUserHandler(userService)
	newsHandler := handlers.NewNewsHandler(newsService)
	opinionHandler := handlers.NewOpinionHandler(opinionService)
	documentHandler := handlers.NewDocumentHandler(documentService)
	heroSlideHandler := handlers.NewHeroSlideHandler(heroSlideService)
	organizationHandler := handlers.NewOrganizationHandler(
		orgPositionService,
		boardMemberService,
		pengurusService,
		departmentService,
		editorialTeamService,
		editorialCouncilService,
	)
	pageHandler := handlers.NewPageHandler(pageService)
	eventFlyerHandler := handlers.NewEventFlyerHandler(eventFlyerService)
	mediaHandler := handlers.NewMediaHandler(mediaService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	tagHandler := handlers.NewTagHandler(tagService)
	contactMessageHandler := handlers.NewContactMessageHandler(contactService)
	settingHandler := handlers.NewSettingHandler(settingService)
	activityLogHandler := handlers.NewActivityLogHandler(activityLogService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	cdnHandler := handlers.NewCDNHandler(cfg)

	// Setup routes
	routes.SetupRoutes(
		router,
		cfg,
		authService,
		authHandler,
		userHandler,
		newsHandler,
		opinionHandler,
		documentHandler,
		heroSlideHandler,
		organizationHandler,
		pageHandler,
		eventFlyerHandler,
		mediaHandler,
		categoryHandler,
		tagHandler,
		contactMessageHandler,
		settingHandler,
		activityLogHandler,
		notificationHandler,
		cdnHandler,
	)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	logger.Info(fmt.Sprintf("Server starting on %s...", addr))
	logger.Info(fmt.Sprintf("API available at http://localhost:%d/api/%s", cfg.App.Port, cfg.App.APIVersion))
	logger.Info(fmt.Sprintf("Health check: http://localhost:%d/health", cfg.App.Port))

	if err := router.Run(addr); err != nil {
		logger.Fatal("Failed to start server: " + err.Error())
	}
}
