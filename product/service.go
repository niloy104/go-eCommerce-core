package product

import "ecommerece/domain"

type service struct {
	prdctRepo ProductRepo
}

func NewService(prdctRepo ProductRepo) Service {
	return &service{
		prdctRepo: prdctRepo,
	}
}

func (svc *service) Create(prdct domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Create(prdct)
}

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prdctRepo.Get(id)
}

func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	return svc.prdctRepo.List(page, limit)
}

func (svc *service) Count() (int64, error) {
	return svc.prdctRepo.Count()
}

func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Update(prdct)
}

func (svc *service) Delete(id int) error {
	return svc.prdctRepo.Delete(id)
}
