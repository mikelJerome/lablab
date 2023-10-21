package test

import (
	"fmt"
	"lab_sys/lab_sys/utils"
	"testing"
)

func Test(t *testing.T) {
	s := utils.RandNumber(10)
	fmt.Print(s)
}

//func Test1(t *testing.T) {
//	s := utils.NewSet()
//	s.Add("nich")
//}
