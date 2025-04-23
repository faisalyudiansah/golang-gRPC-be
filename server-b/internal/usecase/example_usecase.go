package usecase

import (
	"context"
	"server/internal/repository"
	"server/pkg/database/transactor"
)

type ExampleUsecaseInterface interface {
	GetUsecase(ctx context.Context) error
}

type ExampleUsecaseImplementation struct {
	transactor        transactor.Transactor
	exampleRepository repository.ExampleRepositoryInterface
}

func NewExampleUsecaseImplementation(
	trx transactor.Transactor,
	exampleRepo repository.ExampleRepositoryInterface,
) *ExampleUsecaseImplementation {
	return &ExampleUsecaseImplementation{
		transactor:        trx,
		exampleRepository: exampleRepo,
	}
}

func (ir *ExampleUsecaseImplementation) GetUsecase(ctx context.Context) error {
	err := ir.transactor.Atomic(ctx, func(cForTx context.Context) error {
		return ir.exampleRepository.Get()
	})
	return err
}
