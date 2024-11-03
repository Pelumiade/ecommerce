package controllers

import (
	"context"
	"fmt"
	"go/token"
	"net/http"

	"github.com/Pelumiade/ecommerce/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/internal/genid"
)

func Hashpassword(password string) string {

}

func Verifypassword(password string, givenPassword string) (bool, string) {

}

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return

		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.H{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServiceError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.H{"phone": user.Phone})

		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServiceError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone no is already in use"})
			return
		}

		password = HashPassword(*user.Password)
		user.Password = &password

		user.Created_At = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.user_ID = user.ID.HEX()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_name, *user.Last_name, user.user_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.Insertone(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServiceError, gin.H{"error": "user did not get created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Succesfully signed in!")

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := UserCollection.FindOne(ctx, bson.H{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}
		PasswordIsvalid, msg := VerifyPassword(*user.Password, *founduser.Password)

		defer cancel()

		if !PasswordIsvalid {
			c.JSON{http.StatusInternalServerError, gin.H{"error": msg}}
			fmt.println(msg)
			return
		}
		token, refreshToken, _ := generate.TokenGenerator(*founderuser.Email, *founduser.First_name, *founderuser.Last_name, *founderuser.User_ID)
		defer cancel()

		generate.UpdateAllTokens(token, refreshToken, founderuser.user_ID)
		c.JSON(http.StatusFound, founduser)

	}
}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductQuery() gin.HandlerFunc {

}
