<template>
  <div>
    <v-btn color="indigo darken-1" outline @click="selectFile()">
      <v-icon>camera_alt</v-icon>Upload
    </v-btn>
    <input
      id="files"
      type="file"
      name="file"
      ref="uploadInput"
      accept="image/*"
      :multiple="false"
      @change="uploadFile($event)"
    >
  </div>
</template>

<script>
export default {
  methods: {
    data () {
      return {
      }
    },
    selectFile () {
      this.$refs.uploadInput.click()
    },
    uploadFile (e) {
      let files = e.target.files
      if (files.length) {
        this.$emit('fileSelected', files[0])
      } else {
      }
    }
  }
}
</script>

<style scoped>
input[type="file"] {
  position: absolute;
  clip: rect(0, 0, 0, 0);
}
</style>
