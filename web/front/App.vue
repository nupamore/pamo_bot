<template lang="pug">
    el-container
        el-header
            top-nav
        el-main
            keep-alive
                router-view
        el-footer
</template>

<style lang="scss">
.el-container {
    max-width: 1600px;
    margin: 0 auto;
}
</style>

<script>
import TopNav from 'component/TopNav'

export default {
    components: {
        TopNav,
    },
}
</script>
