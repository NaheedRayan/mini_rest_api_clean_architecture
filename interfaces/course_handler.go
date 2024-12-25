package interfaces

import (
	"strconv"

	"github.com/NaheedRayan/mini_rest_api_shikho/entities"
	"github.com/NaheedRayan/mini_rest_api_shikho/usecases"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CourseHandler is a struct that defines the handler for the course entity
type CourseHandler struct {
	UseCase *usecases.CourseUseCase
}
//constructor
func NewCourseHandler(uc *usecases.CourseUseCase) *CourseHandler {
	return &CourseHandler{UseCase: uc}
}

//----------------------------------------------------------------user----------------------------------------------------------------

// CreateUser is a handler function to create a new user
func (h *CourseHandler) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user , err := h.UseCase.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *CourseHandler) GetAllUsers(c *gin.Context) {
	users, err := h.UseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *CourseHandler) GetUserByID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := h.UseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *CourseHandler) UpdateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user , err := h.UseCase.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}


func (h *CourseHandler) DeleteUser(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = h.UseCase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

//----------------------------------------------------------------course----------------------------------------------------------------
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var course entities.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course , err := h.UseCase.AddCourse(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.UseCase.ListCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	course, err := h.UseCase.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) UpdateCourse(c *gin.Context) {
	var course entities.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course , err := h.UseCase.UpdateCourse(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func (h *CourseHandler) DeleteCourse(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	err = h.UseCase.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}


//----------------------------------------------------------------enrollment----------------------------------------------------------------

func (h *CourseHandler) AddEnrollment(c *gin.Context) {
	var enrollment entities.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrollment , err := h.UseCase.AddEnrollment(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Enrollment created successfully"})
}

func (h *CourseHandler) GetAllEnrollments(c *gin.Context) {
	enrollments, err := h.UseCase.GetAllEnrollments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

func (h *CourseHandler) GetEnrollmentByID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}
	enrollment, err := h.UseCase.GetEnrollmentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

func (h *CourseHandler) GetEnrollmentsByUserID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	enrollments, err := h.UseCase.GetEnrollmentsByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

func (h *CourseHandler) UpdateEnrollment(c *gin.Context) {
	var enrollment entities.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrollment , err := h.UseCase.UpdateEnrollment(enrollment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment updated successfully"})
}

func (h *CourseHandler) DeleteEnrollment(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}
	err = h.UseCase.DeleteEnrollment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}

//----------------------------------------------------------------lesson----------------------------------------------------------------

func (h *CourseHandler) AddLesson(c *gin.Context) {
	var lesson entities.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson , err := h.UseCase.AddLesson(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Lesson created successfully"})
}

func (h *CourseHandler) GetLessonsByCourseID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	lessons, err := h.UseCase.GetLessonsByCourseID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

func (h *CourseHandler) GetLessonsByID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}
	lessons, err := h.UseCase.GetLessonsByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

func (h *CourseHandler) UpdateLesson(c *gin.Context) {
	var lesson entities.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson , err := h.UseCase.UpdateLesson(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson updated successfully"})
}

func (h *CourseHandler) DeleteLesson(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}
	err = h.UseCase.DeleteLesson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}

//----------------------------------------------------------------progress----------------------------------------------------------------

func (h *CourseHandler) AddProgress(c *gin.Context) {
	var progress entities.Progress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	progress , err := h.UseCase.AddProgress(progress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Progress created successfully"})
}

func (h *CourseHandler) UpdateProgress(c *gin.Context) {
	var progress entities.Progress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	progress , err := h.UseCase.UpdateProgress(progress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Progress updated successfully"})
}

func (h *CourseHandler) GetProgressByEnrollmentAndLesson(c *gin.Context) {
	str_enrollment_id := c.Param("enrollment_id")
	enrollment_id, err := strconv.Atoi(str_enrollment_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	str_lesson_id := c.Param("lesson_id")
	lesson_id, err := strconv.Atoi(str_lesson_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	progress, err := h.UseCase.GetProgressByEnrollmentAndLesson(enrollment_id, lesson_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

//----------------------------------------------------------------review----------------------------------------------------------------

func (h *CourseHandler) AddReview(c *gin.Context) {
	var review entities.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review , err := h.UseCase.AddReview(review)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review created successfully"})
}

func (h *CourseHandler) GetReviewsByCourseID(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	reviews, err := h.UseCase.GetReviewsByCourseID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *CourseHandler) DeleteReview(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.Atoi(str_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	err = h.UseCase.DeleteReview(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}