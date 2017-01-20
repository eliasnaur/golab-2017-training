package gorev

import (
	"Java/android/os"
	"Java/android/support/v7/app"

	// The Java package for "gorev".
	gopkg "Java/gorev"
)

type GoActivity struct {
	app.AppCompatActivity
}

func (a *GoActivity) OnCreate(this gopkg.GoActivity, b os.Bundle) {
	this.Super().OnCreate(b)
}
