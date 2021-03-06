package vuetify

import (
	"context"
	"fmt"

	h "github.com/theplant/htmlgo"
)

type VBreadcrumbsItemBuilder struct {
	tag *h.HTMLTagBuilder
}

func VBreadcrumbsItem(children ...h.HTMLComponent) (r *VBreadcrumbsItemBuilder) {
	r = &VBreadcrumbsItemBuilder{
		tag: h.Tag("v-breadcrumbs-item").Children(children...),
	}
	return
}

func (b *VBreadcrumbsItemBuilder) ActiveClass(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("active-class", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Append(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":append", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Disabled(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":disabled", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Exact(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":exact", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) ExactActiveClass(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("exact-active-class", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Href(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("href", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Nuxt(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":nuxt", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Replace(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":replace", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Ripple(v bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(":ripple", fmt.Sprint(v))
	return b
}

func (b *VBreadcrumbsItemBuilder) Tag(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("tag", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Target(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("target", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) To(v string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr("to", v)
	return b
}

func (b *VBreadcrumbsItemBuilder) Class(names ...string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Class(names...)
	return b
}

func (b *VBreadcrumbsItemBuilder) ClassIf(name string, add bool) (r *VBreadcrumbsItemBuilder) {
	b.tag.ClassIf(name, add)
	return b
}

func (b *VBreadcrumbsItemBuilder) On(name string, value string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-on:%s", name), value)
	return b
}

func (b *VBreadcrumbsItemBuilder) Bind(name string, value string) (r *VBreadcrumbsItemBuilder) {
	b.tag.Attr(fmt.Sprintf("v-bind:%s", name), value)
	return b
}

func (b *VBreadcrumbsItemBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
