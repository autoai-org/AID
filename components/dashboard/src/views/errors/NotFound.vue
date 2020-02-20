<template>
  <v-container>
    <div class="not-found-text">
      <img src="https://http.cat/404"/>
      <p class="display-1 text--primary not-found-text">404 Not Found.</p>
      <v-btn color="primary" href="https://github.com/autoai-org/aid/issues" target="_blank">Report the issue?</v-btn>
    </div>
  </v-container>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({});
</script>
<style scoped>
.not-found-text {
  align-self: center;
  text-align: center;
}
</style>