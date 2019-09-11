package repositories

import (
	"context"
	"database/sql"
	"goNam/contract"
	"goNam/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"

)
const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlRepository(Conn *sql.DB) contract.Repository {
	return &mysqlRepository{Conn}
}
// Hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// get users by ID
func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res *models.UserModel, err error) {
	tmp:= new(models.UserModel)
	err = m.Conn.QueryRow("SELECT id ,email FROM users WHERE id=$1", id).Scan(&tmp.Id,&tmp.Email)
	if err != nil {
		return nil, err
	}
	//fmt.Print(tmp)

	res =tmp

	return
}

// Del user by Id
func (m *mysqlRepository) DeleteById(ctx context.Context, id int64) (err error) {
	//_,err := m.Conn.Query("DELETE * FROM users WHERE id=$1", id)
	sqlStatement := ` DELETE FROM users WHERE id = $1;`
	_, err = m.Conn.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}

	return nil
}

// Fetall users
func (m *mysqlRepository) Fetch(ctx context.Context) (res []*models.UserModel, err error) {
	//tmp:= new([]models.UserModel)

	rows, err := m.Conn.Query("SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}
	result := make([]*models.UserModel, 0)

	defer rows.Close()
	for rows.Next() {
		t := new(models.UserModel)

		err = rows.Scan(&t.Id, &t.Email)
		if err != nil {
			panic(err)
		}

		result = append(result, t)
	}
	return result, nil

}

func (m *mysqlRepository) Add(ctx context.Context, a *models.UserModel) (err error) {
	query := `INSERT INTO users (email,password) VALUES ($1, $2)`
	hashedPass, _ := HashPassword(a.Password)

	_,err = m.Conn.Exec(query, a.Email, hashedPass)
	if err != nil {
		return err
	}
	fmt.Print(query)
	return nil
}
