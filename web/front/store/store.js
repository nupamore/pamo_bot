import Vue from 'vue'
import Vuex from 'vuex'
import 'babel-polyfill'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        userInfo: Object,
    },
    getters: {
        userInfo: state => state.userInfo,
        serverList(state) {
            return state.userInfo.guilds?.map(guild => ({
               value: guild.id,
               label: guild.name,
               src: `https://cdn.discordapp.com/icons/${guild.id}/${guild.icon}.jpg`,
            }))
        },
    },
    mutations: {
        userInfo(state, info) {
            state.userInfo = info
        },
    },
    actions: {
        async userInfo(state) {
            const res = await fetch('/profile')
            const data = await res.json()
            state.commit('userInfo', data)
        },
        logout(state) {
            location.href = '/logout'    
        },
    },
})