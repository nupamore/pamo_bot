import Vue from 'vue'
import VueRouter from 'vue-router'
import VueI18n from 'vue-i18n'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

import './common.scss'
import lang from './../lang/lang.js'
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
Vue.use(VueI18n)
const i18n = new VueI18n({
    locale: 'en',
    messages: lang,
})
Vue.use(ElementUI, {
    i18n: (key, value) => i18n.t(key, value),
})

new Vue({
    el: '#app',
    router,
    store,
    i18n,
    render: h => h(App),
    beforeCreate() {
        if (router.currentRoute.path === '/') {
            router.replace('/photo')
        }
        if (i18n.availableLocales.includes(navigator.language)) {
            i18n.locale = navigator.language
        }
    },
})
