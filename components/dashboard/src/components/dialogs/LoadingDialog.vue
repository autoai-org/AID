<template>
<v-dialog v-model="show" hide-overlay persistent width="300">
      <v-card color="primary" dark>
        <v-card-text>
          Please stand by...
          <v-progress-linear indeterminate color="white" class="mb-0"></v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>
</template>
<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    props:['show']
})
</script>
<style scoped>

</style>