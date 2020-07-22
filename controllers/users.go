package controllers

import (
	"fmt"
	"github.com/frootlo/dbhelper"
	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {
	if err := dbhelper.Client.Ping(); err != nil {
		fmt.Println("Error in ping: ", err.Error())
	}

	//mobileId := ctx.Request.Header.Get(constants.MobileId)
	//goctx := context.WithValue(ctx.Request.Context(), constants.MobileId, mobileId)
	//fmt.Printf("Context: %v \n", goctx)

	//var user models.User
	//body, err := ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//	fmt.Println("Error in read body. Error: ", err.Error())
	//}
	//
	//if err := json.Unmarshal(body, &user); err != nil {
	//	fmt.Println("Error in unmarshal user, error: ", err.Error())
	//}
	//
	//fmt.Printf(user.FirstName)

}
