package pedido

import "github.com/google/uuid"

func NewPedidoService() PedidoServiceInterface {
	return &pedidoService{}
}

type pedidoService struct {
}

func (p *pedidoService) Create(clienteId int) (*Pedido, error) {
	return &Pedido{
		Id:        uuid.New(),
		ClienteId: clienteId,
		OrderCode: uuid.NewString(),
	}, nil
}

func (p *pedidoService) GetById(id string) (*Pedido, error) {
	return &Pedido{
		Id:        uuid.MustParse(id),
		ClienteId: 1,
		OrderCode: uuid.NewString(),
	}, nil
}

func (p *pedidoService) GetAll() (map[string]*Pedido, error) {
	id1 := uuid.NewString()
	id2 := uuid.NewString()
	return map[string]*Pedido{
		id1: {
			Id:        uuid.MustParse(id1),
			ClienteId: 1,
			OrderCode: uuid.NewString(),
		},
		id2: {
			Id:        uuid.MustParse(id2),
			ClienteId: 2,
			OrderCode: uuid.NewString(),
		},
	}, nil
}
