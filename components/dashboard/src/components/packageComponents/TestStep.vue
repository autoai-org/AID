<template>
<div>
  <v-card>
    <v-card-title>Test Inference</v-card-title>
    <v-card-text>
      <v-file-input v-model="fileBlob" show-size label="Upload File"></v-file-input>
      <!--<canvas id="overlay" class="coveringCanvas"></canvas>-->
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" @click="call">Call</v-btn>
    </v-card-actions>
  </v-card>
  <v-card v-if="requestCompleted">
    <v-card-title>Results</v-card-title>
    <v-card-text>
      <p>The request costs <B>{{(endTime-beginTime).toFixed(2)}}</B> ms. <a href="https://nervanasystems.github.io/distiller/quantization.html">Optimize?</a></p>
      <json-viewer :value="jsonResponse" copyable boxed sort></json-viewer>
    </v-card-text>
  </v-card>
</div>
</template>
<script lang="ts">
import Vue from "vue";
import { drawBBox } from "@/middlewares/drawer/bbox";
import { testContainer } from "@/middlewares/api.mdw";
import JsonViewer from 'vue-json-viewer'
export default Vue.extend({
  data() {
    return {
      fileBlob: new Blob(),
      requestCompleted: false,
      jsonResponse:{},
      beginTime: 0,
      endTime:0,
    };
  },
  methods: {
    call() {
      let self = this;
      this.beginTime =new Date().getTime()
      testContainer(this.fileBlob).then(function(res:any) {
        console.log(res)
        self.endTime = new Date().getTime()
        self.requestCompleted = true
        self.jsonResponse = res
      });
    }
  },
  components:{
    'json-viewer': JsonViewer
  }
});
</script>

<style scoped>
.coveringCanvas {
    width: 100%;
    height: 100%;
}
</style>