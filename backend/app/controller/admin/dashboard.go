package admin

import (
	"fmt"
	"rustdesk-api-server-pro/app/model"

	"github.com/golang-module/carbon/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type DashboardController struct {
	basicController
}

func (c *DashboardController) GetDashboardStat() mvc.Result {

	userCount, err := c.Db.Count(&model.User{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	peerCount, err := c.Db.Count(&model.Peer{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	onlineCount, err := c.Db.Where("is_online = 1").Count(&model.Peer{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	visitsCount, err := c.Db.Count(&model.Audit{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(iris.Map{
		"userCount":   userCount,
		"peerCount":   peerCount,
		"onlineCount": onlineCount,
		"visitsCount": visitsCount,
	}, "ok")
}

func (c *DashboardController) GetDashboardLineCharts() mvc.Result {
	now := carbon.Now()
	startOfWeek := now.SetWeekStartsAt(carbon.Monday).StartOfWeek()
	endOfWeek := now.SetWeekStartsAt(carbon.Monday).EndOfWeek()

	startOfWeekString := startOfWeek.ToDateTimeString()
	endOfWeekString := endOfWeek.ToDateTimeString()

	type lineChartsData struct {
		Value int    `json:"value"`
		Date  string `json:"date"`
	}

	var userList = make([]lineChartsData, 0)
	err := c.Db.Table(&model.User{}).Select("count(*) as `value`, date(created_at) as `date`").Where("created_at between ? and ?", startOfWeekString, endOfWeekString).GroupBy("`date`").Find(&userList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	var userData = make(map[string]int)
	for _, r := range userList {
		userData[r.Date] = r.Value
	}

	var peerList = make([]lineChartsData, 0)
	err = c.Db.Table(&model.Peer{}).Select("count(*) as `value`, date(created_at) as `date`").Where("created_at between ? and ?", startOfWeekString, endOfWeekString).GroupBy("`date`").Find(&peerList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	var peerData = make(map[string]int)
	for _, r := range peerList {
		peerData[r.Date] = r.Value
	}

	var xDateLine []string
	var seriesUser []int
	var seriesPeer []int
	for i := startOfWeek.Day(); i <= endOfWeek.Day(); i++ {
		current := carbon.Parse(fmt.Sprintf("%d-%d-%d", startOfWeek.Year(), startOfWeek.Month(), i))
		s := current.ToDateString()
		xDateLine = append(xDateLine, s)
		seriesUser = append(seriesUser, userData[s])
		seriesPeer = append(seriesPeer, peerData[s])
	}

	return c.Success(iris.Map{
		"xAxis": xDateLine,
		"users": seriesUser,
		"peer":  seriesPeer,
	}, "ok")
}

func (c *DashboardController) GetDashboardPieCharts() mvc.Result {
	type pieChartsData struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	pieMap := make([]pieChartsData, 0)
	err := c.Db.Table(&model.Peer{}).GroupBy("platform").Select("case when `platform` = '' then 'unknown' else `platform` end as `name`, count(*) as `value`").Find(&pieMap)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(pieMap, "ok")
}
