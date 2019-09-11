package contract
import (
	"context"
	"goNam/models"
)

// Repository represent the article's repository contract
type Repository interface {
	Fetch(ctx context.Context) (res []*models.UserModel, err error)
	GetByID(ctx context.Context, id int64) (*models.UserModel, error)
	//GetByTitle(ctx context.Context, title string) (*models.UserModel, error)
	//Update(ctx context.Context, ar *models.UserModel) error
	Add(ctx context.Context, model *models.UserModel) error
	DeleteById(ctx context.Context, id int64) error
}
