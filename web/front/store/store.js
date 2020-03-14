import Vue from 'vue'
import Vuex from 'vuex'
import 'babel-polyfill'
import api from 'module/api/gateway'

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
        async getUserInfo(context) {
            const data = await api('GET_PROFILE')
            context.commit('SET_USER_INFO', data)
        },
        async logout() {
            await api('LOGOUT')
            location.reload()
        },
    },
})
