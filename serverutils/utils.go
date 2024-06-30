package serverutils

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"jacq/config"
	"jacq/encrypt"
	handler2 "jacq/handler"
	"jacq/idgenerator"
	server2 "jacq/server"
	"jacq/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func SetUpDatabase(url, name string) (storage.DataStore, *mongo.Client) {
	repo, client, err := storage.New(url, name)
	if err != nil {
		log.Fatalf("Error failed to open MongoDB: %v", err)
	}
	return repo, client
}

func SetUpHandler(store storage.DataStore) handler2.UserHandler {
	return handler2.NewUserHandler(store, encrypt.NewPasswordEncryptor(), idgenerator.New())
}

func SetUpServer(userHandler *handler2.UserHandler) server2.UserServer {
	return server2.NewUserServer(userHandler)
}

func SetupRouter(server *server2.UserServer) *gin.Engine {
	router := gin.Default()

	// Add Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.POST("user/signup", server.SignUp())
	router.POST("user/email_verification", server.SendVerificationEmail())
	router.PUT("user/update_record", server.UpdateUser())
	router.PUT("user/secure_transaction", server.SecureTransaction())
	router.POST("user/login", server.Login())
	router.POST("user/forgot_password", server.ForgotPassword())
	router.PUT("user/change_password", server.ChangePassword())
	router.POST("user/verify_phone_number", server.VerifyPhoneNumber())
	router.PUT("user/change_pin", server.ChangeTransactionPin())
	router.POST("user/verify_token", server.VerifyToken())
	router.POST("user/verify_bvn", server.VerifyBvn())
	return router
}

func SetupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func StartServer(router *gin.Engine, client *mongo.Client) {
	var c config.Config
	c = config.ImportConfig(config.OSSource{})
	interruptHandler := make(chan os.Signal, 1)
	signal.Notify(interruptHandler, syscall.SIGTERM, syscall.SIGINT)

	addr := fmt.Sprintf(":%s", c.ServicePort)
	go func(addr string) {
		log.Println(fmt.Sprintf("Jacq API service running on %v. Environment=%s", addr, c.AppEnv))
		if err := http.ListenAndServe(addr, router); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}(addr)

	<-interruptHandler
	log.Println("Closing application...")
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal("Failed to disconnect from database")
	}
}
