package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"miniCloudStroage/iserver"
	"net/http"
)

type DirRequest struct {
	filePath string
	addPath  string
	delPath  string
}

func main() {
	router := gin.Default()

	// linux下
	// server := iserver.NewSimpleFileServer("/home/sti/tmp")

	// windows下
	server := iserver.NewSimpleFileServer("D:\\cloudStorage")

	router.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_info": "file upload failed"})
			return
		}

		// 打开文件
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_info": "file open error"})
			return
		}

		data, err := io.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_info": "file read error"})
			return
		}

		server.Upload(data, fileHeader.Filename, "/test")
		c.JSON(http.StatusOK, gin.H{"msg": "file uploaded"})
		return
	})

	router.GET("/download", func(c *gin.Context) {
		// 文件路径
		filePath := c.Query("path")

		// 打开文件
		data, err := server.Download(filePath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_info": err.Error()})
			return
		}

		// 设置响应头
		c.Header("Content-Disposition", "attachment; filename=example.txt")
		c.Header("Content-Type", "application/octet-stream")

		// 流式传输文件
		_, err = c.Writer.Write(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error_info": "file read error"})
			return
		}

	})

	router.POST("/list_dirs", func(c *gin.Context) {
		raw, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_info": "read body fail"})
			return
		}

		var dirRequest DirRequest
		if err := json.Unmarshal(raw, &dirRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error_info": "json parse fail"})
			return
		}

		dirs := server.ListDirs(dirRequest.filePath)

		c.JSON(http.StatusOK, gin.H{"data": dirs})
		return
	})

	router.POST("/add_dir", func(c *gin.Context) {
		raw, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_info": "read body error",
			})
			return
		}

		var dirRequest DirRequest
		if err := json.Unmarshal(raw, &dirRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_info": "json parse fail",
			})
			return
		}

		err = server.AddDir(dirRequest.addPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_info": "add dir error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
		return
	})

	router.POST("/del_dir", func(c *gin.Context) {
		raw, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_info": "read body error",
			})
			return
		}

		var dirRequest DirRequest
		if err := json.Unmarshal(raw, &dirRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_info": "json parse fail",
			})
			return
		}

		err = server.DelDir(dirRequest.delPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_info": "add dir error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
		return
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
