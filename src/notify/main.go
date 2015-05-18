package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {

        router := gin.Default()
        router.GET("/room/", roomGET)
        router.POST("/room/", roomPOST)
        router.Run(":8080")
}

func roomGET(c *gin.Context) {
        c.JSON(200, gin.H{
                "status": "success",
                "message": "post param body e.g :{'mail':true, 'sms':true, 'subject':'test','content':'test','recieve':'recieve'}",
        })
}

func roomPOST(c *gin.Context) {
        var param NotifyType
        if err := c.ParseBody(&param); err != nil {
            fmt.Println(err.Error())
        } else {
            DoNotify(param)
        }
        c.JSON(200, gin.H{
                "status": "success",
                "message": "message",
        })
}
