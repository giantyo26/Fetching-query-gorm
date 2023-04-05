package repositories

import (
	"dumbmerch/models"
	"time"

	"gorm.io/gorm"
)

// declaration of the UserRepository interface, which defines two methods: FindUsers() and GetUser(ID int)
type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

// repository struct, which implements the UserRepository interface.
type repository struct {
	db *gorm.DB // pointer to a GORM database connection
}

// constructor function for the repository struct. It takes a *gorm.DB as an argument
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db} // returns a pointer to a new repository struct initialized with the provided database connection.
}

// queries the "users" table in the database and scans the results into a slice of Users models.
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

// It queries the "users" table with a parameterized query to fetch a single user by their ID. It scans the result into a User model
func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("UPDATE users SET name=?, email=?, password=?, updated_at=? WHERE id=?", user.Name, user.Email, user.Password, time.Now(), ID).Scan(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users where ID=?", ID).Scan(&user).Error

	return user, err
}
