<template>
  <v-card class="mx-auto" max-width="344" outlined>
    <v-card-text>
      <div>{{title}}</div>
      <p class="display-1 text--primary">{{content}}</p>
      <div class="text--primary">{{desc}}</div>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  props: ["title", "desc", "content"]
});
</script>

<style scoped>
</style>