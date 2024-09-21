//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/entity"
	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/event"
	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/infra/database"
	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/infra/web"
	"github.com/vinicius-gregorio/fc_clean_arch_go/internal/usecase"
	"github.com/vinicius-gregorio/fc_clean_arch_go/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventInterface), new(*event.OrderListed)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

var setOrderListedEvent = wire.NewSet(
	event.NewOrderListed,
	wire.Bind(new(events.EventInterface), new(*event.OrderListed)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		usecase.NewListOrderUseCase,
	)
	return &usecase.ListOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
func NewWebOrderListHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderListHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderListedEvent,
		web.NewWebOrderListHandler,
	)
	return &web.WebOrderListHandler{}
}
