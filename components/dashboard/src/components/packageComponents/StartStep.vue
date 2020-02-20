<template>
  <v-card>
    <v-card-text>
      <v-row class="mb-12">
        <v-col cols="2" v-for="(item,index) in containers" :key="index" class="d-flex flex-column justify-center solver_image_box">
          <v-icon size="64pt">mdi-package-variant</v-icon>
            <B>{{item.ID}}</B>
            {{item.CreatedAt}}
            <v-btn color="primary" @click="triggerStartContainer(item.ID)">Start</v-btn> <br/>
            <v-btn color="secondary" outlined>Delete Container</v-btn> <br/>
        </v-col>
        <v-col cols="2" class="d-flex flex-column justify-center solver_image_box solver_image_create_box">
          <v-icon size="64pt" @click="createNewContainer">mdi-plus</v-icon>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { createContainer, startContainer } from "@/middlewares/api.mdw"
export default Vue.extend({
  data:()=>({
    parsedContainers: [],
    code:"",
    showCodeEditorDialog:false
  }),
  props: ["containers", "solverName", "imageId"],
  methods:{
    showCodeEditor(vendorName:string, packageName:string) {
      let self = this
      console.log(this.solverName)
      this.showCodeEditorDialog = true
    },
    createNewContainer () {
        console.log(this.imageId)
    },
    triggerStartContainer (containerId:string) {
      startContainer(containerId).then(function(res){
        location.reload()
      })
    }
  },
  mounted () { 
    this.parsedContainers = this.containers.map(function(each:any){
        each.ID=each.ID.slice(0,6)
        return each
    })
  }
});
</script>
<style scoped>
.solver_image_box {
  text-align: center;
  color: black!important;
}
.solver_image_create_box {
  border: 1px;
  border-style: dotted;
  border-color: gray;
}
</style>