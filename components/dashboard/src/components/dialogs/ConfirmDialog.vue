<template>
  <v-dialog v-model="shouldShow" max-width="60%" @click:outside="closeOutside()">
    <v-card>
      <v-card-title>{{title}}</v-card-title>
      <v-card-text>
        {{message}}
        <v-alert type="info" v-if="info">
          {{info}}
        </v-alert>
      </v-card-text>
      <v-card-actions>
        <v-btn class="ma-2" color="primary" @click="confirm()">Confirm</v-btn>
        <v-spacer></v-spacer>
        <v-btn v-if="!success && !error" class="ma-2" color="secondary" @click="cancel()">Cancel</v-btn>
        <v-btn v-if="success || error" class="ma-2" color="secondary" @click="finish()">Finish</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data: () => ({
    shouldShow: false,
    identifier: "",
    error: "",
    success: false,
    logid: ""
  }),
  props: ["show", "title", "message", "info"],
  methods: {
    closeOutside() {
      this.$emit("closed", true);
    },
    confirm() {
      this.$emit("confirmed", true);
    },
    cancel() {
      this.$emit("cancelled", true);
    },
    finish() {}
  },
  watch: {
    show(newVal) {
      if (newVal) {
        this.shouldShow = true;
      }
    }
  }
});
</script>

<style scoped>
</style>