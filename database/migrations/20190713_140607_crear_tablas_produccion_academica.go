package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablasEventos_20190713_140607 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablasEventos_20190713_140607{}
	m.Created = "20190713_140607"

	migration.Register("CrearTablasEventos_20190713_140607", m)
}

// Run the migrations
func (m *CrearTablasEventos_20190713_140607) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20190713_140607_crear_tablas_eventos.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablasEventos_20190713_140607) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20190713_140607_crear_tablas_eventos.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
