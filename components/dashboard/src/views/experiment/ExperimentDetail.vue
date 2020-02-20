<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-card-title>Experiment #1</v-card-title>
    <v-card-text>
      <v-expansion-panels>
        <v-expansion-panel>
          <v-expansion-panel-header>Loss</v-expansion-panel-header>
          <v-expansion-panel-content eager>
            <div id="trainLineChartLoss" />
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header>Accuracy</v-expansion-panel-header>
          <v-expansion-panel-content eager>
            <div id="trainLineChartAcc" />
          </v-expansion-panel-content>
        </v-expansion-panel>
        <v-expansion-panel>
          <v-expansion-panel-header>Power Consumption</v-expansion-panel-header>
          <v-expansion-panel-content eager>
            <div id="powerConsumptionChart" />
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import Plotly from "plotly.js-dist";
export default Vue.extend({
  methods: {
    drawLineChart() {
      const x = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
      const y = [2.3, 1.9, 1.8, 1.5, 1.6, 1.5, 1.6, 1.3, 1.2, 1.0];
      const data = {
        x: x,
        y: y,
        type: "scatter"
      };
      const plotData = [data];
      Plotly.newPlot("trainLineChartLoss", plotData);
      Plotly.newPlot("trainLineChartAcc", plotData);
      Plotly.newPlot("powerConsumptionChart", plotData);
    }
  },
  mounted() {
    this.drawLineChart();
  }
});
</script>

<style scoped>
</style>