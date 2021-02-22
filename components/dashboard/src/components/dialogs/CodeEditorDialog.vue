<template>
  <v-dialog v-model="dialog" width="60%" @click:outside="closeOutside()">
    <v-card>
      <v-card-title class="headline grey lighten-2" primary-title>Dockerfile</v-card-title>
        <v-card-text>
        <codemirror
          ref="myCm"
          :value="code"
          :options="cmOptions"
          @input="onCmCodeChange"
        ></codemirror>
        </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" text @click="closeOutside(); dialog=false">OK</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { codemirror } from "vue-codemirror";
import "codemirror/lib/codemirror.css";
import "codemirror/mode/dockerfile/dockerfile"
import Vue from "vue";
export default Vue.extend({
  data: () => ({
    dialog: false,
    cmOptions: {
      // codemirror options
      tabSize: 4,
      mode: "text/x-dockerfile",
      // lineNumbers: true,
      line: true
    }
  }),
  components: {
    codemirror
  },
  props: ["show", "title", "code"],
  watch: {
    show(newVal) {
      if (newVal) {
        this.dialog = true;
      }
    }
  },
  methods: {
    onCmCodeChange(newCode: string) {
      console.log("this is new code", newCode);
    },
    closeOutside() {
      this.$emit("closed", true);
    },
  },
});
</script>

<style>
.CodeMirror {
  border: 1px solid #eee;
  height: auto;
}
.CodeMirror pre { 
    white-space: pre-wrap; 
    word-break: break-all; 
    word-wrap: break-word; 
}
</style>