package contract
import (
	"context"
	"goNam/models"
)

// Repository represent the article's repository contract
type UseCase interface {
	Fetch(ctx context.Context) (res []*models.UserModel, err error)
	GetByID(ctx context.Context, id int64) (*models.UserModel, error)
	DeleteById(ctx context.Context, id int64) error
	Add(ctx context.Context, model *models.UserModel) error

}
