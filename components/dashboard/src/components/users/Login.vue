<template>
  <v-card width="300px">
    <v-card-title class="aid-discovery-title">Discovery</v-card-title>
    <v-card-text>
      <v-text-field label="Username"></v-text-field>
      <v-text-field label="Password"></v-text-field>
      <v-btn style="width:100%;" color="primary">Login</v-btn>
      <br />
      <br />
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
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    props: ["discovery", "width"]
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