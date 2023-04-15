package controller

import (
	"chap3-challenge2/model"
	"chap3-challenge2/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	request := model.ProductCreateRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}
	res, err := pc.ProductService.CreateProduct(request, userId.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, res)
	return

}

func (pc *ProductController) GetListProducts(ctx *gin.Context) {
	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}
	role, isExist := ctx.Get("role")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := pc.ProductService.GetAll(userId.(string), role.(string))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}

func (pc *ProductController) GetProductByID(ctx *gin.Context) {
	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}
	role, isExist := ctx.Get("role")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	id := ctx.Param("id")
	res, err := pc.ProductService.GetById(id, role.(string), userId.(string))

	if err != nil {
		if err != model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.MyError{
				Err: model.ErrorNotAuthorized.Err,
			})
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.MyError{
				Err: model.ErrorNotFound.Err,
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, res)
	return

}

func (pc *ProductController) UpdateProductByID(ctx *gin.Context) {
	request := model.ProductUpdateRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	role, isExist := ctx.Get("role")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	id := ctx.Param("id")
	res, err := pc.ProductService.UpdateById(request, id, role.(string))

	if err != nil {
		if err == model.ErrorNotAuthorized {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.MyError{
				Err: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}

func (pc *ProductController) DeleteProductById(ctx *gin.Context) {
	role, isExist := ctx.Get("role")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	id := ctx.Param("id")
	err := pc.ProductService.DeleteById(id, role.(string))

	if err != nil {
		if err == model.ErrorNotAuthorized {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.MyError{
				Err: model.ErrorNotAuthorized.Err,
			})
			return
		} else if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.MyError{
				Err: model.ErrorNotFound.Err,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Berhasil menghapus id %s", id)
	ctx.JSON(http.StatusOK, model.ResponseMessage{
		Message: message,
	})
	return

}
