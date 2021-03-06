package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VRadioGroupBuilder struct {
	tag *h.HTMLTagBuilder
}

func VRadioGroup(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	r = &VRadioGroupBuilder{
		tag: h.Tag("vw-radio-group").Children(children...),
	}
	return
}
func (b *VRadioGroupBuilder) ActiveClass(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VRadioGroupBuilder) AppendIcon(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VRadioGroupBuilder) BackgroundColor(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VRadioGroupBuilder) Color(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VRadioGroupBuilder) Column(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":column", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Dark(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Disabled(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Error(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) ErrorCount(v int) (r *VRadioGroupBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) ErrorMessages(v ...string) (r *VRadioGroupBuilder) {
	setErrorMessages(b.tag, v)
	return b
}

func (b *VRadioGroupBuilder) Height(v int) (r *VRadioGroupBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) HideDetails(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Hint(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VRadioGroupBuilder) Id(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VRadioGroupBuilder) Label(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VRadioGroupBuilder) Light(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Loading(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Mandatory(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":mandatory", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Max(v int) (r *VRadioGroupBuilder) {
	b.tag.Attr(":max", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Messages(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VRadioGroupBuilder) Multiple(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Name(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("name", v)
	return b
}

func (b *VRadioGroupBuilder) PersistentHint(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) PrependIcon(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VRadioGroupBuilder) Readonly(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Row(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":row", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Rules(v []string) (r *VRadioGroupBuilder) {
	b.tag.Attr(":rules", v)
	return b
}

func (b *VRadioGroupBuilder) Success(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) SuccessMessages(v string) (r *VRadioGroupBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VRadioGroupBuilder) ValidateOnBlur(v bool) (r *VRadioGroupBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VRadioGroupBuilder) Value(v interface{}) (r *VRadioGroupBuilder) {
	b.tag.Attr(":value", h.JSONString(v))
	return b
}

func (b *VRadioGroupBuilder) Children(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VRadioGroupBuilder) AppendChildren(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VRadioGroupBuilder) PrependChildren(children ...h.HTMLComponent) (r *VRadioGroupBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VRadioGroupBuilder) Class(names ...string) (r *VRadioGroupBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VRadioGroupBuilder) ClassIf(name string, add bool) (r *VRadioGroupBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VRadioGroupBuilder) On(name string, value string) (r *VRadioGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VRadioGroupBuilder) Bind(name string, value string) (r *VRadioGroupBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VRadioGroupBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
