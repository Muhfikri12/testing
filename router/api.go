package router

import (
	"ecommers/database"
	"ecommers/handler"
	"ecommers/repository"
	"ecommers/service"
	"ecommers/util"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func InitRouter() (*chi.Mux, *zap.Logger, error) {
	// Inisialisasi router
	r := chi.NewRouter()

	// Inisialisasi logger
	logger := util.InitLog()

	// Membaca konfigurasi
	config, err := util.ReadConfiguration()
	if err != nil {
		logger.Error("Failed to read configuration", zap.Error(err))
		return nil, logger, err
	}

	// Inisialisasi database
	db, err := database.InitDB(config)
	if err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
		return nil, logger, err
	}

	// Inisialisasi repository, service, dan handler
	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	handler := handler.NewAllHandler(service, logger, config)

	// Menambahkan endpoint ke router
	r.Route("/api", func(api chi.Router) {
		api.Get("/products", handler.ProductHandler.GetAll)
		api.Get("/products/best_selling", handler.ProductHandler.GetAllBestSelling)
	})

	return r, logger, nil
}
