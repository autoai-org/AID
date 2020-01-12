<template>
  <v-data-table
    :headers="headers"
    :items="items"
    item-key="name"
    class="elevation-1"
    :show-select="select"
    v-model="selected"
  >
  <template v-slot:item.action="{ item }">
      <v-chip v-for="action in actions" :key="action.text" @click="onClick(action.text, item)">{{action.text}}</v-chip>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    data:() =>({
      selected:[],
    }),
    props: ['headers', 'items', 'actions', 'select'],
    methods:{
      onClick(action:string, item:object) {
        this.$emit('actionClicked', {action, item})
      }
    }
})
</script>

<style scoped>

</style>