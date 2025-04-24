package provider

import (
	controllerUser "server/internal/controller"
	"server/internal/gateway/route"
	clientgrpc "server/internal/grpc/client"
	repositoryUser "server/internal/repository"
	usecaseUser "server/internal/usecase"
	"server/pkg/config"
	"server/pkg/database/transactor"
	"server/pkg/logger"

	"github.com/gin-gonic/gin"
)

var (
	userRepository repositoryUser.UserRepositoryInterface
)

var (
	userUseCase usecaseUser.UserUsecaseInterface
)

var (
	userController *controllerUser.UserController
)

var (
	GRPCUserClient *clientgrpc.UserClient
)

func gRPCProvideUserModule(cfg *config.Config) {
	if store == nil {
		store = transactor.NewTransactor(db)
	}
	client := clientgrpc.NewUserClient(cfg)
	if client == nil {
		logger.Log.Error("failed to create gRPC client")
	}
	GRPCUserClient = client
	injectUserModuleRepository()
	injectUserModuleUseCase()
}

func ProvideUserModule(router *gin.Engine, cfg *config.Config) {
	injectUserModuleRepository()
	injectUserModuleUseCase()
	injectUserModuleController()

	route.UserControllerRoute(userController, router)
}

func injectUserModuleRepository() {
	userRepository = repositoryUser.NewUserRepositoryImplementation(db)
}

func injectUserModuleUseCase() {
	userUseCase = usecaseUser.NewUserUsecaseImplementation(store, userRepository, GRPCUserClient)
}

func injectUserModuleController() {
	userController = controllerUser.NewUserController(userUseCase)
}
