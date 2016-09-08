package main

import (
	//variables d'environnement
	"os" 	
	//micro framework routing
	"github.com/gin-gonic/gin"
)




func main(){
	router := gin.Default()
	router.GET("/v1/groupname", getGroupName)
	router.Run(":8080")
}


//renvoie le nom du groupe passé en variable d'environnement
func getGroupName(c *gin.Context){
	
	c.JSON(2000, os.Getenv("DEVO_GROUPNAME"))
}
