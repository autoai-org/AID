<template>
  <v-dialog v-model="dialog" width="50%">
    <v-card>
      <v-card-title class="headline grey lighten-2" primary-title>{{title}}</v-card-title>
      <v-card-text>{{info}}</v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" text @click="dialog = false">OK</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data: () => ({
      dialog: false,
  }),
  props: ["info", "title"],
  watch: {
      info (newVal) {
          this.dialog= newVal
      }
  }
});
</script>

<style scoped>
</style>