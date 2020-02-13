import Vue from 'vue'
import VueRouter from 'vue-router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import './common.scss'
import store from './store/store.js'
import App from './App.vue'

// view
import SimpleDynamicLink from 'view/SimpleDynamicLink.vue'
import PhotoArchive from 'view/PhotoArchive.vue'
// promotion
import P_SimpleDynamicLink from 'view/P_SimpleDynamicLink.vue'
import P_PhotoArchive from 'view/P_PhotoArchive.vue'

const router = new VueRouter({
    mode: 'history',
    routes: [
        {
            path: '/sdl',
            component: SimpleDynamicLink,
        },
        {
            path: '/photo',
            component: PhotoArchive,
        },
    ],
})

Vue.use(VueRouter)
Vue.use(ElementUI)

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App),
    beforeCreate() {
        if (router.currentRoute.path === '/') {
            router.replace('/photo')
        }
    },
})
