<template lang="pug">
    .photo-card(:class="isVideo(item.origin) ? 'video' : ''")
        el-image(
            ref="elImg"
            fit="cover"
            :src="item.thumb"
            :preview-src-list="srcList"
            v-on:click="onImageClick"
        )
            .image-slot(slot="placeholder" v-loading="true")
            .image-slot(slot="error")
                i.el-icon-error
        .description
            .left {{ item.name }}
            .right {{ item.date }}
            el-popconfirm(
                v-if="hasPermission"
                title="Are you sure to delete this?"
                @onConfirm="onDeleteClick"
            )
                el-button(slot="reference" icon="el-icon-close" v-on:click.prevent="")
        i.el-icon-video-play
</template>

<style lang="scss">
@import "./../_vars.scss";
.photo-card {
    display: block; position: relative; overflow: hidden;

    // video icon
    &.video .el-icon-video-play {
        @include vCenter;
        top: -10px; width: 100px; height: 100px;
        font-size: 100px; color: #fff; pointer-events: none;
        opacity: 0.8;
        filter: drop-shadow(0px 0px 6px rgba(1,1,1,0.5));
    }
    .el-image {
        width: 100%; height: 260px;
        & > img { cursor: pointer; }
        .image-slot { 
            height: 100%; color: #aaa; text-align: center;
            .el-icon-error { font-size: 100px; line-height: 250px; }
        }
    }
    .description { 
        position: absolute; width: 100%; bottom: 0; background: rgba(0,0,0,.5);
        font-size: 16px; color: #fff;
        .left { padding: 12px; opacity: .8; }
        .right { padding: 12px; opacity: .5; }
        button {
            @include vCenter;
            background: none; border: none;
            color: #fff; transform: translateY(100%);
        }
    }

    // transition
    & > img, .description, .description > *, 
    .el-icon-video-play { transition: opacity, transform, background, filter, color; transition-duration: .3s; }
    &:hover {
        & > img { transform: scale(1.1); filter: brightness(1.1); }
        .description { background: rgba(0,0,0,1); }
        .description > * { opacity: 1; }
        .description button { transform: translateY(0); }
        .description button:hover { color: #f00; }
        .el-icon-video-play { opacity: 1; transform: translateY(-10%); }
    }
}
</style>

<script>
export default {
    props: {
        item: Object,
        srcList: Array,
    },
    computed: {
        hasPermission() {
            const userInfo = this.$store.getters.userInfo
            const server = userInfo.guilds.find(server => server.id === this.item.serverId)
            return server.permissions === 2147483647
        },
    },
    methods: {
        isVideo(url) {
            return /mp4$/.test(url)
        },
        onDeleteClick() {
            this.$emit('deleteClick', this.item.origin)
        },
        onImageClick() {
            const imageViewer = this.$children[0].$children[0]
            const index = this.srcList.findIndex(src => src === this.item.origin)
            const src = this.srcList[index]
            imageViewer.$data.index = index
            if (this.isVideo(src)) this.$emit('showVideo', src)
        },
    },
    mounted() {
        const imageViewer = this.$children[0].$children[0]
        const wrapper = imageViewer.$refs['el-image-viewer__wrapper']
        const changeBtn = wrapper.querySelectorAll('.el-image-viewer__prev, .el-image-viewer__next')
        const closeBtn = wrapper.querySelector('.el-image-viewer__close')
        const mask = wrapper.querySelector('.el-image-viewer__mask')

        changeBtn.forEach(btn => btn.addEventListener('click', () => {
            const index = imageViewer.$data.index
            const src = this.srcList[index]
            if (this.isVideo(src)) this.$emit('showVideo', src)
            else this.$emit('closeVideo')
        }))
        closeBtn.addEventListener('click', () => {
            this.$emit('closeVideo')
        })
        mask.addEventListener('click', () => {
            closeBtn.dispatchEvent(new MouseEvent('click'))
        })
    },
}
</script>