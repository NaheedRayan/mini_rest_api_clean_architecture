

# **Mini REST API - Clean Architecture**

A feature-rich REST API prototype designed for managing an online learning platform, built with Go, using **Gin** as the web framework and **Sqlite3** for database interactions. This API allows course management, user management, enrollment handling, and more.

It Follows `Clean Architecture`
For more details, visit the [Clean Architecture in Golang](https://github.com/bxcodec/go-clean-arch).

![clear architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)


---

## **Key Features**

### **1. Course Management**
- Perform CRUD operations for courses.
- Store details such as:
  - Course title
  - Description
  - Price
  - Instructor name
  - Category (e.g., programming, design, business).

### **2. User Management**
- Register and manage users (students and instructors).
- Support roles such as `student` or `instructor`.
- Include additional user attributes like bio (optional for instructors).

### **3. Enrollment Management**
- Enroll users in courses.
- Retrieve:
  - Courses a user is enrolled in.
  - Students enrolled in a specific course.

### **4. Lessons and Progress Tracking**
- Associate multiple lessons with courses.
- Track user progress at the lesson level.

### **5. Course Reviews**
- Add and retrieve course reviews.
- Include ratings (1-5 stars) and comments.



---



### File structure
```
course-management/
├── entities/               # Core business entities
│   └── course.go
├── usecases/               # Business logic (use cases)
│   └── course_usecase.go
├── interfaces/             # Controllers (HTTP handlers)
│   └── course_handler.go
├── infrastructure/         # External systems (DB and routing)
│   ├── database/
│   │   └── course_repository.go
│   └── router.go
├── config/                 # Configuration files (DB setup)
│   └── database.go
└── main.go                 # Application entry point
```

Here’s a polished version of your README file, designed to be more professional and structured for clarity:

---

## **API Endpoints**

### **Course Endpoints**
| Method | Endpoint         | Description                        |
|--------|------------------|------------------------------------|
| POST   | `/course`        | Add a new course.                 |
| GET    | `/courses`       | Retrieve all courses.             |
| GET    | `/course/{id}`   | Retrieve details of a course.     |
| PUT    | `/course`        | Update course details.            |
| DELETE | `/course/{id}`   | Delete a course.                  |

### **User Endpoints**
| Method | Endpoint         | Description                        |
|--------|------------------|------------------------------------|
| POST   | `/user`          | Register a new user.              |
| GET    | `/users`         | List all users.                   |
| GET    | `/user/{id}`     | Retrieve details of a user.       |
| PUT    | `/user`          | Update user details.              |
| DELETE | `/user/{id}`     | Delete a user.                    |

### **Enrollment Endpoints**
| Method | Endpoint             | Description                        |
|--------|----------------------|------------------------------------|
| POST   | `/enroll`            | Enroll a user in a course.        |
| GET    | `/enrolls`           | Retrieve all enrollments.         |
| GET    | `/enroll/:id`        | Retrieve enrollment by ID.        |
| GET    | `/enrolls/user/{id}` | Get all enrollments for a user.   |
| PUT    | `/enroll`            | Update enrollment details.        |
| DELETE | `/enroll/:id`        | Delete an enrollment.             |

### **Lesson Endpoints**
| Method | Endpoint             | Description                        |
|--------|----------------------|------------------------------------|
| POST   | `/lesson`            | Add a new lesson to a course.     |
| GET    | `/lessons/course/{id}` | Get all lessons for a course.    |
| GET    | `/lesson/{id}`       | Retrieve details of a lesson.     |
| PUT    | `/lesson`            | Update lesson details.            |
| DELETE | `/lesson/{id}`       | Delete a lesson.                  |

### **Progress Tracking**
| Method | Endpoint                        | Description                                 |
|--------|---------------------------------|---------------------------------------------|
| POST   | `/progress`                     | Add progress for a lesson.                 |
| PUT    | `/progress`                     | Update progress details.                   |
| GET    | `/progress/{enrollmentID}/{lessonID}` | Get progress for a specific lesson. |

### **Review Endpoints**
| Method | Endpoint              | Description                        |
|--------|-----------------------|------------------------------------|
| POST   | `/review`             | Add a review to a course.         |
| GET    | `/reviews/course/{id}` | Retrieve reviews for a course.   |
| DELETE | `/review/{id}`        | Delete a review.                  |

---



## **Data Models**
```go
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


```



## **Usage**

Here’s the revised project setup with the correct module path:

---

### **Project Setup**

#### **1. Initialize the Project**
Create a directory for your project and initialize the Go module:

```bash
mkdir mini_rest_api_clean_architecture
cd mini_rest_api_clean_architecture
go mod init github.com/username/mini_rest_api_clean_architecture
```

---

#### **2. Install Dependencies**
Install the required libraries:

```bash
# Web framework
go get -u github.com/gin-gonic/gin

# SQLite3 driver
go get github.com/mattn/go-sqlite3

```

---


#### **3. Run the Application**
1. Ensure the database file is properly initialized.
2. Start your Go server:
   ```bash
   go run main.go
   ```
3. Access the API endpoints using tools like Postman or Curl. The base URL is `http://localhost:8080`.


