package users

import (
	"github.com/Ameen9044/bookstore_users-api/domain/users"
	"github.com/Ameen9044/bookstore_users-api/services"
	"github.com/Ameen9044/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*Every Request is going to be handled by controller*/
/*This is the entry point of an application*/
/*Verify the request and then send to service*/

func CreateUser(c *gin.Context)  {
	/*take the json request(handle json request) and populate the User data Start*/
	var user users.User
	/*from request and populate First Method (byte array) Start*/
	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		//TODO Handle Error
		return
	}
	if err := json.Unmarshal(bytes, &user); err !=nil{
		//TODO: Handle json error
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil{
		//TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)*/
	/*from request and populate First Method (byte array) END*/

	/*from request and populate Second Method (above work in one line) Start*/
	if err := c.ShouldBindJSON(&user); err !=nil{
		restErr := errors.NewBadRequest("Invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil{
		//TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
	/*from request and populate Second Method (above work in one line) End*/

	/*take the json from request and populate the User data END*/
}

func GetUser(c *gin.Context)  {
	c.String(http.StatusOK,"POST Users")
}




