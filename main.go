package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

var c config.Config

func main() {
	c = config.ImportConfig(config.OSSource{})

	// Establish connection to Mongodb service
	repo, client, err := storage.New(c.DatabaseURL, c.DatabaseName)
	if err != nil {
		log.Printf("Error failed to open mongoDB %v", err)
	}

	//Call user handler
	handler := handler2.NewUserHandler(repo, encrypt.NewPasswordEncryptor(), idgenerator.New())

	//Call server
	server := server2.NewUserServer(&handler)

	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	//Graceful shut down
	interruptHandler := make(chan os.Signal, 1)
	signal.Notify(interruptHandler, syscall.SIGTERM, syscall.SIGINT)

	addr := fmt.Sprintf(":%s", c.ServicePort)
	go func(addr string) {
		log.Println(fmt.Sprintf("Jacq API service running on %v. Environment=%s", addr, c.AppEnv))
		if err := http.ListenAndServe(addr, router); err != nil {
			log.Printf("error starting a server %v", err)
		}
	}(addr)

	<-interruptHandler
	log.Println("closing application...")
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal("failed to disconnect from database")
	}

}

/*Configurations for CORS
confg := cors.DefaultConfig()
confg.AllowAllOrigins = true
confg.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Access-Control-Allow-Origin", "XMLHttpRequest", "*"}
confg.ExposeHeaders = []string{"Content-Length"}
confg.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}
confg.AllowCredentials = true

router := gin.New()
router.Use(cors.New(confg))
router.Use(gin.Logger())
router.Use(gin.Recovery())
*/
