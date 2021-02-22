<template>
  <v-card width="300px">
    <v-card-title class="aid-discovery-title">Discovery</v-card-title>
    <v-card-text v-if="!logged_in && access_token_required==='' ">
      <v-text-field label="Username" v-model="username"></v-text-field>
      <v-text-field label="Password" type="password" v-model="password"></v-text-field>
      <v-btn style="width:100%;" color="primary" @click="login">Login</v-btn>
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
          <v-btn icon tile @click="login_third_party('google')">
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
    <v-card-text v-if="logged_in">
      <v-avatar color="indigo">
        <img
        v-if=" avatar!=='' "
        :src="avatar"
        alt="Avatar"
      >
        <v-icon dark v-else>mdi-account-circle</v-icon>
      </v-avatar>
      {{ username }}
    </v-card-text>
    <v-card-text v-if=" !logged_in && access_token_required!=='' ">
      <p>Paste your access token:</p>
      <v-text-field label="Access Token" v-model="access_token"></v-text-field>
      <v-btn style="width:100%;" color="primary" @click="confirm">Confirm</v-btn>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { login_with_pass, login_with_google, get_profile_google } from "../../services/discovery";
export default Vue.extend({
  data: () => ({
    username: "",
    avatar:"",
    password: "",
    access_token: "",
    logged_in: false,
    access_token_required: '',
  }),
  props: ["discovery", "width"],
  methods: {
    login() {
      let self = this
      login_with_pass(this.username, this.password).then(function (res: any) {
        localStorage.setItem("aid_access_token", res.access_token);
        self.logged_in = true;
      });
    },
    login_third_party(source: string) {
      if (source === 'google') {
        this.access_token_required = "google"
        console.log(this.access_token_required)
        login_with_google()
      }
    },
    confirm() {
      let self = this
      if (this.access_token_required === 'google') {
        get_profile_google(this.access_token).then(function(res:any){
          self.avatar = res.picture
          self.username = res.name
          self.logged_in = true
        })
      }
    }
  },
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