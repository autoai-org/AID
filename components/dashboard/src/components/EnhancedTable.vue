<template>
  <v-data-table
    :headers="headers"
    :items="items"
    item-key="name"
    class="elevation-1"
  >
  </v-data-table>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: ['headers', 'items'],
})
</script>

<style scoped>

</style>