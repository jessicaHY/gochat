package ajax

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"chatroom/service/models"
	"chatroom/helper"
	"chatroom/utils/Constants"
	"log"
	"net/http"
)

func ListGift(req *http.Request, rend render.Render) {
	group := Constants.GetGroupFromReq(req)
	gt := Constants.GroupType(group)
	rss, err := models.ListGift(gt)
	if err != nil {
		log.Println(err)
		rend.JSON(200, helper.Error(helper.ParamsError))
		return
	}
	rend.JSON(200, helper.Success(rss))
}

func ListDonateByRoom(params martini.Params, rend render.Render) {
	roomId := helper.Int64(params["roomId"])
	log.Println(roomId)
	if roomId <= 0 {
		rend.JSON(200, helper.Error(helper.ParamsError))
		return
	}
	rss, err := models.ListDonateByRoom(roomId)
	if err != nil {
		log.Println(err)
		rend.JSON(200, helper.Error(helper.ParamsError))
		return
	}
	rend.JSON(200, helper.Success(rss))
}
