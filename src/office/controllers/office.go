package controllers

import (
//"office/models"
//"office/util"
)

type OfficeController struct {
	baseController
}

/**
首页
*/
func (c *OfficeController) Index() {
	//c.list()
	c.TplName = c.controllerName + "/index.html"
}
