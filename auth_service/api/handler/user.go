package handler

import (
	"log"

	pb "finance_tracker/auth_service/genproto"
	token "finance_tracker/auth_service/token"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  user  body  pb.UserCreateReq  true  "User data"
// @Success 200 {object} string "User registered successfully"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /user/register [post]
func (h *Handler) RegisterUser(c *gin.Context) {
	var req pb.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// _, err := h.stg.Register(c, &req)
	// if err!= nil {
	//     c.JSON(500, gin.H{"error": err.Error()})
	//     return
	// }
	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	err = h.Producer.ProduceMessages("user-create", input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully"})
}

// LoginUser godoc
// @Summary Log in a user
// @Description Log in a user with the provided credentials
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  login  body  pb.LoginReq  true  "Login credentials"
// @Success 200 {object} pb.Token
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string  "Internal server error"
// @Router /user/login [post]
func (h *Handler) LoginUser(c *gin.Context) {
	var req pb.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.stg.Login(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, token)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update user profile with the provided details
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  profile  body  pb.UserUpdateModel  true  "Profile data"
// @Success 200 {object} string  "Profile updated successfully"
// @Failure 400 {object} string "Invalid input"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /user/profile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	claims, err := token.ExtractClaim(tokenStr)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	id := claims["id"]
	var body pb.UserUpdateModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req := pb.UserUpdate{
		Id:   id.(string),
		Body: &body,
	}
	_, err = h.stg.UpdateProfile(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Profile updated successfully"})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change user password with the provided details
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  password  body  pb.PasswordChangeBody  true  "Password change data"
// @Success 200 {object} string "Password updated successfully"
// @Failure 400 {object} string "Invalid input"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal server error"
// @Router /user/password [put]
func (h *Handler) ChangePassword(c *gin.Context) {
	tokenStr := c.Request.Header.Get("Authorization")
	claims, err := token.ExtractClaim(tokenStr)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	id := claims["id"]
	var body pb.PasswordChangeBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req := pb.PasswordChangeReq{
		UserId: id.(string),
		Body:   &body,
	}
	_, err = h.stg.ChangePassword(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password updated successfully"})
}

// ForgotPassword godoc
// @Summary Send a reset password code to the user's email
// @Description Send a reset password code to the user's email
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  email  body  pb.EmailBody  true  "Email data"
// @Success 200 {object} string "Reset password code sent successfully"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /user/forgot_password [post]
func (h *Handler) ForgotPassword(c *gin.Context) {
	var email pb.EmailBody
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req := pb.ForgotPasswordReq{
		Body: &email,
	}

	// _, err := h.stg.ForgotPassword(c, &req)
	// if err!= nil {
	//     c.JSON(500, gin.H{"error": err.Error()})
	//     return
	// }

	input, err := protojson.Marshal(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	err = h.Producer.ProduceMessages("forgot_password", input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "Reset password code sent successfully"})
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Reset user password with the provided reset code and new password
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  resetCode  body  pb.ResetBody  true  "Reset code data"
// @Success 200 {object} string "Password reset successfully"
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /user/reset_password [post]
func (h *Handler) ResetPassword(c *gin.Context) {
	var resetCode pb.ResetBody
	if err := c.ShouldBindJSON(&resetCode); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req := pb.PasswordResetReq{
		Body: &resetCode,
	}

	_, err := h.stg.ResetPassword(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password reset successfully"})
}

// GetUserProfile godoc
// @Summary GetUserProfile
// @Description Gets user profile
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Success 200 {object} pb.UserCreateRes
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /user/profile [get]
func (h *Handler) GetUserProfile(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	log.Println("token from header:", tokenStr)
	claims, err := token.ExtractClaim(tokenStr)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	id := claims["id"]
	req := pb.ById{
		Id: id.(string),
	}
	userProfile, err := h.stg.GetUserProfile(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userProfile)
}

// ConfirmEmail godoc
// @Summary ConfirmEmail
// @Description Confirm Email
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  Email  body  pb.EmailConfirm  true  "Email Confirm data"
// @Success 200 {object} pb.EmailConfirm
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /user/confirm [post]
func (h *Handler) ConfirmEmail(c *gin.Context) {
	var req pb.EmailConfirm
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.stg.ConfirmEmail(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Email confirmed successfully"})
}

// ResendCode godoc
// @Summary ResendCode
// @Description Resends Code
// @Tags User
// @Accept  json
// @Produce  json
// @Security  		BearerAuth
// @Param  Email  body  pb.ResendReq  true  "Email data"
// @Success 200 {object} pb.ResendReq
// @Failure 400 {object} string "Invalid input"
// @Failure 500 {object} string "Internal server error"
// @Router /user/resend [post]
func (h *Handler) ResendCode(c *gin.Context) {
	var req pb.ResendReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.stg.ResendCode(c, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Code resent successfully"})
}
