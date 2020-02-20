<template>
  <v-form
    ref="trainForm"
    v-model="valid"
    lazy-validation
  >
    <v-alert
      v-model="info"
      outline
      type="info"
    >
      You might be interested in using our hosted
      <a
        href="https://cvflow.autoai.org/#/guide/misc/hosted"
      >CVFlow</a>.
    </v-alert>
    <v-select
      v-model="selectedVendor"
      :items="vendors"
      label="Vendors"
    />
    <v-select
      v-if="selectedVendor"
      v-model="selectedPackage"
      :items="packages"
      label="Packages"
    />
    <v-select
      v-if="selectedPackage"
      v-model="selectedSolver"
      :items="solvers"
      label="Solvers"
    />
    <v-select
      v-if="selectedSolver"
      v-model="selectedDataset"
      :items="datasets"
      label="Datasets"
    />
    <v-btn
      outline
      color="indigo"
    >
      Start
    </v-btn>
  </v-form>
</template>

<script>
import { systemService } from '@/services/system'
import toml from 'toml'

export default {
  data () {
    return {
      datasets: [],
      selectedDataset: '',
      vendors: [],
      selectedVendor: '',
      packages: [],
      selectedPackage: '',
      solvers: [],
      selectedSolver: '',
      info: true,
      valid: true
    }
  },
  watch: {
    selectedVendor (val) {
      this.getPackages(val)
    },
    selectedPackage (val) {
      this.getSolvers(this.selectedVendor, val)
    },
    selectedSolver (val) {
      this.getDatasets()
    }
  },
  created () {
    this.getVendors()
  },
  methods: {
    getVendors () {
      const self = this
      systemService.getPackages().then(function (res) {
        res.data.map(function (each) {
          if (self.vendors.indexOf(each.Vendor) === -1) {
            self.vendors.push(each.Vendor)
          }
        })
      })
    },
    getPackages (vendorName) {
      const self = this
      systemService.getPackages().then(function (res) {
        res.data.map(function (each) {
          if (
            self.packages.indexOf(each.Name) === -1 &&
            each.Vendor === vendorName
          ) {
            self.packages.push(each.Name)
          }
        })
      })
    },
    getSolvers (vendorName, packageName) {
      const self = this
      systemService.getRepoMeta(vendorName, packageName).then(function (res) {
        const parsedConfig = toml.parse(res.data.Config)
        parsedConfig.solvers.map(function (each) {
          self.solvers.push(each.name)
        })
      })
    },
    getDatasets () {
      const self = this
      systemService.getMyFiles().then(function (res) {
        console.log(res.data)
        res.data.result.map(function (each) {
          if (
            self.datasets.indexOf(each.filepath) === -1 &&
            each.type === 'dataset'
          ) {
            console.log(each.filepath)
            self.datasets.push(each.filepath)
          }
        })
        console.log(self.datasets)
      })
    },
    getTrainConfig () {}
  }
}
</script>

<style>
</style>
