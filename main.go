package main

import (
    "github.com/gin-gonic/gin"
    "html/template"
    "log"
    "net/http"
    "codeshare/models"
)

func main() {
    // 初始化数据库
    if err := models.InitDB(); err != nil {
        log.Fatal("数据库初始化失败:", err)
    }

    // 初始化数据
    if err := models.InitializeData(); err != nil {
        log.Printf("警告：数据初始化失败: %v", err)
    }

    r := gin.Default()
    
    // 添加模板函数
    r.SetFuncMap(template.FuncMap{
        "unescapeHTML": func(s string) template.HTML {
            return template.HTML(s)
        },
    })
    
    // 加载HTML模板
    r.LoadHTMLGlob("templates/*")
    
    // 设置静态文件目录
    r.Static("/static", "./static")
    
    // 首页 - 显示所有页面列表
    r.GET("/", func(c *gin.Context) {
        pages, err := models.GetAllPages()
        if err != nil {
            c.HTML(http.StatusInternalServerError, "404.html", gin.H{
                "error": "获取页面列表失败",
            })
            return
        }
        
        c.HTML(http.StatusOK, "index.html", gin.H{
            "title": "旅游攻略分享平台",
            "pages": pages,
        })
    })
    
    // 动态页面路由
    r.GET("/page/*path", func(c *gin.Context) {
        path := c.Param("path")
        page, err := models.GetPageByPath(path)
        if err != nil {
            c.HTML(http.StatusInternalServerError, "404.html", gin.H{
                "error": "获取页面失败",
            })
            return
        }
        
        if page == nil {
            c.HTML(http.StatusNotFound, "404.html", gin.H{
                "error": "页面不存在",
            })
            return
        }
        
        c.HTML(http.StatusOK, "view.html", gin.H{
            "page": page,
        })
    })
    
    // 管理接口 - 创建或更新页面
    r.POST("/api/pages", func(c *gin.Context) {
        var page models.Page
        if err := c.BindJSON(&page); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
            return
        }
        
        if err := page.Save(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "保存页面失败"})
            return
        }
        
        c.JSON(http.StatusOK, page)
    })
    
    // 启动服务器
    r.Run(":8914")
} 