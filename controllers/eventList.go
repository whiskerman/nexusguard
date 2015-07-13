package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"nexusguard/models"
)

// Operations about object
type EventListController struct {
	beego.Controller
}

// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *EventListController) Post() {
	/*
		var ob models.Object
		json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
		objectid := models.AddOne(ob)
		o.Data["json"] = map[string]string{"ObjectId": objectid}
		o.ServeJson()
	*/
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *EventListController) Get() {
	spe_name := o.Ctx.Input.Params[":spe_name"]
	sub_verify := o.Ctx.Input.Params[":sub_verify"]
	etag := o.Ctx.Input.Params[":etag"]
	version := o.Ctx.Input.Params[":version"]
	beego.BeeLogger.Debug("spe_name:%s & sub_verify:%s & etag:%s &version:%s \n %v", spe_name, sub_verify, etag, version, o.Ctx.Input.Params)
	o.Data["json"] = `{"errno":0,"errmsg":"success","result":{"status":1,"time":1436254017},"etag":"69a4d8b44c79bcde396f3cab4d56f554"}`
	/*
		if objectId != "" {
			ob, err := models.GetOne(objectId)
			if err != nil {
				o.Data["json"] = err
			} else {
				o.Data["json"] = ob
			}
		}
	*/
	o.ServeJson()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *EventListController) GetAll() {
	/*
		obs := models.GetAll()
		o.Data["json"] = obs
		o.ServeJson()
	*/
	spe_name := o.Ctx.Input.Query("spe_name")
	sub_verify := o.Ctx.Input.Query("_sub_verify")
	etag := o.Ctx.Input.Query("etag")
	version := o.Ctx.Input.Query("version")
	//o.Ctx.Input.Query(key)
	beego.BeeLogger.Debug("spe_name:%s & sub_verify:%s & etag:%s &version:%s \n %v", spe_name, sub_verify, etag, version, o.Ctx.Input)
	o.Data["json"] = `{"errno":0,"errmsg":"success","result":{"status":1,"time":1436254017},"etag":"69a4d8b44c79bcde396f3cab4d56f554"}`
	/*
		if objectId != "" {
			ob, err := models.GetOne(objectId)
			if err != nil {
				o.Data["json"] = err
			} else {
				o.Data["json"] = ob
			}
		}
	*/
	o.ServeJson()
}

// @Title update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *EventListController) Put() {
	objectId := o.Ctx.Input.Params[":objectId"]
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJson()
}

// @Title delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *EventListController) Delete() {
	objectId := o.Ctx.Input.Params[":objectId"]
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJson()
}
