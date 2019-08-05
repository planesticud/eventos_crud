package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "GetAllByPersona",
			Router:           `/:persona`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
