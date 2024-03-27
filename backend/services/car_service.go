package services

type CarServiceProvider interface {
	GetCars() ([]Car, error)
	GetCar(id string) (Car, error)
	CreateCar(car Car) (Car, error)
	UpdateCar(id string, car Car) (Car, error)
	DeleteCar(id string) error
}

type CarService struct {
	Cars []*Car
}

func NewCarService() CarServiceProvider {
	return &CarService{
		Cars: []*Car{},
	}
}
