package overlay

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type PopoverBuilder struct {
	children []h.HTMLComponent

	triggerElement h.HTMLComponent
	tag            *h.HTMLTagBuilder
}

func Popover(children ...h.HTMLComponent) (r *PopoverBuilder) {
	r = &PopoverBuilder{
		tag: h.Tag("bran-popover"),
	}
	r.Placement("top")
	r.Trigger("click")
	r.children = children
	return
}

func (b *PopoverBuilder) TriggerElement(v h.HTMLComponent) (r *PopoverBuilder) {
	b.triggerElement = v
	return b
}

func (b *PopoverBuilder) Trigger(v string) (r *PopoverBuilder) {
	b.tag.Attr("trigger", v)
	return b
}

func (b *PopoverBuilder) DefaultVisible(v bool) (r *PopoverBuilder) {
	if v {
		b.tag.Attr(":default-visible", fmt.Sprint(v))
	}
	return b
}

func (b *PopoverBuilder) PrefixClass(v string) (r *PopoverBuilder) {
	b.tag.Attr(":prefix-cls", v)
	return b
}

func (b *PopoverBuilder) OverlayClassName(v string) (r *PopoverBuilder) {
	b.tag.Attr("overlay-class-name", v)
	return b
}

func (b *PopoverBuilder) Placement(v string) (r *PopoverBuilder) {
	b.tag.Attr("placement", v)
	return b
}

func (b *PopoverBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {

	b.tag.Children(b.triggerElement)
	b.tag.AppendChildren(
		h.Template(b.children...).Attr("v-slot:overlay", "{ parent }"),
	)

	return b.tag.MarshalHTML(ctx)
}
