package material

import (
	"context"

	"github.com/sunfmin/bran/ui"
	h "github.com/theplant/htmlgo"
)

type CardBuilder struct {
	children   []h.HTMLComponent
	classNames []string

	actionButtons []h.HTMLComponent
	actionIcons   []h.HTMLComponent
}

func (b *CardBuilder) Class(names ...string) (r *CardBuilder) {
	b.classNames = names
	return b
}

func Card(children ...h.HTMLComponent) (r *CardBuilder) {
	r = &CardBuilder{}
	r.children = children
	return
}

func (b *CardBuilder) ActionButtons(buttons ...h.HTMLComponent) (r *CardBuilder) {
	b.actionButtons = buttons
	return b
}

func (b *CardBuilder) ActionIcons(icons ...h.HTMLComponent) (r *CardBuilder) {
	b.actionIcons = icons
	return b
}

func (b *CardBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	ui.Injector(ctx).PutStyle(cardstyles)
	root := h.Div(
		h.Div(b.children...).
			Class("mdc-card__primary-action").
			Attr("tabindex", "0"),
	).Class(append([]string{"mdc-card"}, b.classNames...)...)

	if len(b.actionButtons) > 0 || len(b.actionIcons) > 0 {
		actions := h.Div().Class("mdc-card__actions")
		root.AppendChildren(actions)

		if len(b.actionButtons) > 0 {
			actions.AppendChildren(
				h.Div().Class("mdc-card__action-buttons").
					Children(b.actionButtons...),
			)
		}

		if len(b.actionIcons) > 0 {
			actions.AppendChildren(
				h.Div().Class("mdc-card__action-icons").
					Children(b.actionIcons...),
			)
		}
	}

	return root.MarshalHTML(ctx)
}

