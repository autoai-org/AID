<template>
  <v-card>
    <v-card-title>Test Inference</v-card-title>
    <v-card-text>
      <v-file-input v-model="fileBlob" show-size label="Upload File"></v-file-input>
      <canvas id="overlay" class="coveringCanvas"></canvas>
    </v-card-text>
    <v-card-actions>
      <v-btn color="primary" @click="call">Call</v-btn>
    </v-card-actions>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { drawBBox } from "@/middlewares/drawer/bbox";
import { testContainer } from "@/middlewares/api.mdw";
export default Vue.extend({
  data() {
    return {
      fileBlob: new Blob(),
      imgSrc: ""
    };
  },
  methods: {
    call() {
      let self = this;
      this.imgSrc = URL.createObjectURL(this.fileBlob);
      let canvas = document.getElementById("overlay") as HTMLCanvasElement;
      let ctx = canvas.getContext("2d");
      let backgroundImage = new Image();
      backgroundImage.src = this.imgSrc
      backgroundImage.onload = function() {
        if(ctx != null) {
            ctx.drawImage(backgroundImage,0,0, backgroundImage.width, backgroundImage.height,0,0, canvas.width, canvas.height)
        }
      };
      
      testContainer(this.fileBlob).then(function(res) {
        if (ctx != null) {
          drawBBox(ctx, res, backgroundImage.width, backgroundImage.height,canvas.width, canvas.height);
        }
      });
    }
  }
});
</script>

<style scoped>
.coveringCanvas {
    width: 100%;
    height: 100%;
}
</style>