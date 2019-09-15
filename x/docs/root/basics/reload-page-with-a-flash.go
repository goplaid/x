package basics

import (
	ch "github.com/goplaid/x/codehighlight"
	"github.com/goplaid/x/docs/samples"
	"github.com/goplaid/x/docs/utils"
	"github.com/goplaid/x/md"
	. "github.com/theplant/htmlgo"
)

var ReloadPageWithAFlash = Components(
	md.Markdown(`
The results of an ~web.EventFunc~ could be:

- Go to a new page
- Reload the whole current page
- Refresh part of the current page

Let's demonstrate reload the whole current page: 
`),
	ch.Code(samples.ReloadWithFlashSample),
	utils.Demo("", samples.ReloadWithFlashPath),
	md.Markdown(`
~ctx.Flash~ Object is used to pass data between ~web.EventFunc~ to ~web.PageFunc~ just after the event func is executed. quite similar to [Rails's Flash](https://api.rubyonrails.org/classes/ActionDispatch/Flash.html).
Different is here you can pass in any complicated struct. as long as the page func to use that flash properly.

~er.Reload = true~ tells it will reload the whole page by running page func again, and with the result's body to replace the browser's html content. the event func and page func are executed in one AJAX request in the server.
`),
)
