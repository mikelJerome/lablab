package handler

import (
	"github.com/gin-gonic/gin"
	redislock "github.com/jefferyjob/go-redislock"
	"go.uber.org/zap"
	"lab_sys/database"
	"lab_sys/global"
	"lab_sys/model"
	"lab_sys/response"
	"log"
	"net/http"
	"time"
)

//根据条件查询预约

func Reserve(c *gin.Context) {
	// 获取用户名字
	username := c.PostForm("username")
	// 获取器材名字
	equipmentName := c.PostForm("equipmentName")
	// 开始时间
	startTime := c.PostForm("startTime")
	// 结束时间
	overTime := c.PostForm("overTime")

	// 解析开始时间模板
	startTimeStr, err := time.Parse("2006-01-02-15-04", startTime)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "时间格式不准确"})
		return
	}
	// 解析结束时间模板
	overTimeStr, err := time.Parse("2006-01-02-15-04", overTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid over time"})
		return
	}
	// 如果我们数据库没有名字也是不可以
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
	if userstring, _ := database.CheckForReservationConflict(startTimeStr, overTimeStr); err != nil {
		response.Err(c, http.StatusBadRequest, 500, "预约时间冲突", userstring)
		return
	}
	createdReservation, err := database.CreateReservation(reserve)
	if err != nil {
		response.Err(c, http.StatusBadRequest, 500, "创建用户失败", nil)
		return
	}
	if global.Redis == nil {
		log.Println("global.Redis is nil. Redis client has not been initialized.")
	} else {
		log.Println("global.Redis is initialized.")
	}

	// 初始化 lockerClient
	lockerClient := redislock.New(c, global.Redis, user.Username+" "+equipmentName+" "+startTime+" "+overTime, redislock.WithTimeout(time.Duration(10)*time.Second))

	// 检查 lockerClient 是否已经被正确初始化
	if lockerClient == nil {
		log.Println("lockerClient is nil. Redis lock has not been initialized properly.")
	} else {
		log.Println("lockerClient is initialized.")
	}
	//lockerClient := redislock.New(c, global.Redis, time.Now().String(), redislock.WithTimeout(time.Duration(30)*time.Second))

	_ = lockerClient.Lock()
	defer lockerClient.UnLock()
	// 返回预约成功消息

	c.JSON(http.StatusOK, gin.H{"message": "用户预约成功", "reservation": createdReservation})
}
