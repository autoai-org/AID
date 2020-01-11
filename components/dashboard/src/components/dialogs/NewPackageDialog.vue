<template>
  <v-dialog v-model="shouldShow" persistant max-width="60%" @click:outside="closeOutside()">
    <v-card>
      <v-card-title>
        <span class="headline">Install New Package</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-text-field v-model="identifier" label="Package Identifier" hint="e.g. https://github.com/aidmodels/detectron or aicamp/chatbot"></v-text-field>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-btn class="ma-2" color="primary" @click="callInstallPackage()">Confirm</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script lang="ts">
import Vue from "vue";
import { installPackage } from '@/middlewares/api.mdw'
export default Vue.extend({
  data: () => ({
    shouldShow: false,
    identifier: ""
  }),
  props: ["show"],
  methods: {
      closeOutside() {
          console.log('closed')
          this.$emit('closed', true)
      },
      callInstallPackage() {
          installPackage(this.identifier)
      }
  },
  watch:{
    show(newVal) {
        if (newVal) {
            this.shouldShow = true
        }
    }
  },
});
</script>
<style scoped>
</style>