package usecase

import (
	"context"
	"goNam/contract"
	"goNam/models"
	"time"

)

type UserUseCase struct {
	Repo    contract.Repository
	contextTimeout time.Duration
}

func (u *UserUseCase) GetByID(ctx context.Context, id int64) (*models.UserModel, error) {
	//panic("implement me")
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res,err := u.Repo.GetByID(ctx,id)
	if err != nil {
		return nil, err
	}

	return res,err
}

func (u *UserUseCase) DeleteById(ctx context.Context, id int64) error {
	//panic("implement me")
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.Repo.DeleteById(ctx,id)
	if err != nil {
		return err
	}

	return nil
}
//
func (u *UserUseCase) Fetch(ctx context.Context) ([]*models.UserModel, error) {
	//panic("implement me")
	res,err :=u.Repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return res,err
}

func (u *UserUseCase) Add(ctx context.Context, m *models.UserModel) (err error) {
	//panic("implement me")
	err = u.Repo.Add(ctx,m)
	if err != nil {
		return  err
	}

	return nil
}

func NewUserUC(r contract.Repository, timeout time.Duration) contract.UseCase {
	return &UserUseCase{
		Repo:           r,
		contextTimeout: timeout,
	}
}