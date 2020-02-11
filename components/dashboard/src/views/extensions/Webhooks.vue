<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-card-title>Webhooks</v-card-title>
    <v-card-text>
      <v-btn outlined @click="creating=true">New Webhook</v-btn>
      <div class="creating-webhook-form" v-if="creating">
      <v-row>
        <v-col class="mb8">
          <v-text-field v-model="payload_url" placeholder="Payload URL" single-line append-icon="mdi-webhook" />
        </v-col>
        <v-col class="mb4">
          <v-switch v-model="activate" class="ma-2" label="Activate Now?"></v-switch>
        </v-col>
      </v-row>
      <v-row>
          <v-btn color="primary" outlined @click="addNewWebhook()">Confirm</v-btn>
      </v-row>
      </div>
      <v-list></v-list>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import {alert} from "@/middlewares/store.mdw"
import { getWebhooks, addWebhook } from "@/middlewares/api.mdw";
export default Vue.extend({
  data() {
    return {
      webhooks: [],
      creating: false,
      activate: false,
      payload_url: ""
    };
  },
  methods: {
      addNewWebhook() {
          let status = "Disabled"
          if (this.activate) {
              status = "Enabled"
          }
          addWebhook(this.payload_url, status).then(function(res) {
              alert("Your webhook have been created!", "success")
          }).catch(function(err){
              alert("Failed to create webhooks: "+ err, "error",)
          })
      }
  },
  mounted() {
    let self = this;
    getWebhooks().then(function(res: any) {
      self.webhooks = res;
    });
  }
});
</script>

<style scoped>
.creating-webhook-form {
    margin-left: 20px;
}
</style>