package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/planesticud/eventos_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// operations for TrEvento
type TrEventoController struct {
	beego.Controller
}

func (c *TrEventoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAllByPersona", c.GetAllByPersona)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Put", c.Put)
}

// GetAllByPersona ...
// @Title Get All By Persona
// @Description get TrEventoController
// @Param	id		path 	string	true		"Persona"
// @Success 200 {object} models.TrEventoController
// @Failure 404 not found resource
// @router /:persona [get]
func (c *TrEventoController) GetAllByPersona() {
	idPersonaStr := c.Ctx.Input.Param(":persona")
	idPersona, _ := strconv.Atoi(idPersonaStr)
	l, err := models.GetEventosByPersona(idPersona)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title PostTrEvento
// @Description create the TrEvento
// @Param	body		body 	models.TrEvento	true	"body for TrEvento content"
// @Success 201 {int} models.TrEvento
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrEventoController) Post() {
	var v models.TrEvento
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.AddTransaccionEvento(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TrEvento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TrEventoPut	true		"body for TrEvento content"
// @Success 200 {object} models.TrEventoPut
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *TrEventoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var v models.TrEventoPut
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.CalendarioEvento.Id = id
		if err := models.UpdateTransaccionEvento(&v); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Evento
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 404 not found resource
// @router /:id [delete]
func (c *TrEventoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.TrDeleteEvento(id); err == nil {
		c.Data["json"] = map[string]interface{}{"Id": id}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	}
	c.ServeJSON()
}
