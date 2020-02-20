<template>
  <div>
    <tree-view :data="jsonObject" />
  </div>
</template>

<script>
export default {
  props: {
    jsonObject: {
      type: Object,
      default: function () {
        return {}
      }
    }
  },
  created () {
    console.log(this.jsonObject)
  }
}
</script>

<style>

</style>
