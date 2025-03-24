package external

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/tbtec/tremligeiro/internal/infra/httpclient"
)

type IPaymentService interface {
	RequestPayment(ctx context.Context, request PaymentRequest) (PaymentResponse, error)
}

type PaymentService struct {
	httpclient *resty.Client
	config     PaymentConfig
}

func NewPaymentService(config PaymentConfig) IPaymentService {
	return &PaymentService{
		config:     config,
		httpclient: httpclient.New(),
	}
}

func (service *PaymentService) RequestPayment(ctx context.Context, request PaymentRequest) (PaymentResponse, error) {
	paymentResponse := PaymentResponse{}

	url := service.config.Url
	token := "Bearer " + service.config.Token

	response, err := service.httpclient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(request).
		SetResult(&paymentResponse).
		Post(url)
	if err != nil {
		return paymentResponse, err
	}

	if response.StatusCode() != 201 {
		return PaymentResponse{}, response.Error().(error)
	}

	return paymentResponse, nil
}
