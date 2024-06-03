package server

import (
	"github.com/gin-gonic/gin"
	"jacq/handler"
	"jacq/model"
	"net/http"
)

type UserServer struct {
	user *handler.UserHandler
}

func NewUserServer(user *handler.UserHandler) UserServer {
	return UserServer{
		user: user,
	}
}

func (user UserServer) SignUp() gin.HandlerFunc {
	return func(context *gin.Context) {
		var u *model.User

		if err := context.ShouldBindJSON(&u); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 400, "status": "Bad request"})
			return
		}

		//Call the handler to process the request
		if err := user.user.CreateUser(u); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": 500, "status": "Internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "user signed up successfully", "code": 200})
	}
}

func (user UserServer) SendVerificationEmail() gin.HandlerFunc {
	return func(context *gin.Context) {
		var ver model.VerifyEmail

		if err := context.ShouldBindJSON(&ver); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 400, "status": "Bad request"})
			return
		}

		//Call the handler to process the request
		if err := user.user.SendVerificationLink(ver.Email, ver.Link); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": 500, "status": "Internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "email sent out successfully", "code": 200})
	}
}

func (user UserServer) UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var us model.User
		id := context.Query("id")

		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 400, "status": "Internal server error"})
			return
		}

		//Call the handler to process the request
		if err := user.user.UpdateUser(id, &us); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": 500, "status": "Internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "user's record updated successfully", "code": 200})
	}
}

func (user UserServer) SecureTransaction() gin.HandlerFunc {
	return func(context *gin.Context) {
		var us model.User
		id := context.Query("id")

		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 400, "status": "Bad request"})
			return
		}

		//Call the handler to process the request
		if err := user.user.SecureTransaction(id, &us); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": 500, "status": "Internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "user's transaction code secured successfully", "code": 200})
	}
}

func (user UserServer) Login() gin.HandlerFunc {
	return func(context *gin.Context) {

		var us model.User
		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 400, "status": "Bad request"})
			return
		}

		//Call the handler to process the request
		resp, err := user.user.Login(&us)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": 500, "status": "Internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "user's successful logged in", "accessToken": resp, "code": 200})
	}
}

func (user UserServer) ForgotPassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		var us model.User

		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 400, "status": "Bad request"})
			return
		}
		if err := user.user.ForgotPassword(&us); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": 500, "status": "Internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"code": 200, "status": "link is sent successfully"})
	}
}

func (user UserServer) ChangePassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		userID := context.Query("id")
		var chanPass model.ChangePassword

		if err := context.ShouldBindJSON(&chanPass); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": 400, "status": "Bad request"})
			return
		}

		//Call handler to process the request
		if err := user.user.ChangePassword(userID, &chanPass); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "code": 500, "status": "Internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "user's password changed successfully", "code": 200})
	}
}

func (user UserServer) VerifyPhoneNumber() gin.HandlerFunc {
	return func(context *gin.Context) {
		phone := context.Query("phone_number")
		id := context.Query("id")

		//Call handler to process the request
		if err := user.user.VerifyNumber(id, phone); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err, "code": 500, "status": "internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"code": 200, "status": "code was successfully sent to a user"})
		return
	}
}

func (user UserServer) ChangeTransactionPin() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Query("id")
		var tranPin model.TransactionPin

		if err := context.ShouldBindJSON(&tranPin); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"code": 400, "status": "Bad request"})
			return
		}

		//Call handler to process the request
		if err := user.user.ChangeTransactionPin(id, &tranPin); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err, "code": 500, "status": "internal server error"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"code": 200, "status": "code was successfully sent to a user"})
		return
	}

}
