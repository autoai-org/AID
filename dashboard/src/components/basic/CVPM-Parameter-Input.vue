<template>
  <div>
    <h4>Paramters required by CVPM</h4>
    <v-text-field
      v-for="(item, index) in stringParameters"
      :key="index"
      v-model="item.value"
      :label="item.key"
      required
      :hint="item.hint"
    />
    <v-checkbox
      v-for="(item, index) in booleanParamters"
      :key="index"
      v-model="item.value"
      :label="item.key"
      :hint="item.hint"
    />
    <v-divider />
    <h4>Parameters required by Solver</h4>
    <v-btn
      color="indigo darken-1"
      outline
      @click="addParams()"
    >
      <v-icon>add</v-icon>
    </v-btn>
    <v-layout>
      <v-flex xs6>
        <v-text-field
          v-for="(item, index) in solverDefinedParams"
          :key="index"
          v-model="item.key"
          label="Key"
          hint
        />
      </v-flex>
      <v-flex xs6>
        <v-text-field
          v-for="(item, index) in solverDefinedParams"
          :key="index"
          v-model="item.value"
          label="Value"
          hint
        />
      </v-flex>
    </v-layout>
    <v-btn
      color="indigo darken-1"
      outline
      :loading="loading"
      @click="startTest()"
    >
      Start!
    </v-btn>
  </div>
</template>

<script>
import { systemService } from '@/services/system'
export default {
  props: {
    file: {
      type: File,
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
    },
    solverName: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      loading: false,
      systemRequiredParams: [
        {
          key: 'delete_after_process',
          value: true,
          hint: 'Should the service delete the image after process?'
        }
      ],
      solverDefinedParams: []
    }
  },
  computed: {
    stringParameters () {
      return this.systemRequiredParams.filter(
        each => typeof each.value === 'string'
      )
    },
    booleanParamters () {
      return this.systemRequiredParams.filter(
        each => typeof each.value === 'boolean'
      )
    }
  },
  methods: {
    addParams () {
      this.solverDefinedParams.push({
        key: '',
        value: ''
      })
    },
    startTest () {
      const self = this
      self.loading = true
      const requestParams = this.systemRequiredParams.concat(
        this.solverDefinedParams
      )
      if (
        typeof this.selectedSolver === 'undefined' ||
        typeof this.selectedVendor === 'undefined' ||
        typeof this.selectedPackage === 'undefined'
      ) {
        alert('Select Solver/Vendor/Package First!')
      } else {
        systemService
          .testRepoSolver(
            this.vendor,
            this.packageName,
            this.solverName,
            requestParams,
            this.file
          )
          .then(function (res) {
            self.loading = false
            self.$emit('finishInfer', res)
          })
          .catch(function (err) {
            alert('an error occured' + err)
          })
      }
    }
  }
}
</script>

<style>
</style>
