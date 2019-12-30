<template>
  <comp-log :title="title" :message="message" />
</template>

<script lang="ts">
import Vue from "vue";
import Log from "@/components/Log.vue";
import { fetchSystemLog } from "@/middlewares/ws.mdw";

export default Vue.extend({
  data: () => ({
    title: "System Log",
    message: ""
  }),
  components: {
    "comp-log":Log
  },
  mounted() {
    fetchSystemLog((event: any) => {
        this.message = event
    });
  }
});
</script>

<style scoped>
</style>