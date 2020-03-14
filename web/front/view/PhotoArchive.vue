<template lang="pug">
    div
        icon-select(
            :placeholder="$t('UI.PLACEHOLDER.SELECT_DISCORD_SERVER')" 
            :list="serverList" 
            @change="onServerSelect"
        )
        i.el-breadcrumb__separator.el-icon-arrow-right
        icon-select(
            :placeholder="$t('UI.PLACEHOLDER.SELECT_UPLOADER')"  
            :current="uploaderId" 
            :list="uploaderList" 
            @change="onUploaderSelect"
        )
        .divider
        el-row.photo-list
            el-col(
                v-for="item in imageList"
                :md="6"
                :sm="12"
                :xs="12"
            )
                photo-card(
                    :item="item" 
                    :srcList="srcList" 
                    @deleteClick="onDeleteImage"
                    @showVideo="onShowVideo"
                    @closeVideo="onCloseVideo"
                )
        .divider
        el-pagination.center(
            :currentPage="currentPage"
            :page-size="12"
            :total="pageTotal"
            layout="prev, pager, next"
            @current-change="getImageList"
        )
        video.el-image-viewer__video(
            ref="video"
            width="100%"
            controls
            autoplay
        )
            source(:src="videoSrc" type="video/mp4")
</template>

<style lang="scss">
@import './../_vars.scss';

.photo-list {
    font-size: 0;
}
.el-pager li.active {
    color: #fff;
    background: #409eff;
}
.el-image-viewer__video {
    @include vCenter;
    z-index: 2001;
    max-width: 1280px;
    visibility: hidden;
    opacity: 0;
    &.show {
        visibility: visible;
        opacity: 1;
        transition: 0.5s;
    }
}
</style>

<script>
import dayjs from 'dayjs'
import { mapGetters } from 'vuex'
import IconSelect from 'component/IconSelect.vue'
import PhotoCard from 'component/PhotoCard.vue'
import filters from 'module/filters'
import api from 'module/api/gateway'

export default {
    components: {
        IconSelect,
        PhotoCard,
    },
    data() {
        return {
            serverId: '',
            uploaderId: '',
            uploaderList: [],
            imageList: [],
            currentPage: 1,
            pageTotal: 0,
            videoSrc: '',
        }
    },
    computed: {
        ...mapGetters(['serverList']),
        srcList() {
            return this.imageList.map(img => img.origin)
        },
    },
    methods: {
        async getImageList(page) {
            this.currentPage = page

            const data = await api('GET_IMAGES', {
                path: {
                    guildId: this.serverId,
                },
                params: {
                    owner: this.uploaderId,
                    page: this.currentPage,
                },
            })
            this.imageList = data.images.map(image => {
                const imgUrl = filters.imgUrl(
                    image.channel_id,
                    image.file_id,
                    image.file_name,
                )
                return {
                    userName: image.owner_name,
                    userId: image.owner_id,
                    serverId: this.serverId,
                    fileId: image.file_id,
                    date: dayjs(image.reg_date).format('YYYY-MM-DD'),
                    origin: imgUrl,
                    thumb: filters.origin2Thumb(imgUrl),
                }
            })
            this.pageTotal = data.total || this.pageTotal
        },
        async onServerSelect(serverId) {
            this.serverId = serverId
            this.uploaderId = ''
            this.getImageList(1)
            // uploader list
            const data = await api('GET_UPLOADERS', {
                path: {
                    guildId: this.serverId,
                },
            })
            const uploaderList = data.map(item => ({
                value: item.owner_id,
                label: item.owner_name,
                src: filters.avatarUrl(item.owner_id, item.owner_avatar),
                sub: item.amount,
            }))
            const sum = data.reduce((p, n) => p + n.amount, 0)
            this.uploaderList = [
                { value: '', label: 'All', sub: sum },
                ...uploaderList,
            ]
        },
        onUploaderSelect(uploaderId) {
            this.uploaderId = uploaderId
            this.getImageList(1)
        },
        async onDeleteImage(item) {
            const data = await api('DELETE_IMAGE', {
                path: {
                    imageId: item.fileId,
                },
            })
            if (data) {
                this.imageList = this.imageList.filter(
                    img => img.fileId !== item.fileId,
                )
            }
        },
        onShowVideo(url) {
            this.videoSrc = url
            this.$refs.video.classList.add('show')
            this.$refs.video.load()
        },
        onCloseVideo() {
            this.$refs.video.classList.remove('show')
            this.$refs.video.pause()
            this.$refs.video.currentTime = 0
        },
    },
}
</script>
