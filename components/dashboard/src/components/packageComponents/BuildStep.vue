<template>
  <v-card>
    <v-card-text>
      <v-row class="mb-12">
        <v-col cols="2" v-for="(item,index) in images" :key="index" class="d-flex flex-column justify-center solver_image_box">
          <v-icon size="64pt">mdi-package-variant-closed</v-icon>
            <B>{{item.number}}-{{item.solverName}}</B>
            {{item.CreatedAt}}
            <v-btn color="primary" @click="newContainer(item.ID)">Create Container</v-btn> <br/>
            <v-btn color="secondary" outlined @click="showCodeEditor(item.vendor, item.package)">View Dockerfile</v-btn> <br/>
            <v-btn color="secondary" outlined>Delete Image</v-btn>
        </v-col>
        <v-col cols="2" class="d-flex flex-column justify-center solver_image_box solver_image_create_box">
          <v-btn icon @click="triggerBuild" large>
            <v-icon size="64pt">mdi-plus</v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-card-text>
    <build-dialog
      :show="showBuildDialog"
      @closed="showBuildDialog=false"
      :vendor="vendorName"
      :packageName="packageName"
      :solver="solverName"
    ></build-dialog>
    <code-editor-dialog :show="showCodeEditorDialog" :code="code" @closed="showCodeEditorDialog=false"></code-editor-dialog>
  </v-card>
</template>
<script lang="ts">
import BuildDialog from '@/components/dialogs/BuildDialog.vue'
import CodeEditorDialog from '@/components/dialogs/CodeEditorDialog.vue'
import { fetchSolverDockerfile, createContainer } from '@/middlewares/api.mdw'
import Vue from "vue";
export default Vue.extend({
  data:()=>({
    code:"",
    showCodeEditorDialog:false,
    showBuildDialog: false,
  }),
  props: ["images", "solverName", "packageName", "vendorName"],
  components:{
    'code-editor-dialog': CodeEditorDialog,
    'build-dialog': BuildDialog,
  },
  methods:{
    showCodeEditor(vendorName:string, packageName:string) {
      let self = this
      console.log(this.solverName)
      fetchSolverDockerfile(vendorName,packageName, this.solverName).then(function(res:any){
        self.code = res.msg
      })
      this.showCodeEditorDialog = true
    },
    newContainer(imageId: string) {
      createContainer(imageId).then(function(res){
        console.log(res)
        location.reload()
      })
    },
    triggerBuild() {
      this.showBuildDialog = true;
    }
  }
});
</script>
<style scoped>
.solver_image_box {
  text-align: center;
  color: black!important;
}
.solver_image_create_box {
  align-items: center;
  border: 1px;
  border-style: dotted;
  border-color: gray;
}
</style>