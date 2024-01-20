<template>
    <div class='layout-footer mt15' v-show="isDelayFooter">
        <div class="layout-footer-warp"><el-link href="https://github.com/aenjoy/iom" type="warning">{{ version }} ❤️</el-link>
    </div></div>
</template>
<script>
import axios from 'axios'
import { toRefs, reactive } from 'vue';
import { onBeforeRouteUpdate } from 'vue-router';
axios.defaults.baseURL = '/api'
export default {
    name: "footer",
    data() {
        return {
            version: ''
        };
    },
    methods: {
        getversion() {
            axios.get("/version").then(r => {
                //console.log(r.data)
                if (r.data != null) {
                    this.version = 'IOM DashBoard (Server Version:' + r.data + ')'
                }
                else {
                    this.version = 'IOM DashBoard'
                }
            }, e => {
                this.version = 'IOM DashBoard'
            });
        }
    },
    mounted() {
        this.getversion()
    },
    setup() {
		const state = reactive({
			isDelayFooter: true,
		});
		// 路由改变时，等主界面动画加载完毕再显示 footer
		onBeforeRouteUpdate(() => {
			state.isDelayFooter = false;
			setTimeout(() => {
				state.isDelayFooter = true;
			}, 800);
		});
		return {
			...toRefs(state),
		};
	},
}
</script>
<style scoped lang="scss">
.layout-footer {
	width: 100%;
	display: flex;
	&-warp {
		margin: auto;
		color: var(--el-text-color-secondary);
		text-align: center;
		animation: logoAnimation 0.3s ease-in-out;
	}
}
</style>