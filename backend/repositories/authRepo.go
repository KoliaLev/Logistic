package repositories

import (
	"fmt"
	"logistics/models"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

type AuthRepo interface {
	Login()
	Register(user *models.User) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type authRepo struct {
	db *pg.DB
}

func NewAuthRepo(db *pg.DB) AuthRepo {
	return &authRepo{
		db: db}
}

func (r authRepo) Login() {
	fmt.Println("in the authRepo - login")
}

func (r authRepo) Register(user *models.User) (*models.User, error) {
	fmt.Println("in the authRepo - register ", user)

	_, err := r.db.Model(user).Insert()
	if err != nil {
		logrus.Info("err inserting to db ", err)
		return nil, err
	}
	return user, nil
}

func (r authRepo) GetByEmail(email string) (*models.User, error) {
	// fmt.Println("in the authRepo - register ", user)
	var user models.User
	err := r.db.Model(&user).Where("email = ?", email).Select()
	if err != nil {
		logrus.Info("err getting by email from db ", err)
		return nil, err
	}
	return &user, nil
}
