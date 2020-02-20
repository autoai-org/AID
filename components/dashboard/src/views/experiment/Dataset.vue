<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-card-title>Dataset</v-card-title>
    <v-card-text>
      <v-btn outlined @click="creating=true" v-if="!creating">New Dataset</v-btn>
      <v-btn outlined @click="creating=false" v-if="creating">Close</v-btn>
      <div class="creating-dataset-form" v-if="creating">
        <v-row>
          <v-col class="mb8">
            <v-text-field
              v-model="dataset_name"
              placeholder="Dataset Name"
              single-line
              append-icon="mdi-database"
            />
          </v-col>
          <v-col class="mb2"></v-col>
        </v-row>
        <v-row>
          <v-col class="mb8">
            <v-text-field
              v-model="dataset_path"
              placeholder="Dataset Path"
              single-line
              append-icon="mdi-database"
            />
          </v-col>
          <v-col class="mb2">
            <v-btn @click="trigger_upload">Upload New File</v-btn>
          </v-col>
        </v-row>
        <v-row v-if="upload_input">
          <v-col class="mb8">
            <v-file-input v-model="fileBlob" ref="file-input" show-size label="Upload File"></v-file-input>
          </v-col>
          <v-col>
            <v-btn color="primary" @click="confirmUpload">Upload</v-btn>
          </v-col>
        </v-row>
        <v-row>
          <v-col class="mb2">
            <v-switch v-model="uncompress" class="ma-2" label="Uncompress Now?"></v-switch>
          </v-col>
        </v-row>
        <v-row>
          <v-btn color="primary" outlined @click="addNewDataset()">Confirm</v-btn>
        </v-row>
      </div>
      <enhanced-table class="dataset-table" :headers="headers" :items="datasets"></enhanced-table>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import EnhancedTable from "@/components/EnhancedTable.vue"
import { uploadDataset, addDataset, getDatasets } from "@/middlewares/api.mdw";
export default Vue.extend({
  data() {
    return {
      creating: false,
      uncompress: false,
      dataset_path: "",
      dataset_name:"",
      fileBlob: new Blob(),
      upload_input: false,
      datasets:[],
      headers: [
        { text: "ID", value: "ID" },
        { text: "Name", value: "Name" },
        { text: "LocalPath", value: "LocalPath" },
        { text: "Size", value: "Size" },
        { text: "Status", value: "Status" },
        { text: "Created At", value: "CreatedAt" },
      ],
    };
  },
  components: {
    'enhanced-table' : EnhancedTable
  },
  methods: {
    addNewDataset() {
      addDataset(this.dataset_path, this.uncompress, this.dataset_name).then(function(res:any){
        console.log(res)
      })
    },
    trigger_upload() {
      this.upload_input = true;
    },
    confirmUpload() {
      let self = this
      if (this.fileBlob.size!==0) {
        uploadDataset(this.fileBlob).then(function(res:any){
            self.dataset_path = res
        });
      }
    }
  },
  mounted () {
    let self = this
    getDatasets().then(function(res:any){
      self.datasets = res.map(function(each:any){
        each.Size = each.Size.toFixed(2) +" MB"
        return each
      })
    })
  }
});
</script>
<style scoped>
.creating-dataset-form {
  margin-left: 20px;
}
.dataset-table{
  margin-top: 20px;
}
</style>