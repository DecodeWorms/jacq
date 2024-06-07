package server

import (
	"github.com/gin-gonic/gin"
	"jacq/handler"
	"jacq/model"
	"net/http"
	"strings"
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
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "Invalid client request", "", err.Error(), nil)})
			return
		}

		//Call the handler to process the request
		if err := user.user.CreateUser(u); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal Server error", "", err.Error(), nil)})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User Created Successfully", "", nil, nil)})
	}
}

func (user UserServer) SendVerificationEmail() gin.HandlerFunc {
	return func(context *gin.Context) {
		var ver model.VerifyEmail

		if err := context.ShouldBindJSON(&ver); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid request", "", err.Error(), nil)})
			return
		}

		//Call the handler to process the request
		if err := user.user.SendVerificationLink(ver.Email, ver.Link); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server error", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "Verification link sent successfully", "", nil, nil)})
	}
}

func (user UserServer) UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		var us model.User
		id := context.Query("id")

		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid request", "", err.Error(), nil)})
			return
		}

		//Call the handler to process the request
		if err := user.user.UpdateUser(id, &us); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend internal server error", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User's record updated successfully", "", nil, nil)})
	}
}

func (user UserServer) SecureTransaction() gin.HandlerFunc {
	return func(context *gin.Context) {
		var us model.User
		id := context.Query("id")

		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid Request", "", err.Error(), nil)})
			return
		}

		//Call the handler to process the request
		if err := user.user.SecureTransaction(id, &us); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server error", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User Transaction created successfully", "", nil, nil)})
	}
}

func (user UserServer) Login() gin.HandlerFunc {
	return func(context *gin.Context) {

		var us model.User
		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid Request", "", err.Error(), nil)})
			return
		}

		//Call the handler to process the request
		token, err := user.user.Login(&us)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server error", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User has successfully logged in", token, nil, nil)})
	}
}

func (user UserServer) ForgotPassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		var us model.User

		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid request", "", err.Error(), nil)})
			return
		}
		if err := user.user.ForgotPassword(&us); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backed Internal server error", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User forgot password requested succesfully", "", nil, nil)})
	}
}

func (user UserServer) ChangePassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		userID := context.Query("id")
		var chanPass model.ChangePassword

		if err := context.ShouldBindJSON(&chanPass); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid request", "", err.Error(), nil)})
			return
		}

		//Call handler to process the request
		if err := user.user.ChangePassword(userID, &chanPass); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server error", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User Password changed successfully", "", nil, nil)})
	}
}

func (user UserServer) VerifyPhoneNumber() gin.HandlerFunc {
	return func(context *gin.Context) {
		phone := context.Query("phone_number")
		id := context.Query("id")

		//Call handler to process the request
		if err := user.user.VerifyNumber(id, phone); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server error", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User number verified successfully", "", nil, nil)})
	}
}

func (user UserServer) ChangeTransactionPin() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Query("id")
		var tranPin model.TransactionPin

		if err := context.ShouldBindJSON(&tranPin); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid request", "", err.Error(), nil)})
			return
		}

		//Call handler to process the request
		if err := user.user.ChangeTransactionPin(id, &tranPin); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server erro", "", err.Error(), nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User transaction Pin changed successfully", "", nil, nil)})
	}
}

func (user UserServer) VerifyToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Query("id")

		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"data": handleServerResponse(http.StatusUnauthorized, "Authorization header is required", "", nil, nil)})
			context.Abort()
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			context.JSON(http.StatusUnauthorized, gin.H{"data": handleServerResponse(http.StatusUnauthorized, "Authorization header format must be Bearer {token}", "", nil, nil)})
			context.Abort()
			return
		}

		//Call handler to process the request
		if err := user.user.VerifyOtp(id, tokenString); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server error", "", nil, nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"code": handleServerResponse(http.StatusOK, "OTP verified successfully", "", nil, nil)})
	}
}

func (user UserServer) VerifyBvn() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Query("id")
		var us model.User

		if err := context.ShouldBindJSON(&us); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"data": handleServerResponse(http.StatusBadRequest, "User Invalid request", "", nil, nil)})
			return
		}

		//Call handler to process the request
		if err := user.user.VerifyBvn(id, &us); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"data": handleServerResponse(http.StatusInternalServerError, "Backend Internal server error", "", nil, nil)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User BVN verified successfully", "", nil, nil)})
	}
}

func handleServerResponse(code int, status, token string, error any, object *model.User) model.ServerResponse {
	return model.ServerResponse{
		Code:   code,
		Status: status,
		Object: object,
		Error:  error,
		Token:  token,
	}
}
