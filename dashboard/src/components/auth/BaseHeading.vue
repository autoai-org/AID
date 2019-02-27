<template>
  <div
    class="mb-4"
    v-bind="$attrs"
    v-on="$listeners"
  >
    <h1
      class="headline"
      v-text="text"
    />
    <div
      class="subheading"
      v-text="sub"
    />
  </div>
</template>

<script>
export default {
  inheritAttrs: false,

  props: {
    sub: {
      type: String,
      default: undefined
    },
    text: {
      type: String,
      default: undefined
    }
  }
}
</script>
