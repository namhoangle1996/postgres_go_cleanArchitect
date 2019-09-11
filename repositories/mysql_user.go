package repositories

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"


	//"database/sql"
	"goNam/contract"
	"goNam/models"
	"golang.org/x/crypto/bcrypt"
	//"fmt"
)
const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

func test() {
	db, err := gorm.Open("postgres", "user=postgres dbname=test password=postgres port=5432 host=localhost sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	database := db.DB()

	err = database.Ping()
	if err != nil {
		panic(err)
	}
}

type mysqlRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlRepository(Conn *gorm.DB) contract.Repository {
	return &mysqlRepository{Conn}
}
// Hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// get users by ID
func (m *mysqlRepository) GetByID(ctx context.Context, id int64) (res *models.UserModel, err error) {
	//tmp:= new(models.UserModel)
    //user , err =  newUserOrm(m.Conn)
	//test()
	//var users *models.UserModel
	var result *models.UserModel

	//m.Conn.Debug().Where(&models.UserModel{Email: "namnd@gmail.com"}).First(&users.Email)
	m.Conn.Table("users").Select("email").Where("email = ?", "namnd@gmail.com").Scan(&result)
	m.Conn.LogMode(true)


	return nil , nil
}

// Del user by Id
func (m *mysqlRepository) DeleteById(ctx context.Context, id int64) (err error) {
	//_,err := m.Conn.Query("DELETE * FROM users WHERE id=$1", id)
	//sqlStatement := ` DELETE FROM users WHERE id = $1;`
	////_, err = m.Conn.Exec(sqlStatement, id)
	//if err != nil {
	//	panic(err)
	//}

	return nil
}

// Fetall users
func (m *mysqlRepository) Fetch(ctx context.Context) (res []*models.UserModel, err error) {
	//tmp:= new([]models.UserModel)


	//var result []*models.UserModel

	//m.Conn.Debug().Where(&models.UserModel{Email: "namnd@gmail.com"}).First(&users.Email)
	rows, err := m.Conn.Table("users").Select("id,email").Rows()

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
	m.Conn.LogMode(true)

	return result, nil

}

func (m *mysqlRepository) Add(ctx context.Context, a *models.UserModel) (err error) {
	//query := `INSERT INTO users (email,password) VALUES ($1, $2)`
	//hashedPass, _ := HashPassword(a.Password)
	//
	//_,err = m.Conn.Exec(query, a.Email, hashedPass)
	//if err != nil {
	//	return err
	//}
	//fmt.Print(query)
	return nil
}

