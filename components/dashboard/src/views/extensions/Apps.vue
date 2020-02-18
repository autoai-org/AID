<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-card-title>Third Party Apps</v-card-title>
    <v-card-text>
      <v-row class="mb-12">
        <v-col
          cols="2"
          v-for="(item,index) in installed_apps"
          :key="index"
          class="d-flex flex-column justify-center app_image_box align-center"
        >
          <v-avatar color="indigo" size="64pt">
              <img v-if="item.img_url" :src="item.img_url"/>
            <span v-else class="white--text headline">{{getInitials(item.name)}}</span>
          </v-avatar>
          {{item.name}}
          <v-btn color="secondary" outlined>Running History</v-btn> <br/>
          <v-btn color="primary">View Configuration</v-btn>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      installed_apps: [
        {
          name: "Dive Docker Image Analysis",
        },
        {
          name: "Deploy to Amazon Fargate™",
          img_url:"https://mherman.org/presentations/flask-on-fargate/images/fargate.png"
        }
      ]
    };
  },
  methods:{
      getInitials(fullString: string) {
          return fullString.split(" ").map((n)=>n[0]).join(".");
      }
  }
});
</script>

<style scoped>
.app_image_box{
  text-align: center;
  color: black!important;
}
</style>