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
            .left {{ item.userName }}
            .right {{ item.date }}
            el-popconfirm(
                v-if="hasPermission"
                :title="$t('UI.QUESTION.DELETE')"
                :confirmButtonText="$t('UI.CONFIRM')"
                :cancelButtonText="$t('UI.CANCEL')"
                @onConfirm="onDeleteClick"
            )
                el-button(slot="reference" icon="el-icon-close" v-on:click.prevent="")
        i.el-icon-video-play
</template>

<style lang="scss">
@import './../_vars.scss';

$icon-size: 7rem;

.photo-card {
    position: relative;
    display: block;
    overflow: hidden;

    // video icon
    &.video .el-icon-video-play {
        @include vCenter;
        top: -10px;
        width: $icon-size;
        height: $icon-size;
        color: #fff;
        font-size: $icon-size;
        opacity: 0.8;
        filter: drop-shadow(0px 0px 6px rgba(1, 1, 1, 0.5));
        pointer-events: none;
    }
    .el-image {
        width: 100%;
        height: 260px;
        @media screen and (max-width: 768px) {
            height: 35vw;
        }
        & > img {
            cursor: pointer;
        }
        .image-slot {
            height: 100%;
            color: #aaa;
            text-align: center;
            .el-icon-error {
                font-size: $icon-size;
                line-height: #{$icon-size * 2.5};
            }
        }
    }
    .description {
        position: absolute;
        bottom: 0;
        width: 100%;
        color: #fff;
        font-size: 1.1rem;
        background: rgba(0, 0, 0, 0.5);
        .left {
            padding: 0.6em;
            opacity: 0.8;
        }
        .right {
            padding: 0.6em;
            opacity: 0.5;
        }
        button {
            @include vCenter;
            color: #fff;
            background: none;
            border: none;
            transform: translateY(100%);
        }
    }

    // transition
    .el-image > img,
    .description,
    .description > *,
    .el-icon-video-play {
        transition: opacity, transform, background, filter, color;
        transition-duration: 0.3s;
    }
    &:hover {
        .el-image > img {
            transform: scale(1.1);
            filter: brightness(1.1);
        }
        .description {
            background: rgba(0, 0, 0, 1);
        }
        .description > * {
            opacity: 1;
        }
        .description button {
            transform: translateY(0);
        }
        .description button:hover {
            color: #f00;
        }
        .el-icon-video-play {
            transform: translateY(-10%);
            opacity: 1;
        }
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
            const server = userInfo.guilds.find(
                server => server.id === this.item.serverId,
            )
            const isMaster = server.permissions === 2147483647
            const isMine = this.item.userId === userInfo.id
            return isMaster || isMine
        },
    },
    methods: {
        isVideo(url) {
            return /mp4$/.test(url)
        },
        onDeleteClick() {
            this.$emit('deleteClick', this.item)
        },
        onImageClick() {
            const imageViewer = this.$children[0].$children[0]
            const index = this.srcList.findIndex(
                src => src === this.item.origin,
            )
            const src = this.srcList[index]
            imageViewer.$data.index = index
            if (this.isVideo(src)) this.$emit('showVideo', src)
        },
    },
    mounted() {
        const imageViewer = this.$children[0].$children[0]
        const wrapper = imageViewer.$refs['el-image-viewer__wrapper']
        const changeBtn = wrapper.querySelectorAll(
            '.el-image-viewer__prev, .el-image-viewer__next',
        )
        const closeBtn = wrapper.querySelector('.el-image-viewer__close')
        const mask = wrapper.querySelector('.el-image-viewer__mask')

        changeBtn.forEach(btn =>
            btn.addEventListener('click', () => {
                const index = imageViewer.$data.index
                const src = this.srcList[index]
                if (this.isVideo(src)) this.$emit('showVideo', src)
                else this.$emit('closeVideo')
            }),
        )
        closeBtn.addEventListener('click', () => {
            this.$emit('closeVideo')
        })
        mask.addEventListener('click', () => {
            closeBtn.dispatchEvent(new MouseEvent('click'))
        })
    },
}
</script>
