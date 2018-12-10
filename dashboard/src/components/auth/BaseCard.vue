<template>
  <VCard
    class="elevation-4"
    v-bind="$attrs"
    v-on="$listeners"
  >
    <slot />
  </VCard>
</template>

<style lang="scss" scoped>
  .v-card {
    border-radius: 6px;
  }
</style>
