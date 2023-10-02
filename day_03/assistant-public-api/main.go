package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/clients"
	v1 "github.com/ideal-forward/assistant-public-api/controllers/v1"
	"github.com/ideal-forward/assistant-public-api/controllers/v1/audit"
	"github.com/ideal-forward/assistant-public-api/controllers/v1/home"
	"github.com/ideal-forward/assistant-public-api/controllers/v1/private"
	"github.com/ideal-forward/assistant-public-api/controllers/v1/project"
	"github.com/ideal-forward/assistant-public-api/controllers/v1/public"
	"github.com/ideal-forward/assistant-public-api/controllers/v1/report"
	"github.com/ideal-forward/assistant-public-api/controllers/v1/timesheets"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/log"
	"github.com/ideal-forward/assistant-public-api/services"
	"github.com/spf13/viper"
)

func StartHTTPServer() {

	log.Setup()

	router := gin.New()
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/*")

	// healthcheck
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "success", "message": "How to Upload Single and Multiple Files in Golang"})
	})

	// upload file
	router.POST("/upload/single", uploadResizeSingleFile)
	router.POST("/upload/multiple", uploadResizeMultipleFile)
	// router.StaticFS("/images", http.Dir("public"))
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.StaticFS(services.Config.FileStorage.URLPath, http.Dir(services.Config.FileStorage.Folder))
	h := v1.NewHandler()
	publicTaskHandler := public.NewHandler()
	privateTaskHandler := private.NewHandler()
	auditTaskHandler := audit.NewHandler()
	reportHandler := report.NewHandler()
	projectHandler := project.NewHandler()
	timesheetsHandler := timesheets.NewHandler()
	homeHandler := home.NewHandler()
	router.POST("/auth", h.Authenticate)
	router.POST("/api/v1/auth", h.Authenticate)
	// apis
	apiV1 := router.Group("/api/v1", middlewares.JWT())
	{
		//me
		me := apiV1.Group("/me")
		{
			me.GET("", h.ReadSelfProfile)
		}

		// user
		user := apiV1.Group("/users")
		{
			user.GET("", h.ListUser)
			user.GET("/:id", h.ReadUser)
			user.POST("", h.CreateUser)
			user.PUT("/:id", h.UpdateUser)
			user.DELETE("/:id", h.DeleteUser)
		}

		// executor
		executor := apiV1.Group("/executors")
		{
			executor.GET("", h.ListExecutor)
			executor.GET("/:id", h.ReadExecutor)
			executor.POST("", h.CreateExecutor)
			executor.PUT("/:id", h.UpdateExecutor)
			executor.DELETE("/:id", h.DeleteExecutor)
		}

		// region
		region := apiV1.Group("/regions")
		{
			region.GET("", h.ListRegion)
			region.GET("/:id", h.ReadRegion)
			region.POST("", h.CreateRegion)
			region.PUT("/:id", h.UpdateRegion)
			region.DELETE("/:id", h.DeleteRegion)
		}

		// notifications
		notification := apiV1.Group("/notifications")
		{
			notification.GET("", h.ListNotification)
			notification.GET("/:id", h.ReadNotification)
		}

		// project categories
		projectCategory := apiV1.Group("/project-categories")
		{
			projectCategory.GET("", projectHandler.ListProjectCategory)
			projectCategory.GET("/:id", projectHandler.ReadProjectCategory)
			projectCategory.POST("", projectHandler.CreateProjectCategory)
			projectCategory.PUT("/:id", projectHandler.UpdateProjectCategory)
			projectCategory.DELETE("/:id", projectHandler.DeleteProjectCategory)
		}

		// project
		project := apiV1.Group("/projects")
		{
			project.GET("", projectHandler.ListProject)
			project.GET("/:id", projectHandler.ReadProject)
			project.POST("", projectHandler.CreateProject)
			project.PUT("/:id", projectHandler.UpdateProject)
			project.DELETE("/:id", projectHandler.DeleteProject)

			project.GET("/:id/members", projectHandler.ListProjectMember)
			project.POST("/:id/members", projectHandler.AddProjectMember)
			project.DELETE("/:id/members/:user_id", projectHandler.RemoveProjectMember)

			project.GET("/:id/executors", projectHandler.ListProjectExecutor)
			project.POST("/:id/executors", projectHandler.AddProjectExecutor)
			project.DELETE("/:id/executors/:executor_id", projectHandler.RemoveProjectExecutor)

			project.GET("/:id/phases", projectHandler.ListProjectPhase)
			project.GET("/:id/phases/:phase_id", projectHandler.ReadProjectPhase)
			project.POST("/:id/phases", projectHandler.CreateProjectPhase)
			project.PUT("/:id/phases/:phase_id", projectHandler.UpdateProjectPhase)
			project.DELETE("/:id/phases/:phase_id", projectHandler.RemoveProjectPhase)

			project.GET("/:id/areas", projectHandler.ListProjectArea)
			project.GET("/:id/areas/:area_id", projectHandler.ReadProjectArea)
			project.POST("/:id/areas", projectHandler.CreateProjectArea)
			project.PUT("/:id/areas/:area_id", projectHandler.UpdateProjectArea)
			project.DELETE("/:id/areas/:area_id", projectHandler.RemoveProjectArea)
		}

		// task
		task := apiV1.Group("/tasks")
		{
			task.GET("", publicTaskHandler.ListTask)
			task.GET("/:id", publicTaskHandler.ReadTask)
			task.POST("", publicTaskHandler.CreateTask)
			task.PUT("/:id", publicTaskHandler.UpdateTask)
			task.PATCH("/:id", publicTaskHandler.UpdateTaskStatus)
			task.DELETE("/:id", publicTaskHandler.DeleteTask)

			task.GET("/:id/comments", publicTaskHandler.ListTaskComment)
			task.POST("/:id/comments", publicTaskHandler.CreateTaskComment)
			task.DELETE("/:id/comments/:comment_id", publicTaskHandler.DeleteTaskComment)

			task.GET("/:id/images", publicTaskHandler.ListTaskImage)
			task.POST("/:id/images", publicTaskHandler.CreateTaskImage)
			task.DELETE("/:id/images/:image_id", publicTaskHandler.DeleteTaskImage)

			task.POST("/:id/remind", publicTaskHandler.Remind)
		}

		// timesheets
		timesheets := apiV1.Group("/timesheets")
		{
			timesheets.GET("", timesheetsHandler.ListTimesheets)
			timesheets.GET("/:id", timesheetsHandler.ReadTimesheets)
			timesheets.POST("", timesheetsHandler.CreateTimesheets)
			timesheets.PUT("/:id", timesheetsHandler.UpdateTimesheets)
			timesheets.DELETE("/:id", timesheetsHandler.DeleteTimesheets)

			timesheets.GET("/:id/executors", timesheetsHandler.ListExecutor)
			timesheets.GET("/:id/executors/:executor_id", timesheetsHandler.ReadExecutor)
			timesheets.POST("/:id/executors", timesheetsHandler.CreateExecutor)
			timesheets.PUT("/:id/executors/:executor_id", timesheetsHandler.UpdateExecutor)
			timesheets.DELETE("/:id/executors/:executor_id", timesheetsHandler.DeleteExecutor)

			timesheets.GET("/:id/comments", timesheetsHandler.ListComment)
		}

		importantTasks := apiV1.Group("/important-tasks")
		{
			importantTasks.GET("", publicTaskHandler.ListTask)
		}

		taskStatus := apiV1.Group("/task-status")
		{
			taskStatus.PUT("/:id", publicTaskHandler.UpdateTaskStatus)
		}

		privateTask := apiV1.Group("/private-tasks")
		{
			privateTask.GET("", privateTaskHandler.ListTask)
			privateTask.GET("/:id", privateTaskHandler.ReadTask)
			privateTask.POST("", privateTaskHandler.CreateTask)
			privateTask.PUT("/:id", privateTaskHandler.UpdateTask)
			privateTask.PATCH("/:id", privateTaskHandler.UpdateTaskStatus)
			privateTask.DELETE("/:id", privateTaskHandler.DeleteTask)

			privateTask.GET("/:id/comments", privateTaskHandler.ListTaskComment)
			privateTask.POST("/:id/comments", privateTaskHandler.CreateTaskComment)
			privateTask.DELETE("/:id/comments/:comment_id", privateTaskHandler.DeleteTaskComment)

			privateTask.GET("/:id/images", privateTaskHandler.ListTaskImage)
			privateTask.POST("/:id/images", privateTaskHandler.CreateTaskImage)
			privateTask.DELETE("/:id/images/:image_id", privateTaskHandler.DeleteTaskImage)
		}

		auditTask := apiV1.Group("/audit-tasks")
		{
			auditTask.GET("", auditTaskHandler.ListTask)
			auditTask.GET("/:id", auditTaskHandler.ReadTask)
			auditTask.POST("", auditTaskHandler.CreateTask)
			auditTask.PUT("/:id", auditTaskHandler.UpdateTask)
			auditTask.PATCH("/:id", auditTaskHandler.UpdateTaskStatus)
			auditTask.DELETE("/:id", auditTaskHandler.DeleteTask)

			auditTask.GET("/:id/comments", auditTaskHandler.ListTaskComment)
			auditTask.POST("/:id/comments", auditTaskHandler.CreateTaskComment)
			auditTask.DELETE("/:id/comments/:comment_id", auditTaskHandler.DeleteTaskComment)

			auditTask.GET("/:id/images", auditTaskHandler.ListTaskImage)
			auditTask.POST("/:id/images", auditTaskHandler.CreateTaskImage)
			auditTask.DELETE("/:id/images/:image_id", auditTaskHandler.DeleteTaskImage)
		}

		reports := apiV1.Group("/reports")
		{
			reports.POST("/project", reportHandler.CompareProject)
			reports.GET("/project", reportHandler.CompareProject)
			reports.GET("/sefl-constructor", reportHandler.CompareSeflConstructor)
		}

		home := apiV1.Group("/home")
		{
			home.GET("/assigning-tasks", homeHandler.ListAssigningTask)
			home.GET("/assigned-tasks", homeHandler.ListAssignedTask)
			home.GET("/sefl-tasks", homeHandler.ListSelfTask)
		}
	}

	host := fmt.Sprintf("localhost:%d", viper.GetInt("service.port"))
	router.Run(host)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func main() {
	// excel.CreateExcel()

	services.InitViper()

	clients.MySQLClient, _ = clients.NewMySQLClient()
	clients.AutoMigrate()
	SetupFolders()

	services.InitGlobalServices()

	StartHTTPServer()
}
