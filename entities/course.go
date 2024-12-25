package entities

//this is entites or models or domain


type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"` // hashed password
	Role string `json:"role"` // student or instructor
	Bio string `json:"bio"` // instructor bio optional
}


type Course struct {

	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Description string `json:"description"`
	Duration string `json:"duration"`
	Price float64 `json:"price"`
	Instructor string `json:"instructor"` 
	Category string `json:"category"` // programming, design, business, etc
}

//enrollment is the join table between users and courses
type Enrollment struct {
	ID uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"user_id"`
	CourseID uint `json:"course_id"`
	Completed bool `json:"completed"`
}


//courses may have multiple lessons	
type Lesson struct {
	ID uint `json:"id" gorm:"primary_key"`
	CourseID uint `json:"course_id"` 
	Title string `json:"title"`
	Content string `json:"content"`
	VideoURL string `json:"video_url"`
	Order uint `json:"order"` 
}

//progress is the join table between enrollments and lessons
type Progress struct {
    ID        uint `json:"id" gorm:"primaryKey"`
    EnrollmentID uint `json:"enrollment_id"`
    LessonID  uint `json:"lesson_id"`
    Completed bool `json:"completed"`
}

//reviews are associated with courses
type Review struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    CourseID uint   `json:"course_id"`
    UserID   uint   `json:"user_id"`
    Rating   int    `json:"rating"` // e.g., 1-5 stars
    Comment  string `json:"comment"`
}

