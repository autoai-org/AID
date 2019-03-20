<template>
  <v-card>
    <v-card-title>
      <h2>Solvers</h2>
    </v-card-title>
    <cvpm-solver-selection
      :config="config"
      :vendor="vendor"
      :package-name="packageName"
      @finishSelection="onFinishSelection"
    />
    <v-flex class="cvpm-solver-run-btn">
      <v-btn
        color="indigo"
        outline
        @click="run()"
      >
        Run
      </v-btn>
    </v-flex>
    <v-expansion-panel>
      <v-expansion-panel-content>
        <div slot="header">
          Running
        </div>
        <v-list>
          <v-list-tile
            v-for="(solver, index) in runningSolvers"
            :key="index"
          >
            <v-list-tile-content>{{ solver.SolverName }}</v-list-tile-content>
            <v-list-tile-avatar>{{ solver.Port }}</v-list-tile-avatar>
          </v-list-tile>
        </v-list>
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-dialog
      v-model="runningConfirmDialog"
      persistent
      width="60%"
    >
      <v-card>
        <v-card-title>
          <span
            class="headline"
          >Running Solver {{ selectedVendor }}/{{ selectedPackage }}/{{ selectedSolver }}</span>
        </v-card-title>
        <v-card-text>
          <v-text-field
            v-model="runningPort"
            label="Running Port"
            :rules="[rules.validPort]"
            hint="Keep it empty if you want to run it on any one of open ports"
          />
          <p v-if="responseRunningPort">
            This solver is running on {{ responseRunningPort }}
          </p>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
          <v-btn
            color="indigo darken-1"
            flat
            outline
            @click="triggerDialog()"
          >
            Close
          </v-btn>
          <v-btn
            color="indigo darken-1"
            :loading="isRequesting"
            flat
            outline
            @click="confirmedRun()"
          >
            Run!
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-card>
</template>

<script>
import { systemService } from '@/services/system'
import cvpmSolverSelection from '@/components/basic/CVPM-Solver-Selection'
export default {
  components: {
    'cvpm-solver-selection': cvpmSolverSelection
  },
  props: {
    config: {
      type: Object,
      default: function () {
        return {}
      }
    },
    vendor: {
      type: Array,
      default: function () {
        return []
      }
    },
    packageName: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
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
  created () {
    this.fetchRunningSolver()
  },
  methods: {
    fetchRunningSolver () {
      let self = this
      systemService
        .getRunningSolver(this.vendor, this.packageName)
        .then(function (res) {
          self.runningSolvers = res.data
        })
    },
    onFinishSelection (value) {
      this.selectedSolver = value.selectedSolver
      this.selectedVendor = value.selectedVendor
      this.selectedPackage = value.selectedPackage
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
      if (
        this.selectedSolver === '' ||
        this.selectedVendor === '' ||
        this.selectedPackage === ''
      ) {
        alert('Select Solver/Vendor/Package First!')
      } else {
        systemService
          .runRepoSolver(
            this.selectedVendor,
            this.selectedPackage,
            this.selectedSolver,
            this.runningPort
          )
          .then(function (res) {
            self.isRequesting = false
            self.runningConfirmDialog = false
            self.responseRunningPort = res.data.port
            self.fetchRunningSolver()
          })
      }
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
.cvpm-solver-run-btn {
  margin-right: 10%;
}
</style>
