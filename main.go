package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    
    // 加载HTML模板
    r.LoadHTMLGlob("templates/*")
    
    // 设置静态文件目录
    r.Static("/static", "./static")
    
    // 首页
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "旅游攻略分享平台",
        })
    })
    
    // 西安攻略页面
    r.GET("/xianthree", func(c *gin.Context) {
        c.HTML(http.StatusOK, "xianthree.html", nil)
    })
    
    // 山东攻略页面
    r.GET("/shandong", func(c *gin.Context) {
        c.HTML(http.StatusOK, "shandonglaodong.html", nil)
    })
    
    // 启动服务器
    r.Run(":8914")
} 