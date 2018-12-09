<template>
  <div>
    <h4>Paramters required by CVPM</h4>
    <v-text-field
      v-if="typeof(item.value) === 'string'"
      v-for="(item, index) in systemRequiredParams"
      :key="index"
      :label="item.key"
      required
      v-model="item.value"
      :hint="item.hint"
    ></v-text-field>
    <v-checkbox
      v-for="(item, index) in systemRequiredParams"
      :key="index"
      v-if="typeof(item.value) === 'boolean'"
      :label="item.key"
      v-model="item.value"
      :hint="item.hint"
    ></v-checkbox>
    <v-divider></v-divider>
    <h4>Parameters required by Solver</h4>
    <v-btn color="indigo darken-1" outline @click="addParams()">
      <v-icon>add</v-icon>
    </v-btn>
    <v-layout>
      <v-flex xs6>
        <v-text-field
          v-for="(item, index) in solverDefinedParams"
          :key="index"
          label="Key"
          v-model="item.key"
          hint
        ></v-text-field>
      </v-flex>
      <v-flex xs6>
        <v-text-field
          v-for="(item, index) in solverDefinedParams"
          :key="index"
          label="Value"
          v-model="item.value"
          hint
        ></v-text-field>
      </v-flex>
    </v-layout>
    <v-btn color="indigo darken-1" outline @click="startTest()" :loading="loading">Start!</v-btn>
  </div>
</template>

<script>
import { systemService } from '@/services/system'
export default {
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
  props: ['file'],
  methods: {
    addParams () {
      this.solverDefinedParams.push({
        key: '',
        value: ''
      })
    },
    startTest () {
      let self = this
      self.loading = true
      const requestParams = this.systemRequiredParams.concat(this.solverDefinedParams)
      systemService
        .testRepoSolver('8080', requestParams, this.file)
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
</script>

<style>
</style>
