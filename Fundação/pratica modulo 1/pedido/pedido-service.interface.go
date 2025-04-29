package pedido

type PedidoServiceInterface interface {
	Create(clienteId int) (*Pedido, error)
	GetById(id string) (*Pedido, error)
	GetAll() (map[string]*Pedido, error)
}
