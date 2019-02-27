<template>
  <VBtn
    class="white--text"
    color="indigo"
    depressed
    v-bind="$attrs"
    v-on="$listeners"
  >
    <slot />
  </VBtn>
</template>

<script>
export default {
  inheritAttrs: false
}
</script>

<style lang="scss" scoped>
  .v-btn {
    border-radius: 4px;
    text-transform: none;
  }
</style>
