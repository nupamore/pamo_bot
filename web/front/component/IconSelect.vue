<template lang="pug">
    el-select(
        v-model="value" 
        :placeholder="placeholder" 
        @change="onSelect"
        filterable
    )
        el-option(
            v-for="item in list"
            :key="item.value"
            :label="item.label"
            :value="item.value"
        )
            el-avatar(v-if="item.src" :src="item.src")
            .name {{ item.label }}
            .sub {{ item.sub }}
</template>

<style lang="scss">
    .el-select-dropdown__item { 
        height: 46px; line-height: 46px; 
        .el-avatar { float: left; margin: 3px 0; }
        .name { float: left; margin-left: 10px; }
        .sub { float: right; font-size: 0.8em; opacity: .8; }
    }
</style>

<script>
export default {
    props: {
        placeholder: String,
        list: Array,
        current: String,
    },
    data() {
        return {
            value: this.current,
        }
    },
    watch: {
        current(val) {
            this.value = val
        },
    },
    methods: {
        onSelect() {
            this.$emit('change', this.value)
        }
    },
}
</script>