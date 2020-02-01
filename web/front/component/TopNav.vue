<template lang="pug">
    el-row
        el-col(:lg="20")
            el-menu(
                :default-active="activeIndex"
                mode="horizontal"
                router=true
            )
                el-menu-item(index="1" route="/sdl") Simple Dynamic Link
                el-menu-item(index="2" route="/photo") Photo Archive
        el-col(:lg="4")
            el-menu
                el-menu-item Login
</template>

<script>
export default {
    data() {
        return {
            activeIndex: '1'
        }
    },
}
</script>