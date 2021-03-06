import Tooltip from "ant-design-vue/lib/vc-tooltip/Tooltip";
import { getOptionProps } from 'ant-design-vue/lib/_util/props-util';
import './BranPopover.less';

export default {
	name: "BranPopover",
	props: Tooltip.props,
	data() {
		return {
			isVisible: this.defaultVisible,
		};
	},
	methods: {
		forceUpdate() {
			this.$forceUpdate()
		},
		// setIsVisible(v1) {
		// 	console.log("setIsVisible", v1)
		// }
	},
	render() {
		const props = getOptionProps(this)
		const { destroyTooltipOnHide, getTooltipContainer, afterVisibleChange, ...rest } = props
		const appContainer = () => {
			return document.getElementById("app")
		}
		const vcProps = {
			props: {
				destroyTooltipOnHide: true,
				getTooltipContainer: appContainer,
				// afterVisibleChange: this.setIsVisible,
				...rest
			},
			on: {
				...this.$listeners
			}
		}

		const { overlay } = this.$scopedSlots

		return (<Tooltip {...vcProps}>
			{overlay && <template slot="overlay">
				{overlay({ parent: this })}
			</template>}
			<template slot="default">
				{this.$scopedSlots.default({ parent: this })}
			</template>
		</Tooltip>
		)
	}
};
