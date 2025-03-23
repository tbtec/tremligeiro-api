package gateway

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

const (
	SPONSOR_ID = 96469944
)

type PaymentGateway struct {
	paymentService    external.IPaymentService
	paymentRepository repository.IPaymentRepository
}

func NewPaymentGateway(paymentService external.IPaymentService,
	paymentRepository repository.IPaymentRepository) *PaymentGateway {
	return &PaymentGateway{
		paymentService:    paymentService,
		paymentRepository: paymentRepository,
	}
}

func (gtw *PaymentGateway) RequestPayment(ctx context.Context, order entity.Order, orderProduct []entity.OrderProduct, payment *entity.Payment, metadata dto.MetaData) error {

	items := make([]external.Item, 0)

	for _, op := range orderProduct {
		items = append(items, external.Item{
			SkuNumber:   op.ProductID,
			Category:    "marketplace",
			Title:       "Point Trem Ligeiro",
			Description: "Point Trem Ligeiro",
			Quantity:    op.Quantity,
			UnitPrice:   op.Amount,
			TotalAmount: op.TotalAmount,
			UnitMeasure: "unit",
		})
	}

	paymentRequest := external.PaymentRequest{
		ExternalReference: payment.ID,
		Title:             "Trem Ligeiro Payment",
		Description:       "Trem Ligeiro Payment",
		NotificationURL:   *metadata.PaymentWebHookUrl,
		TotalAmount:       order.TotalAmount,
		Item:              items,
		Sponsor:           external.Sponsor{ID: SPONSOR_ID},
		CashOut:           external.CashOut{Amount: 0},
	}

	slog.InfoContext(ctx, "Requesting payment...")

	response, err := gtw.paymentService.RequestPayment(ctx, paymentRequest)
	if err != nil {
		slog.ErrorContext(ctx, "❌ Error requesting payment", "error", err)
		return err
	}

	slog.InfoContext(ctx, "✅ Request payment succesfully", "response", response)

	payment.QrData = response.QRData
	payment.ExternalId = response.InStoreOrderId

	return nil
}

func (gtw *PaymentGateway) Create(ctx context.Context, payment *entity.Payment) error {

	paymentModel := model.Payment{
		ID:         payment.ID,
		OrderId:    payment.OrderId,
		Status:     string(payment.Status),
		QrData:     payment.QrData,
		ExternalId: payment.ExternalId,
		CreatedAt:  payment.CreatedAt,
		UpdatedAt:  payment.UpdatedAt,
	}

	gtw.paymentRepository.Create(ctx, &paymentModel)

	return nil
}

func (gtw *PaymentGateway) FindOne(ctx context.Context, id string) (*entity.Payment, error) {
	paymentModel, err := gtw.paymentRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	payment := entity.Payment{
		ID:         paymentModel.ID,
		OrderId:    paymentModel.OrderId,
		Status:     entity.PaymentStatus(paymentModel.Status),
		QrData:     paymentModel.QrData,
		ExternalId: paymentModel.ExternalId,
		CreatedAt:  paymentModel.CreatedAt,
		UpdatedAt:  paymentModel.UpdatedAt,
	}

	return &payment, nil
}

func (gtw *PaymentGateway) Update(ctx context.Context, payment *entity.Payment) error {

	paymentModel := model.Payment{
		ID:         payment.ID,
		OrderId:    payment.OrderId,
		Status:     string(payment.Status),
		QrData:     payment.QrData,
		ExternalId: payment.ExternalId,
		CreatedAt:  payment.CreatedAt,
		UpdatedAt:  payment.UpdatedAt,
	}

	gtw.paymentRepository.Update(ctx, &paymentModel)

	return nil
}

func (gtw *PaymentGateway) FindByOrderId(ctx context.Context, id string) (*entity.Payment, error) {
	payment := entity.Payment{}
	paymentModel, err := gtw.paymentRepository.FindByOrderId(ctx, id)
	if err != nil {
		return nil, err
	}

	if paymentModel != nil {
		payment = entity.Payment{
			ID:         paymentModel.ID,
			OrderId:    paymentModel.OrderId,
			Status:     entity.PaymentStatus(paymentModel.Status),
			QrData:     paymentModel.QrData,
			ExternalId: paymentModel.ExternalId,
			CreatedAt:  paymentModel.CreatedAt,
			UpdatedAt:  paymentModel.UpdatedAt,
		}
	}

	return &payment, nil
}
