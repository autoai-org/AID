<template>
  <v-data-table
    :headers="headers"
    :items="items"
    item-key="name"
    class="elevation-1"
  >
  <template v-slot:item.action="{ item }">
      <v-chip v-for="action in actions" :key="action.text">{{action.text}}</v-chip>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    props: ['headers', 'items', 'actions'],
})
</script>

<style scoped>

</style>