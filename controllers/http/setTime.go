package http

import (
	"collections/helper"
	"collections/helper/timetn"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

// SetTimeController operations for Set Time
type SetTimeController struct {
	beego.Controller
}

// SetTime ...
func (c *SetTimeController) SetTime() {
	body := c.Ctx.Input.RequestBody
	var t map[string]int
	err := json.Unmarshal(body, &t)
	helper.CheckErr("Failed Unmarshal line 20 controller/setTime", err)

	timetn.GetSetTimeDate(t["set_time"])

	c.Data["json"] = timetn.Now().String()

	c.ServeJSON()
}

// SetSpesificTime ...
func (c *SetTimeController) SetSpesificTime() {
	body := c.Ctx.Input.RequestBody
	var t map[string]string
	err := json.Unmarshal(body, &t)
	helper.CheckErr("Failed Unmarshal line 20 controller/setTime", err)

	timeSp := strings.Split(t["set_time"], "-")
	y, err := strconv.Atoi(timeSp[0])
	helper.CheckErr("", err)
	m, err := strconv.Atoi(timeSp[1])
	helper.CheckErr("", err)
	d, err := strconv.Atoi(timeSp[2])
	helper.CheckErr("", err)

	timetn.SetTimeDateWithSpesificDate(y, time.Month(m), d)

	c.Data["json"] = timetn.Now().String()

	c.ServeJSON()
}
