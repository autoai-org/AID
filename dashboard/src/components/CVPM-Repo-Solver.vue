<template>
  <v-card>
    <v-card-title>
      <h2>Solvers</h2>
    </v-card-title>
    <v-layout class="cvpm-solver-runner-selection">
      <v-select v-model="selectedVendor" :items="vendor" label="Vendor"></v-select>
      <v-select v-model="selectedPackage" :items="packageName" label="Package"></v-select>
      <v-flex xs6>
        <v-select
          v-model="selectedSolver"
          :items="configObject.solvers"
          item-text="name"
          label="Solver"
        ></v-select>
      </v-flex>
    </v-layout>
    <v-flex class="cvpm-solver-run-btn">
      <v-btn @click="run()" color="indigo" outline>Run</v-btn>
    </v-flex>
    <v-expansion-panel>
      <v-expansion-panel-content>
        <div slot="header">Running</div>
        <v-list>
          <v-list-tile v-for="(solver, index) in runningSolvers" :key="index">
            <v-list-tile-content>
                {{solver.SolverName}}
            </v-list-tile-content>
            <v-list-tile-avatar>
                {{solver.Port}}
            </v-list-tile-avatar>
          </v-list-tile>
        </v-list>
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-dialog persistent v-model="runningConfirmDialog" width="60%">
      <v-card>
        <v-card-title>
          <span
            class="headline"
          >Running Solver {{selectedVendor}}/{{selectedPackage}}/{{selectedSolver}}</span>
        </v-card-title>
        <v-card-text>
          <v-text-field
            label="Running Port"
            v-model="runningPort"
            :rules="[rules.validPort]"
            hint="Keep it empty if you want to run it on any one of open ports"
          ></v-text-field>
          <p v-if="responseRunningPort">This solver is running on {{responseRunningPort}}</p>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="indigo darken-1" flat @click="triggerDialog()" outline>Close</v-btn>
          <v-btn
            color="indigo darken-1"
            @loading="isRequesting"
            flat
            @click="confirmedRun()"
            outline
          >Run!</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-card>
</template>

<script>
import { systemService } from '@/services/system'
import toml from 'toml'
export default {
  data () {
    return {
      selectedSolver: '',
      selectedVendor: this.vendor[0],
      selectedPackage: this.packageName[0],
      runningConfirmDialog: false,
      runningPort: '',
      runningSolvers: [],
      responseRunningPort: '',
      isRequesting: false,
      rules: {
        validPort: value => {
          if (value === '') {
            return true
          }
          const pattern = /^\+?(0|[1-9]\d*)$/
          if (pattern.test(value)) {
            if (parseInt(value) > 1024 && parseInt(value) < 65535) {
              return true
            } else {
              return 'Not a Valid Port Number (Must be a postive integer larger than 1024 and less than 65535)'
            }
          } else {
            return 'Not a Valid Port Number (Only numeric value is accepted)'
          }
        }
      }
    }
  },
  props: ['config', 'vendor', 'packageName'],
  computed: {
    configObject () {
      return toml.parse(this.config)
    }
  },
  methods: {
    fetchRunningSolver () {
      let self = this
      systemService
        .getRunningSolver(this.vendor, this.packageName)
        .then(function (res) {
          self.runningSolvers = res.data
          console.log(self.runningSolvers)
        })
    },
    triggerDialog () {
      this.runningConfirmDialog = !this.runningConfirmDialog
    },
    run () {
      if (this.selectedSolver === '') {
        alert('no solver specified')
      } else {
        this.triggerDialog()
      }
    },
    confirmedRun () {
      this.isRequesting = true
      let self = this
      systemService
        .runRepoSolver(
          this.selectedVendor,
          this.selectedPackage,
          this.selectedSolver,
          this.runningPort
        )
        .then(function (res) {
          self.isRequesting = false
          self.responseRunningPort = res.data.port
        })
    }
  },
  created () {
    this.fetchRunningSolver()
  }
}
</script>

<style scoped>
.cvpm-solver-runner-selection {
  margin-left: auto;
  margin-right: auto;
  width: 95%;
}
.cvpm-solver-run-btn {
  margin-right: 10%;
}
</style>
