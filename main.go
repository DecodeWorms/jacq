package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
	router.POST("/signup", server.SignUp())
	router.POST("/email_verification", server.SendVerificationEmail())
	router.PUT("/update_record", server.UpdateUser())
	router.PUT("/secure_transaction", server.SecureTransaction())
	router.POST("/login", server.Login())
	router.POST("/forgot_password", server.ForgotPassword())
	router.PUT("/change_password", server.ChangePassword())
	router.POST("/verify_phone_number", server.VerifyPhoneNumber())
	router.PUT("/change_pin", server.ChangeTransactionPin())
	router.POST("/verify_token", server.VerifyToken())
	router.POST("/verify_bvn", server.VerifyBvn())

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
