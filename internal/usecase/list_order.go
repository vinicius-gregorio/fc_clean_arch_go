package usecase

import (
	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/entity"
	"github.com/vinicius-gregorio/fc_clean_arch_go/pkg/events"
)

type ListOrderInputDTO struct {
}

type ListOrderOutputDTO struct {
	Orders []entity.Order `json:"orders"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrderUseCase) Execute(input ListOrderInputDTO) (ListOrderOutputDTO, error) {

	orders, err := l.OrderRepository.GetAll()
	if err != nil {
		return ListOrderOutputDTO{}, err
	}

	dto := ListOrderOutputDTO{
		Orders: orders,
	}
	l.OrderListed.SetPayload(dto)
	l.EventDispatcher.Dispatch(l.OrderListed)

	return dto, nil
}
