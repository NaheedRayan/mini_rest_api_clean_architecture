package usecases

import (
	"fmt"

	"github.com/NaheedRayan/mini_rest_api_shikho/entities"
)

type CourseRepository interface {

	// User
	CreateUser(user entities.User) (entities.User, error) 
	GetUserByID(id int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	DeleteUser(id int) error

	// Course
	AddCourse(course entities.Course) (entities.Course, error) 
	GetAllCourses() ([]entities.Course, error)
	GetCourseByID(id int) (entities.Course, error)
	UpdateCourse(course entities.Course) (entities.Course, error)
	DeleteCourse(id int) error

	// Enroll
	AddEnrollment(enrollment entities.Enrollment) (entities.Enrollment, error)
	GetAllEnrollments() ([]entities.Enrollment, error)
	GetEnrollmentByID(id int) (entities.Enrollment, error)
	GetEnrollmentsByUserID(userID int) ([]entities.Enrollment, error)
	UpdateEnrollment(enrollment entities.Enrollment) (entities.Enrollment, error)
	DeleteEnrollment(id int) error

	// Lesson
	AddLesson(lesson entities.Lesson) (entities.Lesson, error)
	GetLessonsByCourseID(courseID int) ([]entities.Lesson, error)
	GetLessonsByID(lessonID int) ([]entities.Lesson, error)
	UpdateLesson(lesson entities.Lesson) (entities.Lesson, error)
	DeleteLesson(id int) error

	// Progress
	AddProgress(progress entities.Progress) (entities.Progress, error)
	UpdateProgress(progress entities.Progress) (entities.Progress, error)
	GetProgressByEnrollmentAndLesson(enrollmentID, lessonID int) (entities.Progress, error)

	// Review
	AddReview(review entities.Review) (entities.Review, error)
	GetReviewsByCourseID(courseID int) ([]entities.Review, error)
	DeleteReview(id int) error
}


type CourseUseCase struct {
	Repo CourseRepository
}
//constructor
func NewCourseUseCase(repo CourseRepository) *CourseUseCase {
	return &CourseUseCase{Repo: repo}
}

//----------------------------------------------------------------user----------------------------------------------------------------
func (uc *CourseUseCase) CreateUser(user entities.User) (entities.User, error) {
	return uc.Repo.CreateUser(user)
}

func (uc *CourseUseCase) GetUserByID(id int) (entities.User, error) {
	return uc.Repo.GetUserByID(id)
}

func (uc *CourseUseCase) GetAllUsers() ([]entities.User, error) {
	return uc.Repo.GetAllUsers()
}

func (uc *CourseUseCase) UpdateUser(user entities.User) (entities.User, error) {
	return uc.Repo.UpdateUser(user)
}
func (uc *CourseUseCase) DeleteUser(id int) error {
	return uc.Repo.DeleteUser(id)
}


//----------------------------------------------------------------course----------------------------------------------------------------
func (uc *CourseUseCase) AddCourse(course entities.Course) (entities.Course ,error) {
	// Business rule: validate course price
	if course.Price <= 0 {
		return entities.Course{}, fmt.Errorf("price must be greater than zero")
	}
	return uc.Repo.AddCourse(course)
}

func (uc *CourseUseCase) ListCourses() ([]entities.Course, error) {
	return uc.Repo.GetAllCourses()
}

func (uc *CourseUseCase) GetCourseByID(id int) (entities.Course, error) {
	return uc.Repo.GetCourseByID(id)
}

func (uc *CourseUseCase) UpdateCourse(course entities.Course) (entities.Course, error) {
	return uc.Repo.UpdateCourse(course)
}

func (uc *CourseUseCase) DeleteCourse(id int) error {
	return uc.Repo.DeleteCourse(id)
}

//----------------------------------------------------------------enrollment----------------------------------------------------------------

func (uc *CourseUseCase) AddEnrollment(enrollment entities.Enrollment) (entities.Enrollment, error) {
	return uc.Repo.AddEnrollment(enrollment)
}

func (uc *CourseUseCase) GetAllEnrollments() ([]entities.Enrollment, error) {
	return uc.Repo.GetAllEnrollments()
}

func (uc *CourseUseCase) GetEnrollmentByID(id int) (entities.Enrollment, error) {	
	return uc.Repo.GetEnrollmentByID(id)
}

func (uc *CourseUseCase) GetEnrollmentsByUserID(userID int) ([]entities.Enrollment, error) {
	return uc.Repo.GetEnrollmentsByUserID(userID)
}

func (uc *CourseUseCase) UpdateEnrollment(enrollment entities.Enrollment) (entities.Enrollment, error) {
	return uc.Repo.UpdateEnrollment(enrollment)
}

func (uc *CourseUseCase) DeleteEnrollment(id int) error {
	return uc.Repo.DeleteEnrollment(id)
}

//----------------------------------------------------------------lesson----------------------------------------------------------------

func (uc *CourseUseCase) AddLesson(lesson entities.Lesson) (entities.Lesson, error) {
	return uc.Repo.AddLesson(lesson)
}

func (uc *CourseUseCase) GetLessonsByCourseID(courseID int) ([]entities.Lesson, error) {
	return uc.Repo.GetLessonsByCourseID(courseID)
}	

func (uc *CourseUseCase) GetLessonsByID(lessonID int) ([]entities.Lesson, error) {
	return uc.Repo.GetLessonsByID(lessonID)
}

func (uc *CourseUseCase) UpdateLesson(lesson entities.Lesson) (entities.Lesson, error) {
	return uc.Repo.UpdateLesson(lesson)
}

func (uc *CourseUseCase) DeleteLesson(id int) error {
	return uc.Repo.DeleteLesson(id)
}

//----------------------------------------------------------------progress----------------------------------------------------------------

func (uc *CourseUseCase) AddProgress(progress entities.Progress) (entities.Progress, error) {
	return uc.Repo.AddProgress(progress)
}

func (uc *CourseUseCase) UpdateProgress(progress entities.Progress) (entities.Progress, error) {
	return uc.Repo.UpdateProgress(progress)
}

func (uc *CourseUseCase) GetProgressByEnrollmentAndLesson(enrollmentID, lessonID int) (entities.Progress, error) {
	return uc.Repo.GetProgressByEnrollmentAndLesson(enrollmentID, lessonID)
}

//----------------------------------------------------------------review----------------------------------------------------------------

func (uc *CourseUseCase) AddReview(review entities.Review) (entities.Review, error) {
	return uc.Repo.AddReview(review)
}

func (uc *CourseUseCase) GetReviewsByCourseID(courseID int) ([]entities.Review, error) {
	return uc.Repo.GetReviewsByCourseID(courseID)
}

func (uc *CourseUseCase) DeleteReview(id int) error {
	return uc.Repo.DeleteReview(id)
}
