package main

import (
	//provides HTTP clientand server implementation
	"net/http"
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
	// os.Setenv("GROUPNAME", "Beryllium")
	content := "{ \n  \"status\" : \"Success\", \n  \"data\": {\n \t  \"groupname\":\"" + os.Getenv("GROUPNAME") + "\"\n  } \n}"
	// content2 := gin.H{"status" : "Success", "groupname" : os.Getenv("DEVO_GROUPNAME")}
	c.String(http.StatusOK, content)
	// c.JSON(2000, content2)
}