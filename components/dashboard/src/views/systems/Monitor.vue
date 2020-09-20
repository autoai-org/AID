<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-list>
      <v-card-title>System Monitor</v-card-title>
      <v-row>
        <v-col class="d-flex" cols="12" sm="6">
            <v-spacer></v-spacer>
          <v-select v-model="selectedPeriod" :items="periods" label="Update Period"></v-select>
        </v-col>
      </v-row>
      <div id="requestsMonitor" />
    </v-list>
  </v-card>
</template>
<script>
import { fetchPrometheus } from "../../middlewares/api.mdw";
import Plotly from "plotly.js-dist";
export default {
  data: () => ({
    selectedPeriod: "2s",
    metrics: [],
    periods: ["2s", "5s", "10s"],
    successCount: [],
    errorCount: [],
  }),
  mounted() {
    this.periodicallyUpdate();
  },
  methods: {
    periodicallyUpdate() {
      let milesecPeriod = Number(this.selectedPeriod.split("s")[0]) * 1000;
      setInterval(this.queryPrometheus, milesecPeriod);
    },
    prepareRequestData() {
      for (const each of this.metrics) {
        if (each.name === "promhttp_metric_handler_requests_total") {
          let total_failed = 0
          for (const each_req of each.metrics) {
            if (each_req.labels.code === "200") {
              this.successCount.push(each_req.value);
            } else {
              total_failed += each_req.value;
            }
          }
          this.errorCount.push(total_failed)
        }
      }
      this.drawLineChart();
    },
    drawLineChart() {
      let success = this.successCount;
      let error = this.errorCount;
      let successTrace = {
        x: Array.from({ length: success.length }, (v, k) => k + 1),
        y: success,
        name: "Successful Requests",
      };
      let errorTrace = {
        x: Array.from({ length: error.length }, (v, k) => k + 1),
        y: error,
        name: "Failed Requests",
      };
      let data = [successTrace, errorTrace];

      let layout = {
        title: "Requests Success Rate",
      };

      Plotly.newPlot("requestsMonitor", data, layout);
    },
    queryPrometheus() {
      console.log('querying...')
      let self = this;
      fetchPrometheus().then(function (res) {
        self.metrics = res;
        self.prepareRequestData();
      });
    },
  },
};
</script>
<style scoped>
</style>