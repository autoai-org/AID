<template>
  <v-layout class="cvpm-solver-runner-selection">
    <v-select v-model="selectedVendor" :items="vendor" label="Vendor"></v-select>
    <v-select v-model="selectedPackage" :items="packageName" label="Package"></v-select>
    <v-flex xs6>
      <v-select
        v-model="selectedSolver"
        :items="config.solvers"
        item-text="name"
        label="Solver"
      ></v-select>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  data () {
    return {
      selectedVendor: this.vendor[0],
      selectedPackage: this.packageName[0],
      selectedSolver: ''
    }
  },
  props: ['config', 'vendor', 'packageName'],
  watch: {
    hasFinishSelection: function (val) {
      if (val) {
        this.finishSelection()
      }
    }
  },
  computed: {
    hasFinishSelection () {
      if (this.selectedVendor && this.selectedPackage && this.selectedSolver) {
        return true
      } else {
        return false
      }
    }
  },
  methods: {
    finishSelection () {
      this.$emit('finishSelection', {
        selectedVendor: this.selectedVendor,
        selectedPackage: this.selectedPackage,
        selectedSolver: this.selectedSolver
      })
    }
  }
}
</script>

<style scoped>
.cvpm-solver-runner-selection {
  margin-left: auto;
  margin-right: auto;
  width: 95%;
}
</style>
