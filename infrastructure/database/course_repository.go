package database

import (
	"database/sql"

	"github.com/NaheedRayan/mini_rest_api_shikho/entities"
)

type CourseRepository struct {
	DB *sql.DB
}


func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB: db,
	}
}

//----------------------------------------------------------------user----------------------------------------------------------------

// create a new user
func (r *CourseRepository) CreateUser(user entities.User) (entities.User, error) {
	query := "INSERT INTO users (first_name, last_name, email , password , role , bio ) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.DB.Exec(query, user.FirstName,user.LastName,user.Email , user.Password , user.Role , user.Bio)
	if err != nil {
		return entities.User{}, err
	}
	id, _ := result.LastInsertId()
	user.ID = uint(id)
	return user, nil
}

// Get a user by ID
func (r *CourseRepository) GetUserByID(id int) (entities.User, error) {
	query := "SELECT id, first_name, last_name, email, password, role, bio FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var user entities.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Bio)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

// Get all users
func (r *CourseRepository) GetAllUsers() ([]entities.User, error) {
	query := "SELECT id, first_name, last_name, email, password, role, bio FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role, &user.Bio)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}


// Update a user
func (r *CourseRepository) UpdateUser(user entities.User) (entities.User ,error) {
	query := "UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ?, role = ?, bio = ? WHERE id = ?"
	_, err := r.DB.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.Bio, user.ID)
	return user, err
}

// Delete a user
func (r *CourseRepository) DeleteUser(id int) (error) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}


//----------------------------------------------------------------course----------------------------------------------------------------

// create a new course
func (r *CourseRepository) AddCourse(course entities.Course) (entities.Course, error) {
	query := "INSERT INTO courses (title, description, duration, price, instructor, category) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.DB.Exec(query, course.Title, course.Description, course.Duration, course.Price, course.Instructor, course.Category)
	if err != nil {
		return entities.Course{}, err
	}
	id, _ := result.LastInsertId()
	course.ID = uint(id)
	return course, nil
}

// Get a course by ID
func (r *CourseRepository) GetCourseByID(id int) (entities.Course, error) {
	query := "SELECT id, title, description, duration, price, instructor, category FROM courses WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var course entities.Course
	err := row.Scan(&course.ID, &course.Title, &course.Description, &course.Duration, &course.Price, &course.Instructor, &course.Category)
	if err != nil {
		return entities.Course{}, err
	}
	return course, nil
}

