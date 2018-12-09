<template>
  <v-card>
    <v-card-title>
      <h2>Actions</h2>
    </v-card-title>
    <v-card-actions>
      <v-btn color="indigo darken-1" outline @click="triggerDialog('test')">Test</v-btn>
    </v-card-actions>
    <v-dialog v-model="enableTest">
      <cvpm-request 
        v-if="enableTest" 
        v-on:closeDialog="triggerDialog('test')"
        :config="config"
        :vendor="vendor"
        :packageName="packageName"
        ></cvpm-request>
    </v-dialog>
  </v-card>
</template>

<script>
import cvpmRequest from '@/components/CVPM-Request'
export default {
  data () {
    return {
      enableTest: false
    }
  },
  methods: {
    triggerDialog (name) {
      if (name === 'test') {
        this.enableTest = !this.enableTest
      } else {
      }
    }
  },
  components: {
    'cvpm-request': cvpmRequest
  },
  props: ['config', 'vendor', 'packageName']
}
</script>

<style>
</style>
