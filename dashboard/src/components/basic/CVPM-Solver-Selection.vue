<template>
  <v-layout class="cvpm-solver-runner-selection">
    <v-select
      v-model="selectedVendor"
      :items="vendor"
      label="Vendor"
    />
    <v-select
      v-model="selectedPackage"
      :items="packageName"
      label="Package"
    />
    <v-flex xs6>
      <v-select
        v-model="selectedSolver"
        :items="config.solvers"
        item-text="name"
        label="Solver"
      />
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  props: {
    config: {
      type: Object,
      default: function () {
        return {}
      }
    },
    vendor: {
      type: String,
      default: ''
    },
    packageName: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      selectedVendor: this.vendor[0],
      selectedPackage: this.packageName[0],
      selectedSolver: ''
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
  watch: {
    hasFinishSelection: function (val) {
      if (val) {
        this.finishSelection()
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
