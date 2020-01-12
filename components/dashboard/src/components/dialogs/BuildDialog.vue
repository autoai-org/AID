<template>
  <v-dialog v-model="shouldShow" persistant max-width="60%" @click:outside="closeOutside()">
    <v-card>
      <v-card-title>Are you sure?</v-card-title>
      <v-card-text>
        You are going to build
        <b>{{vendor}}/{{packageName}}/{{solver}}</b>
        <v-alert type="info" v-if="success">
          Build Request submitted!
          <a target="_blank" class="logs" @click="navLog()">logs</a>
        </v-alert>
        <v-alert type="error" v-if="error">
          Build Request Failed. {{error}}
          <a class="logs" target="_blank" @click="navLog()">logs</a>
        </v-alert>
      </v-card-text>
      <v-card-actions>
        <v-btn class="ma-2" color="primary" @click="callBuildImage()">Confirm</v-btn>
        <v-btn class="ma-2" color="secondary" v-if="success || error" @click="finishBuild()">Finish</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { buildImage } from "@/middlewares/api.mdw";
import Vue from "vue";
export default Vue.extend({
  data: () => ({
    shouldShow: false,
    identifier: "",
    error: "",
    success: false,
    logid:""
  }),
  props: ["show", "vendor", "packageName", "solver"],
  methods: {
    navLog() {
        this.$router.push({path:"/system/logs", query:{
            logid: this.logid
        }})
    },
    closeOutside() {
      this.$emit("closed", true);
    },
    finishBuild() {
      this.shouldShow = false;
      this.$emit("closed", true);
    },
    callBuildImage() {
      let self = this;
      buildImage(this.vendor, this.packageName, this.solver)
        .then((res: any) => {
          self.success = true
          self.logid = res.logid
          console.log(res);
        })
        .catch(err => {
          console.log(err.response);
          self.error = err.response.data.msg;
        });
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
.logs {
  color: white;
}
</style>