package cmd

import (
	"ecommerece/config"
	"ecommerece/infra/db"
	"ecommerece/product"
	"ecommerece/repo"
	"ecommerece/rest"
	prdctHandler "ecommerece/rest/handlers/product"
	"ecommerece/user"
	"fmt"
	"os"

	usrHandler "ecommerece/rest/handlers/user"
	middleware "ecommerece/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	//fmt.Println("%+v", cnf.DB)

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	// handlers
	productHandler := prdctHandler.NewHandler(middlewares, prdctSvc)
	userHandler := usrHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
