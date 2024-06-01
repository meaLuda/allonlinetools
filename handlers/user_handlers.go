package handlers

import (
	"allonlinetools/models"
	"allonlinetools/sessionstore"
	"allonlinetools/utils"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type UserController struct {
	collection  *mongo.Collection
	ctx         context.Context
}

func NewUserController(collection *mongo.Collection, ctx context.Context) *UserController {
	return &UserController{
		collection:  collection,
		ctx:         ctx,
	}
}

func (uc *UserController) Handler_RenderSignUpPage(c *fiber.Ctx) error {
	fmt.Println(c.Locals("csrf_token"))
	// // Get CSRF token from session
	// Get session from storage
	sess, err := sessionstore.Store.Get(c)
	if err != nil {
		panic(err)
	}
	fmt.Println("----------- Middleware Save Session token -------------------")
	sess.Set("csrf_token", c.Locals("csrf_token"))
	sess.Save()
	return c.Render("home/pages/signup",  fiber.Map{
		"CsrfToken": c.Locals("csrf_token"),
	})
}


func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	// Parse form data
	signupReq := new(models.UserSignUp)
	if err := c.BodyParser(signupReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}

	// RETRIEVE AND PRINT ALL USERS
	// RETRIEVE AND PRINT ALL USERS
	var users []models.User
	cursor, err := uc.collection.Find(uc.ctx, bson.M{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error retrieving users")
	}
	if err := cursor.All(uc.ctx, &users); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error decoding users")
	}
	for _, user := range users {
		fmt.Println(user)
	}

	//// Process form data and create user
	hashedPassword, err := utils.HashPassword(signupReq.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Server Error")
	}
	
	user := new(models.User)
	user.ID = primitive.NewObjectID()
	user.Created_at = time.Now()
	user.UserName = signupReq.UserName
	user.Email = signupReq.Email
	user.Phone = signupReq.Phone
	user.Password = hashedPassword
	// fmt.Println(user)

	resp, err := uc.collection.InsertOne(uc.ctx, user)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create user")
	}
	fmt.Println(resp)
	return c.Redirect("/auth/login")
}


func (uc *UserController) Handler_RenderLoginPage(c *fiber.Ctx) error {
	fmt.Println(c.Locals("csrf_token"))

	// Get session from storage
	sess, err := sessionstore.Store.Get(c)
	if err != nil {
		panic(err)
	}
	fmt.Println("----------- Middleware Save Session token -------------------")
	sess.Set("csrf_token", c.Locals("csrf_token"))
	sess.Save()
	return c.Render("home/pages/login", fiber.Map{
		"CsrfToken": c.Locals("csrf_token"),
	})
}

func (uc *UserController) LoginUser(c *fiber.Ctx) error {
	// Parse form data
	loginReq := new(models.UserLogin)
	if err := c.BodyParser(loginReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
	}

	// Find user by email
	var user models.User
	filter := bson.M{"email": loginReq.Email}
	err := uc.collection.FindOne(uc.ctx, filter).Decode(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Check if the password is correct
	err = utils.VerifyPassword(loginReq.Password, user.Password) 
	if err != nil{
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}


	// Create session and set user ID
	sess, err := sessionstore.Store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Server Error")
	}
	sess.Set("user_id", user.ID.Hex())
	sess.Save()
	// redirect to main page
	return c.Redirect("/")
}
