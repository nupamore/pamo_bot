<template lang="pug">
    div
        icon-select(placeholder="Select a group" :list="serverList")
        i.el-breadcrumb__separator.el-icon-arrow-right
        icon-select(placeholder="Select a user" :list="userList")
        .divider
        el-row.photo-list
            el-col(
                v-for="item in serverList"
                :md="6"
                :sm="12"
                :xs="12"
            )
                photo-card(:item="item")
        .divider
        el-pagination.center(:total="1000" layout="prev, pager, next")
</template>

<style lang="scss">
    .photo-list { font-size: 0; }
    .el-pager li.active { color: #fff; background: #409EFF; }
</style>

<script>
import IconSelect from 'component/IconSelect.vue'
import PhotoCard from 'component/PhotoCard.vue'

// test
const serverList = [
    { value: 1, label: 'test1', src: 'http://test.jpg' },
    { value: 2, label: 'test2', src: 'http://test.jpg' },
]
const userList = [
    { value: 1, label: 'user1' },
    { value: 2, label: 'user2' },
]

export default {
    components: {
        IconSelect,
        PhotoCard,
    },
    data() {
        return {
            serverList,
            userList,
        }
    },
}
</script>