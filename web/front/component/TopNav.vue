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
            el-menu
                el-menu-item.btn-login(@click="login")
                    el-image(fit="cover" src="https://discordapp.com/assets/f8389ca1a741a115313bede9ac02e2c0.svg")
                    .name Login
</template>

<style lang="scss">
    .btn-login {
        border-radius: 50px; max-width: 180px; font-size: 20px; margin: 0 0 0 auto;
        .el-image { float: left; width: 50px; height: 50px; margin: 4px 10px; }
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
            menuList
        }
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
        }
    },
    created() {
        this.findActiveIndex()
        this.$store.dispatch('userInfo')
    },
}
</script>