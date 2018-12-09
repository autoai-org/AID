<template>
  <v-card>
    <v-card-title>
      <h2>Test</h2>
    </v-card-title>
    <cvpm-solver-selection
      :config="config"
      :vendor="vendor"
      :packageName="packageName"
      v-on:finishSelection="onFinishSelection"
    ></cvpm-solver-selection>
    <v-card class="cvpm-test-steps">
      <v-stepper v-model="step">
        <v-stepper-header>
          <v-stepper-step :complete="step > 1" step="1">Choose Image</v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step :complete="step > 2" step="2">Complete Parameters</v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step step="3">Check Results</v-stepper-step>
        </v-stepper-header>

        <v-stepper-items class="cvpm-test-content">
          <v-stepper-content step="1">
            <cvpm-image-upload v-on:fileSelected="onFileSelected"></cvpm-image-upload>
          </v-stepper-content>
          <v-stepper-content step="2">
            <cvpm-parameter-input :file=file v-on:finishInfer="onFinishInfer"></cvpm-parameter-input>
          </v-stepper-content>
          <v-stepper-content step="3">
            <cvpm-json-view :jsonObject="inferResponse"></cvpm-json-view>
          </v-stepper-content>
        </v-stepper-items>

      </v-stepper>
    </v-card>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="indigo darken-1" outline @click="closeDialog()">Close</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import cvpmSolverSelection from '@/components/basic/CVPM-Solver-Selection'
import cvpmImageUpload from '@/components/basic/CVPM-Image-Upload'
import cvpmParameterInput from '@/components/basic/CVPM-Parameter-Input'
import cvpmJSONView from '@/components/basic/CVPM-JSON-View'
export default {
  data () {
    return {
      selectedPackage: '',
      selectedVendor: '',
      selectedSolver: '',
      step: 1,
      file: null,
      inferResponse: {}
    }
  },
  methods: {
    closeDialog () {
      this.$emit('closeDialog', true)
    },
    onFinishSelection (value) {
      this.selectedSolver = value.selectedSolver
      this.selectedVendor = value.selectedVendor
      this.selectedPackage = value.selectedPackage
    },
    onFileSelected (value) {
      this.step = 2
      this.file = value
    },
    onFinishInfer (value) {
      this.inferResponse = value.data
      this.step = 3
    }
  },
  components: {
    'cvpm-solver-selection': cvpmSolverSelection,
    'cvpm-image-upload': cvpmImageUpload,
    'cvpm-parameter-input': cvpmParameterInput,
    'cvpm-json-view': cvpmJSONView
  },
  props: ['config', 'vendor', 'packageName']
}
</script>

<style scoped>
.cvpm-test-content {
  align-items: center;
  justify-content: center;
  text-align: center;
}
.cvpm-test-steps {
  width: 95%;
  margin-left: auto;
  margin-right: auto;
}
</style>
