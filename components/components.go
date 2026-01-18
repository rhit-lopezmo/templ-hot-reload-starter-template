package components

import (
	firstComponent "tmpl-first-time/components/first_component"

	"github.com/a-h/templ"
)

func FirstComponent(name string) templ.Component {
	return firstComponent.FirstComponent(name)
}
