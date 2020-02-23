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
            const guilds = state.userInfo.guilds
            if (!guilds) return []
            return guilds.map(guild => ({
                value: guild.id,
                label: guild.name,
                src: `https://cdn.discordapp.com/icons/${guild.id}/${guild.icon}.jpg`,
            }))
        },
    },
    mutations: {
        SET_USER_INFO(state, info) {
            state.userInfo = info
        },
    },
    actions: {
        async getUserInfo(state) {
            const res = await fetch('/profile')
            const data = await res.json()
            state.commit('SET_USER_INFO', data)
        },
        logout(state) {
            location.href = '/logout'
        },
    },
})
