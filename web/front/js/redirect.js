
import Vue from 'vue/dist/vue.js'
import MuseUI from 'muse-ui'
import 'muse-ui/dist/muse-ui.css'
import '../css/redirect.scss'

Vue.use(MuseUI)
const app = new Vue({
    el: '#app',
    data: {
        normal: 10
    }
})