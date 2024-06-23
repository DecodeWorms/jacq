package server

import (
	"github.com/gin-gonic/gin"
	_ "jacq/docs"
	"jacq/handler"
	"jacq/model"
	"net/http"
	"strings"
)

// @Jacq User Service Api Documentation
// @version 1.0
// @description API documentation for User Service
// @host localhost:8001
// @BasePath /

type UserServer struct {
	user *handler.UserHandler
}

func NewUserServer(user *handler.UserHandler) UserServer {
	return UserServer{
		user: user,
	}
}

// SignUp User
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Produce  json
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/signup [Post]
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

// SendVerificationEmail Send Verification email
// @Summary Send Verification email
// @Description Send Verification email
// @Tags user
// @Produce  json
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/email_verification [Post]
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

// UpdateUser Updates user's existing record
// @Summary Updates user's existing record
// @Description Updates user's existing record
// @Tags user
// @Produce  json
// @Param model body model.User true "User request data"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/update_record [Put]
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

// SecureTransaction  Secures user`s transaction pin
// @Summary Secures user`s transaction pin
// @Description Secures user`s transaction pin
// @Tags user
// @Produce  json
// @Param model body model.User true "User request data"
// @Param id query string true "ID"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/secure_transaction [Post]
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

// Login  Logins user
// @Summary Login user
// @Description Login user
// @Tags user
// @Produce  json
// @Param model body model.User true "User request data"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/login [Post]
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

// ForgotPassword  sends user forgot password 6 digits code
// @Summary send user forgot password 6 digits code
// @Description sends user forgot password 6 digits code
// @Tags user
// @Produce  json
// @Param model body model.User true "User request data"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/forgot_password [Put]
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
		context.JSON(http.StatusOK, gin.H{"data": handleServerResponse(http.StatusOK, "User forgot password requested successfully", "", nil, nil)})
	}
}

// ChangePassword changes user's existing password
// @Summary changes user's existing password
// @Description changes user's existing password
// @Tags user
// @Produce  json
// @Param model body model.ChangePassword true "User request data"
// @Param id query string true "ID"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/change_password [Put]
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

// VerifyPhoneNumber verifies user's number
// @Summary verify user's number
// @Description verify user's number
// @Tags user
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/verify_phone_number [Post]
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

// ChangeTransactionPin changes user's transaction pin
// @Summary change user's transaction pin
// @Description change user's transaction pin
// @Tags user
// @Produce  json
// @Param model body model.TransactionPin true "User request data"
// @Param id query string true "ID"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/change_pin [Put]
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

// VerifyToken verifies user's token
// @Summary verify user's token
// @Description verify user's token
// @Tags user
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/verify_token [Post]
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

// VerifyBvn verifies user's bvn
// @Summary verify user's bvn
// @Description verify user's bvn
// @Tags user
// @Produce  json
// @Param id query string true "ID"
// @Param model body model.User true "User request data"
// @Success 200 {object} model.ServerResponse{code=int,status=string,error=interface{},token=string}
// @Router /user/verify_bvn [Post]
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
