<template lang="pug">
    .photo-card(:class="isVideo ? 'video' : ''")
        el-image(
            ref="elImg"
            fit="cover"
            :src="item.thumb"
            :preview-src-list="[item.origin]"
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
        img { cursor: pointer; }
        .image-slot { 
            height: 100%; color: #aaa; text-align: center;
            .el-icon-error { font-size: 100px; line-height: 250px; }
        }
        .el-image-viewer__video { position: relative; max-width: 1280px; }
        .el-image-viewer__img { z-index: 1; }
        .el-image-viewer__actions { z-index: 2; }
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
    img, .description, .description > *, 
    .el-icon-video-play { transition: opacity, transform, background, filter, color; transition-duration: .3s; }
    &:hover {
        img { transform: scale(1.1); filter: brightness(1.1); }
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
        item: Object
    },
    data() {
        return {
            viewerImage: null,
            viewerVideo: null,
        }
    },
    computed: {
        hasPermission() {
            const userInfo = this.$store.getters.userInfo
            const server = userInfo.guilds.find(server => server.id === this.item.serverId)
            return server.permissions === 2147483647
        },
        isVideo() {
            return /mp4$/.test(this.item.origin)
        },
    },
    methods: {
        onDeleteClick() {
            this.$emit('deleteClick', this.item.origin)
        },
        onImageClick() {
            const imageViewer = this.$children[0].$children[0]
            imageViewer.$data.index = 0

            const wrapper = imageViewer.$refs['el-image-viewer__wrapper']
            const canvas = wrapper.querySelector('.el-image-viewer__canvas')
            this.viewerImage = wrapper.querySelector('.el-image-viewer__img')

            // video -> img
            if (this.viewerVideo && !this.isVideo) {
                if (this.viewerImage) {
                    this.viewerImage.src = this.item.origin
                    this.viewerImage.style.display = 'block'
                }
                this.viewerVideo.style.display = 'none'
            }
            // video -> video
            if (this.viewerVideo && this.isVideo) {
                if (this.viewerImage) this.viewerImage.style.display = 'none'
                this.viewerVideo.style.display = 'block'
                this.viewerVideo.play()
            }
            // img -> video
            if (!this.viewerVideo && this.isVideo) {
                if (this.viewerImage) this.viewerImage.style.display = 'none'
                canvas.innerHTML += `
                    <video controls width="100%" autoplay class="el-image-viewer__video">
                        <source src="${this.item.origin}" type="video/mp4">
                    </video>
                `
                this.viewerVideo = wrapper.querySelector('.el-image-viewer__video')

                const closeBtn = wrapper.querySelector('.el-image-viewer__close')
                closeBtn.addEventListener('click', () => {
                    this.viewerVideo.pause()
                    this.viewerVideo.currentTime = 0
                })
            }
        },
    },
    mounted() {
        const imageViewer = this.$children[0].$children[0]
        const wrapper = imageViewer.$refs['el-image-viewer__wrapper']
        const mask = wrapper.querySelector('.el-image-viewer__mask')
        const closeBtn = wrapper.querySelector('.el-image-viewer__close')
        mask.addEventListener('click', () => {
            closeBtn.dispatchEvent(new MouseEvent('click'))
        })
    },
}
</script>