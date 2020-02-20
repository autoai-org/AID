<template>
  <v-tabs>
    <v-tab ripple>
      Full Log
    </v-tab>
    <v-tab ripple>
      Less Log
    </v-tab>
    <v-tab ripple>
      Line Chart
    </v-tab>
    <v-tab-item>
      <cvpm-log
        :title="'Full Log for Train Task ' + rayId"
        :message="fullMessage"
      />
    </v-tab-item>
    <v-tab-item>
      <cvpm-log
        :title="'Log for Train Task ' + rayId"
        :message="lessMessage"
      />
    </v-tab-item>
    <v-tab-item>
      <div id="trainLineChart" />
    </v-tab-item>
  </v-tabs>
</template>

<script>
import Log from '@/components/CVPM-Log'
import { systemService } from '@/services/system'
import Plotly from 'plotly.js-dist'
import store from '@/store'
export default {
  components: {
    'cvpm-log': Log
  },
  data () {
    return {
      rayId: '',
      fullMessage: '',
      lessMessage: ''
    }
  },
  created () {
    this.initSocket()
  },
  methods: {
    initSocket () {
      store.state.isLoading = true
      this.rayId = this.$route.params.rayId
      const fullConn = new WebSocket(
        systemService.getWsEndpoint(this.rayId + '.full')
      )
      const lessConn = new WebSocket(
        systemService.getWsEndpoint(this.rayId + '.less')
      )
      const self = this
      fullConn.onmessage = function (evt) {
        self.fullMessage = evt.data
        store.state.isLoading = false
      }
      lessConn.onmessage = function (evt) {
        self.lessMessage = evt.data
        self.parsedLessMessage = self.parseLessMessage(self.lessMessage)
        self.drawLineChart(self.parsedLessMessage)
        store.state.isLoading = false
      }
    },
    parseLessMessage (lessLogMessage) {
      const lines = lessLogMessage.split('\n')
      lines.splice(0, 1)
      return JSON.parse(lines.join('\n'))
    },
    drawLineChart (lessLogMessage) {
      const x = lessLogMessage.map(function (each) {
        return Number(each.global_step)
      })
      const y = lessLogMessage.map(function (each) {
        return Number(each.loss)
      })
      const data = {
        x: x,
        y: y,
        type: 'scatter'
      }
      const plotData = [data]
      console.log(y)
      Plotly.newPlot('trainLineChart', plotData)
    }
  }
}
</script>

<style>
</style>
