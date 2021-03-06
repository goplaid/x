package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VTimePickerBuilder struct {
	tag *h.HTMLTagBuilder
}

func VTimePicker() (r *VTimePickerBuilder) {
	r = &VTimePickerBuilder{
		tag: h.Tag("v-time-picker"),
	}
	return
}

func (b *VTimePickerBuilder) Color(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VTimePickerBuilder) Dark(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Disabled(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Format(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("format", v)
	return b
}

func (b *VTimePickerBuilder) FullWidth(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":full-width", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) HeaderColor(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("header-color", v)
	return b
}

func (b *VTimePickerBuilder) Landscape(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":landscape", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Light(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Max(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("max", v)
	return b
}

func (b *VTimePickerBuilder) Min(v string) (r *VTimePickerBuilder) {
	b.tag.Attr("min", v)
	return b
}

func (b *VTimePickerBuilder) NoTitle(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":no-title", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Readonly(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Scrollable(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":scrollable", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) UseSeconds(v bool) (r *VTimePickerBuilder) {
	b.tag.Attr(":use-seconds", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Value(v interface{}) (r *VTimePickerBuilder) {
	b.tag.Attr(":value", v)
	return b
}

func (b *VTimePickerBuilder) Width(v int) (r *VTimePickerBuilder) {
	b.tag.Attr(":width", fmt.Sprint(v))
	return b
}

func (b *VTimePickerBuilder) Class(names ...string) (r *VTimePickerBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VTimePickerBuilder) ClassIf(name string, add bool) (r *VTimePickerBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VTimePickerBuilder) On(name string, value string) (r *VTimePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VTimePickerBuilder) Bind(name string, value string) (r *VTimePickerBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VTimePickerBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
