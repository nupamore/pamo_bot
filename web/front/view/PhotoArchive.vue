<template lang="pug">
    div
        icon-select(placeholder="Select a group" :list="serverList" @change="onServerSelect")
        i.el-breadcrumb__separator.el-icon-arrow-right
        icon-select(placeholder="Select a user" :list="uploaderList" @change="onUploaderSelect")
        .divider
        el-row.photo-list
            el-col(
                v-for="item in imageList"
                :md="6"
                :sm="12"
                :xs="12"
            )
                photo-card(:item="item")
        .divider
        el-pagination.center(
            :page-size="12"
            :total="pageTotal"
            layout="prev, pager, next"
            @current-change="getImageList"
        )
</template>

<style lang="scss">
    .photo-list { font-size: 0; }
    .el-pager li.active { color: #fff; background: #409EFF; }
</style>

<script>
import dayjs from 'dayjs'
import { mapGetters } from 'vuex'
import IconSelect from 'component/IconSelect.vue'
import PhotoCard from 'component/PhotoCard.vue'

export default {
    components: {
        IconSelect,
        PhotoCard,
    },
    data() {
        return {
            serverId: '',
            uploaderId: 'All',
            uploaderList: [],
            imageList: [],
            currentPage: 1,
            pageTotal : 0,
        }
    },
    computed: {
        ...mapGetters(['serverList']),
    },
    methods: {
        toThumb(url) {
            return url.replace('cdn.discordapp.com', 'media.discordapp.net')
        },
        toOrigin(url) {
            return url.replace('media.discordapp.net', 'cdn.discordapp.com')
        },
        async getImageList(page) {
            this.currentPage = page
            const res = await fetch(`/images?galleryId=${this.serverId}&owner=${this.uploaderId}&page=${this.currentPage}`)
            const data = await res.json()
            this.imageList = data.images.map(image => ({
                name: image.OWNER,
                date: dayjs(image.REG_DATE).format('YYYY-MM-DD'),
                origin: image.ORIGIN_URL,
                thumb: this.toThumb(image.ORIGIN_URL) + '?width=400&height=225',
            }))
            this.pageTotal = data.total || this.pageTotal
        },
        async onServerSelect(serverId) {
            this.serverId = serverId
            this.getImageList(1)
            // uploader list
            const res = await fetch(`/uploaders?galleryId=${this.serverId}`)
            const data = await res.json()
            const uploaderList = data.map(item => ({
                value: item.owner,
                label: item.owner,
            }))
            this.uploaderList = [{ value: 'All', label: 'All' }, ...uploaderList]
        },
        onUploaderSelect(uploaderId) {
            this.uploaderId = uploaderId
            this.getImageList(1)
        },
    }
}
</script>