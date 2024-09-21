package web

import (
	"encoding/json"
	"net/http"

	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/entity"
	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/usecase"
	"github.com/vinicius-gregorio/fc_clean_arch_go/pkg/events"
)

type WebOrderListHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	OrderListedEvent events.EventInterface
}

func NewWebOrderListHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderListedEvent events.EventInterface,
) *WebOrderListHandler {
	return &WebOrderListHandler{
		EventDispatcher:  EventDispatcher,
		OrderRepository:  OrderRepository,
		OrderListedEvent: OrderListedEvent,
	}
}

func (h *WebOrderListHandler) Create(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrderUseCase(h.OrderRepository, h.OrderListedEvent, h.EventDispatcher)
	output, err := listOrders.Execute(usecase.ListOrderInputDTO{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
