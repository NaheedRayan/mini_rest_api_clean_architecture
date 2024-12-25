package infrastructure

import (
	"github.com/NaheedRayan/mini_rest_api_shikho/interfaces"

	"github.com/gin-gonic/gin"
)

func SetupRouter(courseHandler *interfaces.CourseHandler) *gin.Engine {
	router := gin.Default()

	// Define routes


	// User routes
	router.POST("/user", courseHandler.CreateUser)
	router.GET("/users", courseHandler.GetAllUsers)
	router.GET("/user/:id", courseHandler.GetUserByID)
	router.PUT("/user", courseHandler.UpdateUser)
	router.DELETE("/user/:id", courseHandler.DeleteUser)
	
	// Course routes
	router.POST("/course", courseHandler.CreateCourse)
	router.GET("/courses", courseHandler.GetAllCourses)
	router.GET("/course/:id", courseHandler.GetCourseByID)
	router.PUT("/course", courseHandler.UpdateCourse)
	router.DELETE("/course/:id", courseHandler.DeleteCourse)

	// Enroll routes
	router.POST("/enroll", courseHandler.AddEnrollment)
	router.GET("/enrolls", courseHandler.GetAllEnrollments)
	router.GET("/enroll/:id", courseHandler.GetEnrollmentByID)
	router.GET("/enrolls/user/:id", courseHandler.GetEnrollmentsByUserID)
	router.PUT("/enroll", courseHandler.UpdateEnrollment)
	router.DELETE("/enroll/:id", courseHandler.DeleteEnrollment)

	// Lesson routes
	router.POST("/lesson", courseHandler.AddLesson)
	router.GET("/lessons/course/:id", courseHandler.GetLessonsByCourseID)
	router.GET("/lesson/:id", courseHandler.GetLessonsByID)
	router.PUT("/lesson", courseHandler.UpdateLesson)
	router.DELETE("/lesson/:id", courseHandler.DeleteLesson)

	// Progress routes
	router.POST("/progress", courseHandler.AddProgress)
	router.PUT("/progress", courseHandler.UpdateProgress)
	router.GET("/progress/:enrollmentID/:lessonID", courseHandler.GetProgressByEnrollmentAndLesson)

	// Review routes
	router.POST("/review", courseHandler.AddReview)
	router.GET("/reviews/course/:id", courseHandler.GetReviewsByCourseID)
	router.DELETE("/review/:id", courseHandler.DeleteReview)




	return router
}
