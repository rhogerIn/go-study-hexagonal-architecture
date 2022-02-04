package main

import (
	"hexagonArchitecture/internals/core/services"
	"hexagonArchitecture/internals/handlers"
	"hexagonArchitecture/internals/repositories"
	"hexagonArchitecture/internals/server"
)

func main() {
	mongoConn := "localhost"

	userRepository := repositories.NewUserRepository(mongoConn)
	// services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandlers(userService)
	//server
	httpServer := server.NewServer(
		userHandlers,
	)
	httpServer.Initialize()
}