const cardstyles = `
/*!
 Material Components for the Web
 Copyright (c) 2019 Google Inc.
 License: MIT
*/
.mdc-card {
  border-radius: 4px;
  background-color: #fff;
  /* @alternate */
  background-color: var(--mdc-theme-surface, #fff);
  box-shadow: 0px 2px 1px -1px rgba(0, 0, 0, 0.2), 0px 1px 1px 0px rgba(0, 0, 0, 0.14), 0px 1px 3px 0px rgba(0, 0, 0, 0.12);
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
}

.mdc-card--outlined {
  box-shadow: 0px 0px 0px 0px rgba(0, 0, 0, 0.2), 0px 0px 0px 0px rgba(0, 0, 0, 0.14), 0px 0px 0px 0px rgba(0, 0, 0, 0.12);
  border-width: 1px;
  border-style: solid;
  border-color: #e0e0e0;
}

.mdc-card__media {
  position: relative;
  box-sizing: border-box;
  background-repeat: no-repeat;
  background-position: center;
  background-size: cover;
}
.mdc-card__media::before {
  display: block;
  content: "";
}

.mdc-card__media:first-child {
  border-top-left-radius: inherit;
  border-top-right-radius: inherit;
}

.mdc-card__media:last-child {
  border-bottom-left-radius: inherit;
  border-bottom-right-radius: inherit;
}

.mdc-card__media--square::before {
  margin-top: 100%;
}

.mdc-card__media--16-9::before {
  margin-top: 56.25%;
}

.mdc-card__media-content {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  box-sizing: border-box;
}

.mdc-card__primary-action {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  position: relative;
  outline: none;
  color: inherit;
  text-decoration: none;
  cursor: pointer;
  overflow: hidden;
}

.mdc-card__primary-action:first-child {
  border-top-left-radius: inherit;
  border-top-right-radius: inherit;
}

.mdc-card__primary-action:last-child {
  border-bottom-left-radius: inherit;
  border-bottom-right-radius: inherit;
}

.mdc-card__actions {
  display: flex;
  flex-direction: row;
  align-items: center;
  box-sizing: border-box;
  min-height: 52px;
  padding: 8px;
}

.mdc-card__actions--full-bleed {
  padding: 0;
}

.mdc-card__action-buttons,
.mdc-card__action-icons {
  display: flex;
  flex-direction: row;
  align-items: center;
  box-sizing: border-box;
}

.mdc-card__action-icons {
  color: rgba(0, 0, 0, 0.6);
  flex-grow: 1;
  justify-content: flex-end;
}

.mdc-card__action-buttons + .mdc-card__action-icons {
  /* @noflip */
  margin-left: 16px;
  /* @noflip */
  margin-right: 0;
}
[dir=rtl] .mdc-card__action-buttons + .mdc-card__action-icons, .mdc-card__action-buttons + .mdc-card__action-icons[dir=rtl] {
  /* @noflip */
  margin-left: 0;
  /* @noflip */
  margin-right: 16px;
}

.mdc-card__action {
  display: inline-flex;
  flex-direction: row;
  align-items: center;
  box-sizing: border-box;
  justify-content: center;
  cursor: pointer;
  -webkit-user-select: none;
     -moz-user-select: none;
      -ms-user-select: none;
          user-select: none;
}
.mdc-card__action:focus {
  outline: none;
}

.mdc-card__action--button {
  /* @noflip */
  margin-left: 0;
  /* @noflip */
  margin-right: 8px;
  padding: 0 8px;
}
[dir=rtl] .mdc-card__action--button, .mdc-card__action--button[dir=rtl] {
  /* @noflip */
  margin-left: 8px;
  /* @noflip */
  margin-right: 0;
}
.mdc-card__action--button:last-child {
  /* @noflip */
  margin-left: 0;
  /* @noflip */
  margin-right: 0;
}
[dir=rtl] .mdc-card__action--button:last-child, .mdc-card__action--button:last-child[dir=rtl] {
  /* @noflip */
  margin-left: 0;
  /* @noflip */
  margin-right: 0;
}

.mdc-card__actions--full-bleed .mdc-card__action--button {
  justify-content: space-between;
  width: 100%;
  height: auto;
  max-height: none;
  margin: 0;
  padding: 8px 16px;
  /* @noflip */
  text-align: left;
}
[dir=rtl] .mdc-card__actions--full-bleed .mdc-card__action--button, .mdc-card__actions--full-bleed .mdc-card__action--button[dir=rtl] {
  /* @noflip */
  text-align: right;
}

.mdc-card__action--icon {
  margin: -6px 0;
  padding: 12px;
}

.mdc-card__action--icon:not(:disabled) {
  color: rgba(0, 0, 0, 0.6);
}

@-webkit-keyframes mdc-ripple-fg-radius-in {
  from {
    -webkit-animation-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
            animation-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    -webkit-transform: translate(var(--mdc-ripple-fg-translate-start, 0)) scale(1);
            transform: translate(var(--mdc-ripple-fg-translate-start, 0)) scale(1);
  }
  to {
    -webkit-transform: translate(var(--mdc-ripple-fg-translate-end, 0)) scale(var(--mdc-ripple-fg-scale, 1));
            transform: translate(var(--mdc-ripple-fg-translate-end, 0)) scale(var(--mdc-ripple-fg-scale, 1));
  }
}

@keyframes mdc-ripple-fg-radius-in {
  from {
    -webkit-animation-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
            animation-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
    -webkit-transform: translate(var(--mdc-ripple-fg-translate-start, 0)) scale(1);
            transform: translate(var(--mdc-ripple-fg-translate-start, 0)) scale(1);
  }
  to {
    -webkit-transform: translate(var(--mdc-ripple-fg-translate-end, 0)) scale(var(--mdc-ripple-fg-scale, 1));
            transform: translate(var(--mdc-ripple-fg-translate-end, 0)) scale(var(--mdc-ripple-fg-scale, 1));
  }
}
@-webkit-keyframes mdc-ripple-fg-opacity-in {
  from {
    -webkit-animation-timing-function: linear;
            animation-timing-function: linear;
    opacity: 0;
  }
  to {
    opacity: var(--mdc-ripple-fg-opacity, 0);
  }
}
@keyframes mdc-ripple-fg-opacity-in {
  from {
    -webkit-animation-timing-function: linear;
            animation-timing-function: linear;
    opacity: 0;
  }
  to {
    opacity: var(--mdc-ripple-fg-opacity, 0);
  }
}
@-webkit-keyframes mdc-ripple-fg-opacity-out {
  from {
    -webkit-animation-timing-function: linear;
            animation-timing-function: linear;
    opacity: var(--mdc-ripple-fg-opacity, 0);
  }
  to {
    opacity: 0;
  }
}
@keyframes mdc-ripple-fg-opacity-out {
  from {
    -webkit-animation-timing-function: linear;
            animation-timing-function: linear;
    opacity: var(--mdc-ripple-fg-opacity, 0);
  }
  to {
    opacity: 0;
  }
}
.mdc-ripple-surface--test-edge-var-bug {
  --mdc-ripple-surface-test-edge-var: 1px solid #000;
  visibility: hidden;
}
.mdc-ripple-surface--test-edge-var-bug::before {
  border: var(--mdc-ripple-surface-test-edge-var);
}

.mdc-card__primary-action {
  --mdc-ripple-fg-size: 0;
  --mdc-ripple-left: 0;
  --mdc-ripple-top: 0;
  --mdc-ripple-fg-scale: 1;
  --mdc-ripple-fg-translate-end: 0;
  --mdc-ripple-fg-translate-start: 0;
  -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
  will-change: transform, opacity;
}
.mdc-card__primary-action::before, .mdc-card__primary-action::after {
  position: absolute;
  border-radius: 50%;
  opacity: 0;
  pointer-events: none;
  content: "";
}
.mdc-card__primary-action::before {
  transition: opacity 15ms linear, background-color 15ms linear;
  z-index: 1;
}
.mdc-card__primary-action.mdc-ripple-upgraded::before {
  -webkit-transform: scale(var(--mdc-ripple-fg-scale, 1));
          transform: scale(var(--mdc-ripple-fg-scale, 1));
}
.mdc-card__primary-action.mdc-ripple-upgraded::after {
  top: 0;
  /* @noflip */
  left: 0;
  -webkit-transform: scale(0);
          transform: scale(0);
  -webkit-transform-origin: center center;
          transform-origin: center center;
}
.mdc-card__primary-action.mdc-ripple-upgraded--unbounded::after {
  top: var(--mdc-ripple-top, 0);
  /* @noflip */
  left: var(--mdc-ripple-left, 0);
}
.mdc-card__primary-action.mdc-ripple-upgraded--foreground-activation::after {
  -webkit-animation: mdc-ripple-fg-radius-in 225ms forwards, mdc-ripple-fg-opacity-in 75ms forwards;
          animation: mdc-ripple-fg-radius-in 225ms forwards, mdc-ripple-fg-opacity-in 75ms forwards;
}
.mdc-card__primary-action.mdc-ripple-upgraded--foreground-deactivation::after {
  -webkit-animation: mdc-ripple-fg-opacity-out 150ms;
          animation: mdc-ripple-fg-opacity-out 150ms;
  -webkit-transform: translate(var(--mdc-ripple-fg-translate-end, 0)) scale(var(--mdc-ripple-fg-scale, 1));
          transform: translate(var(--mdc-ripple-fg-translate-end, 0)) scale(var(--mdc-ripple-fg-scale, 1));
}
.mdc-card__primary-action::before, .mdc-card__primary-action::after {
  top: calc(50% - 100%);
  /* @noflip */
  left: calc(50% - 100%);
  width: 200%;
  height: 200%;
}
.mdc-card__primary-action.mdc-ripple-upgraded::after {
  width: var(--mdc-ripple-fg-size, 100%);
  height: var(--mdc-ripple-fg-size, 100%);
}
.mdc-card__primary-action::before, .mdc-card__primary-action::after {
  background-color: #000;
}
.mdc-card__primary-action:hover::before {
  opacity: 0.04;
}
.mdc-card__primary-action:not(.mdc-ripple-upgraded):focus::before, .mdc-card__primary-action.mdc-ripple-upgraded--background-focused::before {
  transition-duration: 75ms;
  opacity: 0.12;
}
.mdc-card__primary-action:not(.mdc-ripple-upgraded)::after {
  transition: opacity 150ms linear;
}
.mdc-card__primary-action:not(.mdc-ripple-upgraded):active::after {
  transition-duration: 75ms;
  opacity: 0.12;
}
.mdc-card__primary-action.mdc-ripple-upgraded {
  --mdc-ripple-fg-opacity: 0.12;
}
`