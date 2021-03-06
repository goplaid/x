package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VStepperStepBuilder struct {
	tag *h.HTMLTagBuilder
}

func VStepperStep(children ...h.HTMLComponent) (r *VStepperStepBuilder) {
	r = &VStepperStepBuilder{
		tag: h.Tag("v-stepper-step").Children(children...),
	}
	return
}

func (b *VStepperStepBuilder) Color(v string) (r *VStepperStepBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VStepperStepBuilder) Complete(v bool) (r *VStepperStepBuilder) {
	b.tag.Attr(":complete", fmt.Sprint(v))
	return b
}

func (b *VStepperStepBuilder) CompleteIcon(v string) (r *VStepperStepBuilder) {
	b.tag.Attr("complete-icon", v)
	return b
}

func (b *VStepperStepBuilder) EditIcon(v string) (r *VStepperStepBuilder) {
	b.tag.Attr("edit-icon", v)
	return b
}

func (b *VStepperStepBuilder) Editable(v bool) (r *VStepperStepBuilder) {
	b.tag.Attr(":editable", fmt.Sprint(v))
	return b
}

func (b *VStepperStepBuilder) ErrorIcon(v string) (r *VStepperStepBuilder) {
	b.tag.Attr("error-icon", v)
	return b
}

func (b *VStepperStepBuilder) Rules(v []string) (r *VStepperStepBuilder) {
	b.tag.Attr(":rules", v)
	return b
}

func (b *VStepperStepBuilder) Step(v int) (r *VStepperStepBuilder) {
	b.tag.Attr(":step", fmt.Sprint(v))
	return b
}

func (b *VStepperStepBuilder) Class(names ...string) (r *VStepperStepBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VStepperStepBuilder) ClassIf(name string, add bool) (r *VStepperStepBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VStepperStepBuilder) On(name string, value string) (r *VStepperStepBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VStepperStepBuilder) Bind(name string, value string) (r *VStepperStepBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VStepperStepBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
