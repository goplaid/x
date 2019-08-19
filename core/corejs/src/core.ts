
import debounce from 'lodash/debounce';
import 'whatwg-fetch';
import Vue, { VueConstructor } from 'vue';

import {
	setPushState,
	EventData,
	setFormValue,
	jsonEvent,
} from './utils';

import {
	EventFuncID,
	EventResponse,
	PushState,
} from './types';

// Vue.config.productionTip = true;
declare var window: any;



export class Core {
	public debounce = debounce;

	private debounceFetchEventThenRefresh = debounce(this.fetchEventThenRefresh, 800);
	private form: FormData;
	private rootChangeCurrent: any;
	private changeCurrent: any;


	constructor(form: FormData, rootChangeCurrent: any, changeCurrent: any) {
		this.form = form;
		this.rootChangeCurrent = rootChangeCurrent;
		this.changeCurrent = changeCurrent;
	}

	public loadPage(pushState: PushState, popstate?: boolean) {
		for (const key of this.form.keys()) {
			this.form.delete(key);
		}

		this.fetchEventThenRefresh(
			{
				id: '__reload__',
				pushState,
			},
			{},
			popstate,
		);
	}

	public onpopstate(event: any) {
		this.loadPage(event.state, true);
	}

	public fetchEvent(
		eventFuncId: EventFuncID,
		event: EventData,
		popstate?: boolean,
	): Promise<EventResponse> {

		const { newEventFuncId, eventURL } = setPushState(
			eventFuncId,
			(window.location.pathname + '?' + window.location.search),
			window.history,
			popstate,
		);

		const eventData = JSON.stringify({
			eventFuncId: newEventFuncId,
			event,
		});

		this.form.set('__event_data__', eventData);
		return fetch(eventURL, {
			method: 'POST',
			body: this.form,
		}).then((r) => {
			return r.json();
		}).then((r: EventResponse) => {

			if (r.pageTitle) {
				document.title = r.pageTitle;
			}

			if (r.redirectURL) {
				// window.location.replace(r.redirectURL);
				this.loadPage({ url: r.redirectURL });
			}

			if (r.pushState) {
				this.loadPage(r.pushState);
			}
			return r;
		});
	}

	public componentByTemplate(template: string, afterLoaded?: any): VueConstructor {
		return Vue.extend({
			provide: { core: this },
			template: '<div>' + template + '</div>', // to make only one root.
			methods: this.newVueMethods(),
			mounted() {
				this.$nextTick(() => {
					if (afterLoaded) {
						afterLoaded(this);
					}
				});
			},
			data() {
				return {
					vars: {},
				};
			},
		});
	}

	public setFormValue(fieldName: string, val: any) {
		setFormValue(this.form, fieldName, val);
	}

	// public getFormValue(fieldName: string): string {
	// 	return getFormValue(this.form, fieldName);
	// }

	// public getFormValueAsArray(fieldName: string): string[] {
	// 	return getFormValueAsArray(this.form, fieldName);
	// }

	private fetchEventThenRefresh(
		eventFuncId: EventFuncID,
		event: EventData,
		popstate?: boolean,
	) {
		this.fetchEvent(eventFuncId, event, popstate)
			.then((r: EventResponse) => {
				if (r.reloadPortals && r.reloadPortals.length > 0) {
					for (const portalName of r.reloadPortals) {
						const portal = window.branLazyPortals[portalName];
						if (portal) {
							portal.reload();
						}
					}
					return r;
				}

				if (r.updatePortals && r.updatePortals.length > 0) {
					for (const pu of r.updatePortals) {
						const portal = window.branLazyPortals[pu.name];
						if (portal) {
							let afterLoaded;
							if (pu.afterLoaded) {
								afterLoaded = new Function('comp', pu.afterLoaded);
							}
							portal.changeCurrent(this.componentByTemplate(pu.schema, afterLoaded));
						}
					}
					return r;
				}

				if (r.schema && r.reload) {
					this.rootChangeCurrent(this.componentByTemplate(r.schema));
					return r;
				}

				if (r.schema) {
					this.changeCurrent(this.componentByTemplate(r.schema));
					return r;
				}

				return r;
			});
	}

	private newVueMethods(): any {
		const self = this;
		return {
			topage(pushState: any) {
				self.loadPage(pushState);
			},
			triggerEventFunc(eventFuncId: EventFuncID, evt: any) {
				self.fetchEventThenRefresh(eventFuncId, jsonEvent(evt));
			},
			oninput(eventFuncId?: EventFuncID, fieldName?: string, evt?: any) {
				self.controlsOnInput(eventFuncId, fieldName, evt);
			},
		};
	}

	private controlsOnInput(
		eventFuncId?: EventFuncID,
		fieldName?: string,
		evt?: any,
	) {
		if (fieldName) {
			this.form.set(fieldName, evt.target.value);
		}
		if (eventFuncId) {
			this.debounceFetchEventThenRefresh(eventFuncId, jsonEvent(evt));
		}
	}

}
