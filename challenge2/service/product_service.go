package service

import (
	"chap3-challenge2/helper"
	"chap3-challenge2/model"
	"chap3-challenge2/repository"
	"errors"

	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (ps *ProductService) CreateProduct(request model.ProductCreateRequest, userId string) (model.ProductCreateResponse, error) {
	id := helper.GenerateID()
	product := model.Product{
		ID:     id,
		UserID: userId,
		Name:   request.Name,
		Price:  request.Price,
	}

	res, err := ps.ProductRepository.Add(product)
	if err != nil {
		return model.ProductCreateResponse{}, err
	}

	return model.ProductCreateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Name:      res.Name,
		Price:     res.Price,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (ps *ProductService) GetAll(userId string, role string) ([]model.ProductResponse, error) {
	productResponse := []model.ProductResponse{}
	res, err := ps.ProductRepository.Get()

	if !helper.IsAdminTrue(role) {
		res, err = ps.ProductRepository.GetByUserId(userId)
	}

	if err != nil {
		return []model.ProductResponse{}, err
	}

	for _, val := range res {
		productResponse = append(productResponse, model.ProductResponse{
			ID:        val.ID,
			UserID:    val.UserID,
			Name:      val.Name,
			Price:     val.Price,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return productResponse, nil

}

func (ps *ProductService) GetById(id string, role string, userId string) (model.ProductResponse, error) {

	res, err := ps.ProductRepository.GetOne(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.ProductResponse{}, model.ErrorNotFound
	}

	if !helper.IsAdminTrue(role) && res.UserID != userId {
		return model.ProductResponse{}, model.ErrorNotAuthorized
	}

	return model.ProductResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Name:      res.Name,
		Price:     res.Price,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (ps *ProductService) UpdateById(request model.ProductUpdateRequest, id string, role string) (model.ProductResponse, error) {
	if !helper.IsAdminTrue(role) {
		return model.ProductResponse{}, model.ErrorNotAuthorized
	}

	product := model.Product{
		Name:  request.Name,
		Price: request.Price,
	}

	res, err := ps.ProductRepository.UpdateOne(product, id)

	if err != nil {
		return model.ProductResponse{}, err
	}

	return model.ProductResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Name:      res.Name,
		Price:     res.Price,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (ps *ProductService) DeleteById(id string, role string) error {
	product := model.Product{}
	if !helper.IsAdminTrue(role) {
		return model.ErrorNotAuthorized
	}

	_, err := ps.ProductRepository.GetOne(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.ErrorNotFound
	}

	err = ps.ProductRepository.DeleteOne(product, id)
	if err != nil {
		return err
	}

	return nil
}
