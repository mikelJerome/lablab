package handler

import (
	"github.com/gin-gonic/gin"
	"lab_sys/lab_sys/database"
	"lab_sys/lab_sys/model"
	"net/http"
	"time"
)

//根据条件查询预约

func Reserve(c *gin.Context) {
	// 获取表单请求
	// 获取用户名字
	username := c.PostForm("username")
	// 获取器材名字
	equipmentName := c.PostForm("equipmentName")
	// 开始时间
	startTime := c.PostForm("startTime")
	// 结束时间
	overTime := c.PostForm("overTime")

	// 解析开始时间模板
	startTimeStr, err := time.Parse("2006-01-02 15:04", startTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start time"})
		return
	}
	// 解析结束时间模板
	overTimeStr, err := time.Parse("2006-01-02 15:04", overTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid over time"})
		return
	}
	// 如果我们数据库没有名字也是不可以的
	//_, err = database.FindUserByName(username)
	//if err != nil {
	//	c.JSON(200, gin.H{"errors": "请确保你的名字已经存在数据库"})
	//	return
	//}
	// 获取用户id，根据我们用户名字返回用户id
	user, err := database.FindUserByName(username)
	if err != nil {
		c.JSON(200, gin.H{"errors": "用户名不正确"})
		return
	}

	// 获取实验器材id
	equipment, err := database.FindEquipmentByName(equipmentName)
	if err != nil {
		c.JSON(200, gin.H{"errors": "没有这个实验器材"})
		return
	}
	reserve := model.Reservation{
		UserID:      user.ID,
		EquipmentID: equipment.ID,
		StartTime:   startTimeStr,
		OverTime:    overTimeStr,
	}
	createdReservation, err := database.CreateReservation(reserve)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册用户失败"})
		return
	}

	// 返回预约成功消息
	c.JSON(http.StatusOK, gin.H{"message": "用户预约成功", "reservation": createdReservation})
}
