<template lang="pug">
    el-select(v-model="value" :placeholder="placeholder")
        el-option(
            v-for="item in list"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        )
            el-avatar(:src="item.src")
            .name {{ item.label }}
</template>

<style lang="scss">
    .el-select-dropdown__item { 
        height: 46px; line-height: 46px; 
        .el-avatar { float: left; margin: 3px 0; }
        .name { float: left; margin-left: 10px; }
    }
</style>

<script>
export default {
    props: {
        placeholder: String,
        list: Array,
    },
    data() {
        return {
            value: ''
        }
    },
}
</script>