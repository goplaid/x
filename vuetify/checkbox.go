package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VCheckboxBuilder struct {
	tag *h.HTMLTagBuilder
}

func VCheckbox(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	r = &VCheckboxBuilder{
		tag: h.Tag("vw-checkbox").Children(children...),
	}
	//r.FalseValue("false").TrueValue("true")
	return
}
func (b *VCheckboxBuilder) AppendIcon(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("append-icon", v)
	return b
}

func (b *VCheckboxBuilder) BackgroundColor(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("background-color", v)
	return b
}

func (b *VCheckboxBuilder) Color(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("color", v)
	return b
}

func (b *VCheckboxBuilder) Dark(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":dark", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Disabled(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Error(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":error", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) ErrorCount(v int) (r *VCheckboxBuilder) {
	b.tag.Attr(":error-count", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) ErrorMessages(v ...string) (r *VCheckboxBuilder) {
	setErrorMessages(b.tag, v)
	return b
}

func (b *VCheckboxBuilder) FalseValue(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("false-value", v)
	return b
}

func (b *VCheckboxBuilder) Height(v int) (r *VCheckboxBuilder) {
	b.tag.Attr(":height", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) HideDetails(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":hide-details", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Hint(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("hint", v)
	return b
}

func (b *VCheckboxBuilder) Id(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("id", v)
	return b
}

func (b *VCheckboxBuilder) Indeterminate(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":indeterminate", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) IndeterminateIcon(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("indeterminate-icon", v)
	return b
}

func (b *VCheckboxBuilder) InputValue(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":input-value", h.JSONString(v))
	return b
}

func (b *VCheckboxBuilder) Label(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("label", v)
	return b
}

func (b *VCheckboxBuilder) Light(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":light", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Loading(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":loading", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Messages(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("messages", v)
	return b
}

func (b *VCheckboxBuilder) Multiple(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":multiple", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) OffIcon(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("off-icon", v)
	return b
}

func (b *VCheckboxBuilder) OnIcon(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("on-icon", v)
	return b
}

func (b *VCheckboxBuilder) PersistentHint(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":persistent-hint", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) PrependIcon(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("prepend-icon", v)
	return b
}

func (b *VCheckboxBuilder) Readonly(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":readonly", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Ripple(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Rules(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr("rules", v)
	return b
}

func (b *VCheckboxBuilder) Success(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":success", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) SuccessMessages(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("success-messages", v)
	return b
}

func (b *VCheckboxBuilder) TrueValue(v string) (r *VCheckboxBuilder) {
	b.tag.Attr("true-value", v)
	return b
}

func (b *VCheckboxBuilder) ValidateOnBlur(v bool) (r *VCheckboxBuilder) {
	b.tag.Attr(":validate-on-blur", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Value(v interface{}) (r *VCheckboxBuilder) {
	b.tag.Attr(":value", fmt.Sprint(v))
	return b
}

func (b *VCheckboxBuilder) Children(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	b.tag.Children(children...)
	return b
}

func (b *VCheckboxBuilder) AppendChildren(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	b.tag.AppendChildren(children...)
	return b
}

func (b *VCheckboxBuilder) PrependChildren(children ...h.HTMLComponent) (r *VCheckboxBuilder) {
	b.tag.PrependChildren(children...)
	return b
}

func (b *VCheckboxBuilder) Class(names ...string) (r *VCheckboxBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VCheckboxBuilder) ClassIf(name string, add bool) (r *VCheckboxBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VCheckboxBuilder) On(name string, value string) (r *VCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VCheckboxBuilder) Bind(name string, value string) (r *VCheckboxBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VCheckboxBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
