package main

import (
	"chap3-challenge2/controller"
	"chap3-challenge2/middleware"
	"chap3-challenge2/model"
	"chap3-challenge2/repository"
	"chap3-challenge2/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	const (
		host         = "localhost"
		user         = "postgres"
		password     = "admin"
		databasePort = "5432"
		databaseName = "test"
	)

	fix := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, databasePort)

	db, err = gorm.Open(postgres.Open(fix), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&model.User{}, &model.Product{})
}

func main() {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(*productRepository)
	productController := controller.NewProductController(*productService)

	x := gin.Default()

	x.POST("/admin/register", userController.CreateAdmin) //ok
	x.POST("/user/register", userController.CreateUser)   //ok
	x.POST("/login", userController.Login)                //ok

	productGroup := x.Group("/admin", middleware.AuthMiddleware)     //ok
	productGroup.POST("/", productController.CreateProduct)          //ok
	productGroup.GET("/", productController.GetListProducts)         //ok
	productGroup.GET("/:id", productController.GetProductByID)       //ok
	productGroup.PUT("/:id", productController.UpdateProductByID)    //ok
	productGroup.DELETE("/:id", productController.DeleteProductById) //ok

	productUserGroup := x.Group("/user", middleware.AuthMiddleware) //ok
	productUserGroup.POST("/", productController.CreateProduct)     //ok
	productUserGroup.GET("/:id", productController.GetProductByID)  //ok
	productUserGroup.GET("/", productController.GetListProducts)    //ok

	err := x.Run("localhost:8083")
	if err != nil {
		panic(err)
	}
}
