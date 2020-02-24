<template>
  <v-system-bar height="42px" color="primary">
    <v-spacer />
    
    <v-btn color="primary" small fab icon tile>
      <v-icon color="white">mdi-emoticon-happy-outline</v-icon>
    </v-btn>
    <v-menu :close-on-content-click="false" top offset-y>
      <template v-slot:activator="{ on }">
        <v-btn color="primary" small fab icon tile>
          <v-icon color="white" v-on="on">mdi-alert-circle-outline</v-icon>
        </v-btn>
      </template>
      <error-card :errors="errors" />
    </v-menu>

    <v-menu :close-on-content-click="false" top offset-y>
      <template v-slot:activator="{ on }">
        <v-btn color="primary" small fab icon tile>
          <v-icon color="white" v-on="on" v-if="discovery.connected">mdi-cloud-outline</v-icon>
          <v-icon color="white" v-else>mdi-cloud-off-outline</v-icon>
        </v-btn>
      </template>
      <login-card :discovery="discovery"/>
    </v-menu>
  </v-system-bar>
</template>

<script lang="ts">
import Vue from "vue";
import LoginCard from "@/components/users/Login.vue";
import ErrorCard from "@/components/cards/ErrorCard.vue";
import { get_discovery_version } from "@/services/discovery";
export default Vue.extend({
  data() {
    return {
      discovery: {
        connected: false,
        version: ""
      },
      errors:[{
        "Msg":"Init Service Failed!",
        "OccuredAt":"0000-0000"
      }],
    };
  },
  components:{
    "login-card": LoginCard,
    "error-card": ErrorCard,
  },
  mounted() {
    let self = this;
    get_discovery_version()
      .then(function(res: any) {
        self.discovery.connected = true;
        self.discovery.version = res;
      })
      .catch(function(err) {
        self.discovery.connected = false;
      });
  }
});
</script>

<style scoped>

</style>