<template lang="pug">
    el-row
        el-col(:sm="20" :md="21")
            el-menu(
                :default-active="activeIndex"
                mode="horizontal"
                router=true
            )
                el-menu-item(
                    v-for="item in menuList"
                    :index="item.index"
                    :route="item.route"
                ) {{ item.name }}
        el-col(:sm="4" :md="3")
            el-menu(v-if="!userInfo.id")
                el-menu-item.btn-profile(@click="login")
                    el-image(fit="cover" src="https://discordapp.com/assets/f8389ca1a741a115313bede9ac02e2c0.svg")
                    .login Login
            el-menu(v-else)
                el-popconfirm(title="Are you sure to logout?" @onConfirm="logout")
                    el-menu-item.btn-profile(slot="reference")
                        el-image(fit="cover" :src="`https://cdn.discordapp.com/avatars/${ userInfo.id }/${ userInfo.avatar }.jpg`")
                        .name {{ userInfo.username }}
                        .code {{ '#' + userInfo.discriminator }}
</template>

<style lang="scss">
    .btn-profile {
        max-width: 180px; font-size: 20px; margin: 0 0 0 auto;
        .el-image { float: left; width: 50px; height: 50px; margin: 4px 10px; }
        .name { font-size: 18px; line-height: 2em; }
        .code { font-size: 12px; line-height: 1em; }
    }
</style>

<script>
const menuList = [
    { index: '1', route: '/sdl', name: 'Simple Dynamic Link' },
    { index: '2', route: '/photo', name: 'Photo Archive' },
]

function findActiveIndex($route) {
    return menuList.find(item => item.route === $route.path).index
}

export default {
    data() {
        return {
            activeIndex: String,
            menuList,
        }
    },
    computed: {
        userInfo() {
            return this.$store.getters.userInfo
        },
    },
    watch: {
        $route (to, from){
            this.findActiveIndex()
        }
    },
    methods: {
        findActiveIndex() {
            this.activeIndex = menuList.find(item => item.route === this.$route.path).index
        },
        login() {
            location.href = "/auth/discord"
        },
        logout() {
            this.$confirm('Are you sure logout?')
            .then(() => {
                this.$store.dispatch('logout')
            })
        },
    },
    created() {
        this.findActiveIndex()
        this.$store.dispatch('userInfo')
    },
}
</script>