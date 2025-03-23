package container

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/env"
	rdbms "github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/postgres"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

type Container struct {
	Config                 env.Config
	TremLigeiroDB          rdbms.RDBMS
	ProductRepository      repository.IProductRepository
	OrderRepository        repository.IOrderRepository
	OrderProductRepository repository.IOrderProductRepository
	CustomerRepository     repository.ICustomerRepository
	CategoryRepository     repository.ICategoryRepository
	PaymentRepository      repository.IPaymentRepository
	PaymentService         external.IPaymentService
}

func New(config env.Config) (*Container, error) {
	factory := Container{}
	factory.Config = config

	return &factory, nil
}

func (container *Container) Start() error {

	err := postgres.Migrate(getPostgreSQLConf(container.Config))
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
	}
	container.TremLigeiroDB, err = postgres.New(getPostgreSQLConf(container.Config))
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return err
	}

	container.ProductRepository = repository.NewProductRepository(container.TremLigeiroDB)
	container.OrderRepository = repository.NewOrderRepository(container.TremLigeiroDB)
	container.CustomerRepository = repository.NewCustomerRepository(container.TremLigeiroDB)
	container.CategoryRepository = repository.NewCategoryRepository()
	container.OrderProductRepository = repository.NewOrderProductRepository(container.TremLigeiroDB)
	container.PaymentRepository = repository.NewPaymentRepository(container.TremLigeiroDB)
	container.PaymentService = external.NewPaymentService(getPaymentConf(container.Config))

	return nil
}

func (container *Container) Stop() error {
	db, err := container.TremLigeiroDB.DB.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func getPostgreSQLConf(config env.Config) postgres.PostgreSQLConf {
	return postgres.PostgreSQLConf{
		User:   config.DbUser,
		Pass:   config.DbPassword,
		Url:    config.DbHost,
		Port:   config.DbPort,
		DbName: config.DbName,
	}
}

func getPaymentConf(config env.Config) external.PaymentConfig {
	return external.PaymentConfig{
		Url:   config.PaymentUrl,
		Token: config.PaymentAuthToken,
	}
}
