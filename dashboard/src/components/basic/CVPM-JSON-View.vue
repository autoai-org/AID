<template>
    <div>
        <tree-view :data="jsonObject"></tree-view>
    </div>
</template>

<script>
export default {
  props: ['jsonObject'],
  created () {
    console.log(this.jsonObject)
  }
}
</script>

<style>

</style>
