package routers

import (
	"deviceshouse/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		/*beego.NSNamespace("/users",
			beego.NSRouter("/register",
				&controllers.UserController{},
				"post:Register"),
			beego.NSRouter("/login",
				&controllers.UserController{},
				"post:Login"),
			beego.NSRouter("/logout",
				&controllers.UserController{},
				"post:Logout"),
			beego.NSRouter("/passwd",
				&controllers.UserController{},
				"post:Passwd"),
			beego.NSRouter("/uploads",
				&controllers.UserController{},
				"post:Uploads"),
			beego.NSRouter("/downloads",
				&controllers.UserController{},
				"get:Downloads"),
		),*/
		beego.NSNamespace("/users",
			beego.NSRouter("/:id",
				&controllers.UserController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.UserController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/sessions",
			beego.NSRouter("/:id",
				&controllers.SessionController{},
				"delete:Delete"),
			beego.NSRouter("/",
				&controllers.SessionController{},
				"post:Post"),
		),
		beego.NSNamespace("/devices",
			beego.NSRouter("/:id",
				&controllers.DeviceController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/devicetypes",
			beego.NSRouter("/:id",
				&controllers.DeviceTypeController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceTypeController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/deviceattrs",
			beego.NSRouter("/:id",
				&controllers.DeviceAttrController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceAttrController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/deviceattrinfos",
			beego.NSRouter("/:id",
				&controllers.DeviceAttrInfoController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceAttrInfoController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/deviceboilerainfos",
			beego.NSRouter("/:id",
				&controllers.DeviceBoilerAInfoController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceBoilerAInfoController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/deviceboilerbinfos",
			beego.NSRouter("/:id",
				&controllers.DeviceBoilerBInfoController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.DeviceBoilerBInfoController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/userbinddevices",
			beego.NSRouter("/:id",
				&controllers.UserBindDeviceController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.UserBindDeviceController{},
				"get:GetAll;post:Post"),
		),
		beego.NSNamespace("/roles",
			beego.NSRouter("/:id",
				&controllers.RoleController{},
				"get:GetOne;put:Put;delete:Delete"),
			beego.NSRouter("/",
				&controllers.RoleController{},
				"get:GetAll;post:Post"),
		),
	)
	beego.AddNamespace(ns)
}
