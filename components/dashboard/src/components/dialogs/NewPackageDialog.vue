<template>
  <v-dialog v-model="shouldShow" persistant max-width="60%" @click:outside="closeOutside()">
    <v-card>
      <v-card-title>
        <span class="headline">Install New Package</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-text-field
            v-model="identifier"
            label="Package Identifier"
            hint="e.g. https://github.com/aidmodels/detectron or aicamp/chatbot"
          ></v-text-field>
          <v-alert type="info" v-if="success">Installtion Request submitted! <a target="_blank" class="logs" href="/system/logs">logs</a></v-alert>
          <v-alert type="error" v-if="error">Installation Request Failed. {{error}} <a class="logs" target="_blank" href="/system/logs">logs</a></v-alert>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-btn class="ma-2" color="primary" @click="callInstallPackage()">Confirm</v-btn>
        <v-btn class="ma-2" color="secondary" v-if="success || error" @click="finishInstall()">Finish</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script lang="ts">
import Vue from "vue";
import { installPackage } from "@/middlewares/api.mdw";
export default Vue.extend({
  data: () => ({
    shouldShow: false,
    identifier: "",
    error:"",
    success:false,
  }),
  props: ["show"],
  methods: {
    closeOutside() {
      console.log("closed");
      this.$emit("closed", true);
    },
    callInstallPackage() {
      let self = this
      installPackage(this.identifier).then((res:any) => {
          if (res.msg==="submitted success") {
              self.success = true
          } else {
              self.error = res.msg
          }
      });
    },
    finishInstall() {
        if (this.success) {
            location.reload()
        }
    }
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
.logs{
    color:white;
}
</style>