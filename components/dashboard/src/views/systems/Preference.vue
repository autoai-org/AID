<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-list>
      <v-subheader>General</v-subheader>
      <v-list-item>
        <v-list-item-content>
          <v-switch v-model="config.RemoteReport" class="ma-2" label="Allow Remote Report"></v-switch>
        </v-list-item-content>
      </v-list-item>
    </v-list>
    <v-card-actions>
        <v-btn color="primary" @click="submit">Apply</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { fetchConfig, updateConfig } from "@/middlewares/api.mdw";
import Vue from "vue";
import { alert } from '@/middlewares/store.mdw'
export default Vue.extend({
  data: () => ({
    config: {}
  }),
  mounted() {
    let self = this;
    fetchConfig().then(function(res: any) {
      self.config = res;
    });
  },
  methods: {
      submit() {
          let self: Vue = this
          updateConfig(this.config).then(function(res) {
              alert("Your preference has been updated! ","Successfully")
          }).catch(function(err){
              alert("Your preference cannot be updated! " + err,"Failed!")
          })
      }
  }
});
</script>

<style scoped>
</style>