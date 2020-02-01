<template lang="pug">
    a.photo-card(:href="item.src" target="blank")
        el-image(
            fit="cover"
            :src="item.src"
        )
            .image-slot(slot="placeholder" v-loading="true")
            .image-slot(slot="error")
                i.el-icon-circle-close
        .description
            .left {{ item.label }}
            .right {{ item.value }}
</template>

<style lang="scss">
    .photo-card {
        display: block; position: relative; overflow: hidden;
        
        .el-image {
            width: 100%; height: 250px;
            .image-slot { 
                height: 100%; font-size: 100px; color: #aaa;
                text-align: center; line-height: 250px;
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
            img { transform: scale(1.1); filter: brightness(1.2); }
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