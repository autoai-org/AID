<template>
  <comp-log :title="title" :message="message" />
</template>

<script lang="ts">
import Vue from "vue";
import Log from "@/components/Log.vue";
import { fetchBuildLog } from "@/middlewares/ws.mdw";

export default Vue.extend({
  data: () => ({
    title: "Build aicamp-detectron-image_segmentation.201912300000",
    message: ""
  }),
  components: {
    "comp-log":Log
  },
  mounted() {
    fetchBuildLog((event: any) => {
        this.message = event
    }, "aicamp-detectron-image_segmentation.201912300000");
  }
});
</script>

<style scoped>
</style>