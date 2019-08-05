package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrEvento struct {
	CalendarioEvento    *CalendarioEvento
	EncargadosEvento  *[]EncargadoEvento
	TiposPublico        *[]TipoPublico
}

type TrEventoPut struct {
	CalendarioEvento    *CalendarioEvento
	EncargadosEvento  *[]EncargadoEvento
	TiposPublico        *[]TipoPublico
	EncargadosEventoBorrados  *[]EncargadoEvento
	TiposPublicoBorrados       *[]TipoPublico
}

// GetEventosByPersona Transacción para consultar todos los eventos con toda la información de las mismas
func GetEventosByPersona(persona int) (v []interface{}, err error) {
	o := orm.NewOrm()
	var encargados []*EncargadoEvento
	if _, err := o.QueryTable(new(EncargadoEvento)).RelatedSel().Filter("EncargadoId",persona).Filter("CalendarioEventoId__Activo",true).All(&encargados); err == nil{
		for _, encargado := range encargados {

			evento := encargado.CalendarioEventoId

			var encargadosEvento []EncargadoEvento
			if _, err := o.QueryTable(new(EncargadoEvento)).RelatedSel().Filter("CalendarioEventoId__Id",evento.Id).Filter("Activo", true).All(&encargadosEvento); err != nil{
				return nil, err
			}

			var tiposPublico []TipoPublico
			if _, err := o.QueryTable(new(TipoPublico)).RelatedSel().Filter("CalendarioEventoId__Id",evento.Id).Filter("Activo", true).All(&tiposPublico); err != nil{
				return nil, err
			}

			v = append(v,map[string]interface{}{
				"CalendarioEvento": &evento,
				"EncargadosEvento": &encargadosEvento,
				"TiposPublico": &tiposPublico,
			})
		}

		return v, nil
	}
	return nil, err
}

// AddTransaccionEvento Transacción para registrar toda la información de un evento
func AddTransaccionEvento(m *TrEvento) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	
	if idEvento, errTr := o.Insert(m.CalendarioEvento); errTr == nil {
		fmt.Println(idEvento)

		for _, v := range *m.EncargadosEvento {
			v.CalendarioEventoId.Id = int(idEvento)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

		for _, v := range *m.TiposPublico {
			v.CalendarioEventoId.Id = int(idEvento)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

// UpdateTransaccionEvento updates Evento by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransaccionEvento(m *TrEventoPut) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	v := CalendarioEvento{Id: m.CalendarioEvento.Id}
	// ascertain id exists in the database
	if errTr := o.Read(&v); errTr == nil {
		var num int64
		if num, errTr = o.Update(m.CalendarioEvento,"Descripcion","FechaInicio","FechaFin"); errTr == nil {
			fmt.Println("Number of records updated (CalendarioEvento) in database:", num)
			
			// Eliminar tipospublico
			for _, v := range *m.TiposPublicoBorrados {
				v.Activo = false
				if _, errTr = o.Update(&v,"Activo"); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
			}
			// Agregar y actualizar TiposPublico
			for _, v := range *m.TiposPublico {
				if (v.Id != 0) {
					if _, errTr = o.Update(&v,"Nombre"); errTr != nil {
						err = errTr
						fmt.Println(err)
						_ = o.Rollback()
						return
					}
				} else {
					if _, errTr = o.Insert(&v); errTr != nil {
						err = errTr
						fmt.Println(err)
						_ = o.Rollback()
						return
					}
				}
			}
			// Agregar encagadosEvennto
			for _, v := range *m.EncargadosEvento {
				if _, errTr = o.Insert(&v); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
			}
			// Eliminar encargadodsEvento
			for _, v := range *m.EncargadosEventoBorrados {
				v.Activo = false
				if _, errTr = o.Update(&v,"Activo"); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
			}

			_ = o.Commit()
		}	else {
			err = errTr
			fmt.Println(err)
			_ = o.Rollback()
			return
		}
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}

// TrDeleteEvento deletes Evento by Id and returns error if
// the record to be deleted doesn't exist
func TrDeleteEvento(id int) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	v := CalendarioEvento{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		// if num, err = o.Delete(&CalendarioEvento{Id: id}); err == nil {
			// fmt.Println("Number of records deleted in database:", num)
		if num, err = o.Update(&CalendarioEvento{Id: id, Activo: false },"Activo"); err == nil {
			fmt.Println("Number of records updated in database:", num)
		} else {
			fmt.Println(err)
			_ = o.Rollback()
			return
		}

		// Buscar tipos publico e inactivarlos.
		var tiposPublico []*TipoPublico
		if _, err = o.QueryTable(new(TipoPublico)).RelatedSel().Filter("CalendarioEventoId__Id",id).All(&tiposPublico); err == nil{
			for _, tipoPublico := range tiposPublico {
				tipoPublico.Activo = false
				if num, err = o.Update(tipoPublico,"Activo"); err == nil {
					fmt.Println("Tipo publico inactivated in database:", num)
				} else {
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
			}
		} else {
			fmt.Println(err)
			_ = o.Rollback()
			return
		}


		// Buscar encargados e inactivarlos a todos.
		var encargadosEvento []*EncargadoEvento
		if _, err = o.QueryTable(new(EncargadoEvento)).RelatedSel().Filter("CalendarioEventoId__Id",id).All(&encargadosEvento); err == nil{
			for _, encargadoEvento := range encargadosEvento {
				encargadoEvento.Activo = false
				if num, err = o.Update(encargadoEvento,"Activo"); err == nil {
					fmt.Println("Encargado Evento inactivated in database:", num)
				} else {
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
			}
		} else {
			fmt.Println(err)
			_ = o.Rollback()
			return
		}		

		_ = o.Commit()
	} else {
		fmt.Println(err)
		_ = o.Rollback()
	}
	return
}