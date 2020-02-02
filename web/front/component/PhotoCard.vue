<template lang="pug">
    a.photo-card(:href="item.origin" target="blank")
        el-image(
            fit="cover"
            :src="item.thumb"
        )
            .image-slot(slot="placeholder" v-loading="true")
            .image-slot(slot="error")
                i.el-icon-error
        .description
            .left {{ item.name }}
            .right {{ item.date }}
</template>

<style lang="scss">
    .photo-card {
        display: block; position: relative; overflow: hidden;
        
        .el-image {
            width: 100%; height: 260px;
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
        }

        img, .description { transition: .3s; }
        &:hover {
            img { transform: scale(1.1); filter: brightness(1.1); }
            .description { transform: translateY(100%); }
        }
    }
</style>

<script>
export default {
    props: {
        item: Object
    },
}
</script>