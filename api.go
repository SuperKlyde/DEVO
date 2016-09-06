package main

import (
	"os"
	//"syscall"
	//"fmt"
	//micro framework routing
	"github.com/gin-gonic/gin"
	//package for the conversion from and to string
	//"strconv"
)


var groupName = "Beryllium"

func main(){
	router := gin.Default()
	router.GET("/DEVO/GROUPNAME", getGroupName)
	router.Run(":8080")
}

func getGroupName(c *gin.Context){
	os.Setenv("DEVO_GROUPNAME", "Beryllium")
	// syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
	c.JSON(2000, groupName)
}