// Get all courses
func (r *CourseRepository) GetAllCourses() ([]entities.Course, error) {
	query := "SELECT id, title, description, duration, price, instructor, category FROM courses"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []entities.Course
	for rows.Next() {
		var course entities.Course
		err := rows.Scan(&course.ID, &course.Title, &course.Description, &course.Duration, &course.Price, &course.Instructor, &course.Category)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// Update a course
func (r *CourseRepository) UpdateCourse(course entities.Course) (entities.Course ,error) {
	query := "UPDATE courses SET title = ?, description = ?, duration = ?, price = ?, instructor = ?, category = ? WHERE id = ?"
	_, err := r.DB.Exec(query, course.Title, course.Description, course.Duration, course.Price, course.Instructor, course.Category, course.ID)
	return course, err
}

// Delete a course
func (r *CourseRepository) DeleteCourse(id int) (error) {
	query := "DELETE FROM courses WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

//----------------------------------------------------------------enrollment----------------------------------------------------------------

// Create a new enrollment
func (r *CourseRepository) AddEnrollment(enrollment entities.Enrollment) (entities.Enrollment, error) {
	query := "INSERT INTO enrollments (user_id, course_id, completed) VALUES (?, ?, ?)"
	result, err := r.DB.Exec(query, enrollment.UserID, enrollment.CourseID, enrollment.Completed)
	if err != nil {
		return entities.Enrollment{}, err
	}
	id, _ := result.LastInsertId()
	enrollment.ID = uint(id)
	return enrollment, nil
}

// Get all enrollments
func (r *CourseRepository) GetAllEnrollments() ([]entities.Enrollment, error) {
	query := "SELECT id, user_id, course_id, completed FROM enrollments"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []entities.Enrollment
	for rows.Next() {
		var enrollment entities.Enrollment
		err := rows.Scan(&enrollment.ID, &enrollment.UserID, &enrollment.CourseID, &enrollment.Completed)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}

// Get an enrollment by ID
func (r *CourseRepository) GetEnrollmentByID(id int) (entities.Enrollment, error) {
	query := "SELECT id, user_id, course_id, completed FROM enrollments WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var enrollment entities.Enrollment
	err := row.Scan(&enrollment.ID, &enrollment.UserID, &enrollment.CourseID, &enrollment.Completed)
	if err != nil {
		return entities.Enrollment{}, err
	}
	return enrollment, nil
}

// get all enrollments by user ID
func (r *CourseRepository) GetEnrollmentsByUserID(userID int) ([]entities.Enrollment, error) {
	query := "SELECT id, user_id, course_id, completed FROM enrollments WHERE user_id = ?"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []entities.Enrollment
	for rows.Next() {
		var enrollment entities.Enrollment
		err := rows.Scan(&enrollment.ID, &enrollment.UserID, &enrollment.CourseID, &enrollment.Completed)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}			

// Update an enrollment
func (r *CourseRepository) UpdateEnrollment(enrollment entities.Enrollment) (entities.Enrollment, error) {
	query := "UPDATE enrollments SET user_id = ?, course_id = ?, completed = ? WHERE id = ?"
	_, err := r.DB.Exec(query, enrollment.UserID, enrollment.CourseID, enrollment.Completed, enrollment.ID)
	if err != nil {
		return entities.Enrollment{}, err
	}
	return enrollment, nil
}

// Delete an enrollment				
func (r *CourseRepository) DeleteEnrollment(id int) error {
	query := "DELETE FROM enrollments WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

//----------------------------------------------------------------lesson----------------------------------------------------------------

// Create a new lesson
func (r *CourseRepository) AddLesson(lesson entities.Lesson) (entities.Lesson, error) {
	query := "INSERT INTO lessons (course_id, title, content, video_url, `order`) VALUES (?, ?, ?, ?, ?)"
	result, err := r.DB.Exec(query, lesson.CourseID, lesson.Title, lesson.Content, lesson.VideoURL, lesson.Order)
	if err != nil {
		return entities.Lesson{}, err
	}
	id, _ := result.LastInsertId()
	lesson.ID = uint(id)
	return lesson, nil
}



// Get all lessons by course ID
func (r *CourseRepository) GetLessonsByCourseID(courseID int) ([]entities.Lesson, error) {
	query := "SELECT id, course_id, title, content, video_url, `order` FROM lessons WHERE course_id = ? ORDER BY `order`"
	rows, err := r.DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []entities.Lesson
	for rows.Next() {
		var lesson entities.Lesson
		err := rows.Scan(&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.VideoURL, &lesson.Order)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

//get all lessons by course ID 
func (r *CourseRepository) GetLessonsByID(id int) ([]entities.Lesson, error) {
	query := "SELECT id, course_id, title, content, video_url, `order` FROM lessons WHERE id = ?"
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var lessons []entities.Lesson
	for rows.Next() {
		var lesson entities.Lesson
		err := rows.Scan(&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.VideoURL, &lesson.Order)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

// Update a lesson
func (r *CourseRepository) UpdateLesson(lesson entities.Lesson) (entities.Lesson, error) {
	query := "UPDATE lessons SET course_id = ?, title = ?, content = ?, video_url = ?, `order` = ? WHERE id = ?"
	_, err := r.DB.Exec(query, lesson.CourseID, lesson.Title, lesson.Content, lesson.VideoURL, lesson.Order, lesson.ID)
	if err != nil {
		return entities.Lesson{}, err
	}
	return lesson, nil
}

// Delete a lesson
func (r *CourseRepository) DeleteLesson(id int) error {
	query := "DELETE FROM lessons WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

//----------------------------------------------------------------progress----------------------------------------------------------------

// Create a new progress
func (r *CourseRepository) AddProgress(progress entities.Progress) (entities.Progress, error) {
	query := "INSERT INTO progress (enrollment_id, lesson_id, completed) VALUES (?, ?, ?)"
	result, err := r.DB.Exec(query, progress.EnrollmentID, progress.LessonID, progress.Completed)
	if err != nil {
		return entities.Progress{}, err
	}
	id, _ := result.LastInsertId()
	progress.ID = uint(id)
	return progress, nil
}

// Update lesson progress for a user
func (r *CourseRepository) UpdateProgress(progress entities.Progress) (entities.Progress, error) {
	query := "UPDATE progress SET completed = ? WHERE enrollment_id = ? AND lesson_id = ?"
	_, err := r.DB.Exec(query, progress.Completed, progress.EnrollmentID, progress.LessonID)
	if err != nil {
		return entities.Progress{}, err
	}
	return progress, nil
}

// Get progress by enrollment ID and lesson ID
func (r *CourseRepository) GetProgressByEnrollmentAndLesson(enrollmentID, lessonID int) (entities.Progress, error) {
	query := "SELECT id, enrollment_id, lesson_id, completed FROM progress WHERE enrollment_id = ? AND lesson_id = ?"
	row := r.DB.QueryRow(query, enrollmentID, lessonID)

	var progress entities.Progress
	err := row.Scan(&progress.ID, &progress.EnrollmentID, &progress.LessonID, &progress.Completed)
	if err != nil {
		return entities.Progress{}, err
	}
	return progress, nil
}

//----------------------------------------------------------------review----------------------------------------------------------------

// Add a new review
func (r *CourseRepository) AddReview(review entities.Review) (entities.Review, error) {
	query := "INSERT INTO reviews (course_id, user_id, rating, comment) VALUES (?, ?, ?, ?)"
	result, err := r.DB.Exec(query, review.CourseID, review.UserID, review.Rating, review.Comment)
	if err != nil {
		return entities.Review{}, err
	}
	id, _ := result.LastInsertId()
	review.ID = uint(id)
	return review, nil
}

// Get all reviews for a course
func (r *CourseRepository) GetReviewsByCourseID(courseID int) ([]entities.Review, error) {
	query := "SELECT id, course_id, user_id, rating, comment FROM reviews WHERE course_id = ?"
	rows, err := r.DB.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []entities.Review
	for rows.Next() {
		var review entities.Review
		err := rows.Scan(&review.ID, &review.CourseID, &review.UserID, &review.Rating, &review.Comment)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// Delete a review
func (r *CourseRepository) DeleteReview(id int) error {
	query := "DELETE FROM reviews WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	return err
}

