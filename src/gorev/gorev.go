package gorev

import (
	"Java/android/os"
	"Java/android/support/v7/app"

	"Java/android/databinding/DataBindingUtil"
	"Java/gorev/databinding/ActivityMainBinding"

	rlayout "Java/gorev/R/layout"

	// The Java package for "gorev".
	gopkg "Java/gorev"
)

type GoActivity struct {
	app.AppCompatActivity
}

func (a *GoActivity) OnCreate(this gopkg.GoActivity, b os.Bundle) {
	this.Super().OnCreate(b)
	db := DataBindingUtil.SetContentView(this, rlayout.Activity_main)
	binding := ActivityMainBinding.Cast(db)
	binding.SetText("Hello from the reverse")
}
