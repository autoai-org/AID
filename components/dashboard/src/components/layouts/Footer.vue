<template>
  <v-system-bar height="42px" color="primary">
    <v-spacer />
    <v-btn color="primary" small fab icon tile>
      <v-icon color="white">mdi-emoticon-happy-outline</v-icon>
    </v-btn>
    <v-menu :close-on-content-click="false" top offset-y>
      <template v-slot:activator="{ on }">
        <v-btn color="primary" small fab icon tile>
          <v-icon color="white" v-on="on" v-if="discovery.connected">mdi-cloud-outline</v-icon>
          <v-icon color="white" v-else>mdi-cloud-off-outline</v-icon>
        </v-btn>
      </template>
      <v-card width="300px">
        <v-card-title class="aid-discovery-title">Discovery</v-card-title>
        <v-card-text>
          <v-text-field label="Username"></v-text-field>
          <v-text-field label="Password"></v-text-field>
          <v-btn style="width:100%;" color="primary">Login</v-btn> <br/><br/>
          <v-btn style="width:100%;" color="secondary">Signup</v-btn>
          <v-row class="md4">
            <v-col>
              <v-btn icon tile>
                <v-icon>fab fa-github</v-icon>
              </v-btn>
            </v-col>
            <v-col>
              <v-btn icon tile>
                <v-icon>mdi-google</v-icon>
              </v-btn>
            </v-col>
            <v-col>
              <v-btn icon tile>
                <v-icon>mdi-gitlab</v-icon>
              </v-btn>
            </v-col>
            <v-col>
              <v-btn icon tile>
                <v-icon>mdi-twitter</v-icon>
              </v-btn>
            </v-col>
          </v-row>
          <v-spacer></v-spacer>
          <p class="aid-discovery-version">Version: {{discovery.version}}</p>
        </v-card-text>
      </v-card>
    </v-menu>
  </v-system-bar>
</template>

<script lang="ts">
import Vue from "vue";
import { get_discovery_version } from "@/services/discovery";
export default Vue.extend({
  data() {
    return {
      discovery: {
        connected: false,
        version: ""
      }
    };
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
.aid-discovery-title {
  justify-content: center;
}
.aid-discovery-version {
  justify-content: center;
  text-align: center;
}
</style>