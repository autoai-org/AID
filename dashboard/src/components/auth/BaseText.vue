<template>
  <Component
    :is="component"
    :target="$attrs.href ? '_blank' : undefined"
    class="indigo--text subheading font-weight-bold"
    v-bind="$attrs"
    v-on="$listeners"
  >
    <slot />
    <VIcon
      v-if="$attrs.href"
      color="indigo lighten-1"
      small
    >
      mdi-open-in-new
    </VIcon>
  </Component>
</template>

<script>
  export default {
    inheritAttrs: false,

    computed: {
      component () {
        return this.$attrs.to
          ? 'router-link'
          : this.$attrs.href ? 'a' : 'span'
      }
    }
  }
</script>

<style lang="scss" scoped>
a {
  text-decoration: none;
}
</style>
