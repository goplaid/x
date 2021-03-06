package e02_hello_material_button

import (
	"math/rand"
	"time"

	"github.com/goplaid/web"
	. "github.com/goplaid/x/material"
	h "github.com/theplant/htmlgo"
)

func randText() string {
	randTexts := []string{"Reload", "Hello world", "Do it", "Save..."}

	rand.Seed(int64(time.Now().Nanosecond()))

	return randTexts[rand.Intn(len(randTexts))]
}

func HelloButton(ctx *web.EventContext) (pr web.PageResponse, err error) {
	ctx.Hub.RegisterEventFunc("reload", reload)

	pr.Body = h.Div(
		web.Bind(Button(randText()).Variant(ButtonVariantRaised)).
			OnClick("reload"),
		web.Bind(Button(randText()).Variant(ButtonVariantRaised).Disabled(true)).
			OnClick("reload"),
		web.Bind(Button(randText()).Variant(ButtonVariantUnelevated)).
			OnClick("reload"),
		web.Bind(Button(randText()).Variant(ButtonVariantOutlined)).
			OnClick("reload"),
		web.Bind(Button(randText()).Variant(ButtonVariantText)).
			OnClick("reload"),
		web.Bind(Button("").Variant(ButtonVariantOutlined).
			Children(
				h.RawHTML(`
					<svg class="mdc-button__icon" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="#000000">
						<path fill="none" d="M0 0h24v24H0z"/>
						<path d="M23 12c0-6.07-4.93-11-11-11S1 5.93 1 12s4.93 11 11 11 11-4.93 11-11zM5 17.64C3.75 16.1 3 14.14 3 12c0-2.13.76-4.08 2-5.63v11.27zM17.64 5H6.36C7.9 3.75 9.86 3 12 3s4.1.75 5.64 2zM12 14.53L8.24 7h7.53L12 14.53zM17 9v8h-4l4-8zm-6 8H7V9l4 8zm6.64 2c-1.55 1.25-3.51 2-5.64 2s-4.1-.75-5.64-2h11.28zM21 12c0 2.14-.75 4.1-2 5.64V6.37c1.24 1.55 2 3.5 2 5.63z"/>
					</svg>
				`),
				h.RawHTML(randText()),
			)).
			OnClick("reload"),
	)
	return
}

func reload(ctx *web.EventContext) (r web.EventResponse, err error) {
	r.Reload = true
	return
}
