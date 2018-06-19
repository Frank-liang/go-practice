package controllers

import "github.com/Frank-liang/go/lenslocked.com/views"

//import "lenslocked.com/views"

func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "static/home"),
		Contact: views.NewView(
			"bootstrap", "static/contact"),
	}
}

type Static struct {
	Home    *views.View
	Contact *views.View
}
