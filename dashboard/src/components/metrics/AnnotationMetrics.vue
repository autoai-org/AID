<template>
  <v-card>
    <v-card-title
      class="headline grey lighten-2"
      primary-title
    >
      Metrics of Dataset
    </v-card-title>

    <v-card-text>
      <v-btn
        flat
        small
        @click="drawClassPieChart()"
      >
        Pie Chart (train/test/val)
      </v-btn>
      <div id="classPieChart" />
    </v-card-text>

    <v-divider />

    <v-card-actions />
  </v-card>
</template>

<script>
import Plotly from 'plotly.js-dist'
export default {
  props: {
    annodata: {
      type: Array,
      default: function () {
        return []
      }
    }
  },
  data () {
    return {
      trainNum: 0,
      testNum: 0,
      valNum: 0,
      // number of images in each class
      classNum: {},
      // the total count of existing classes
      classesNum: 0,
      classesName: {}
    }
  },
  watch: {
    annodata (newVal) {
      this.parseData()
    }
  },
  methods: {
    parseData () {
      const self = this
      this.annodata.map(function (res) {
        if (typeof self.classNum[res.class.label.toString()] === 'undefined') {
          self.classNum[res.class.label.toString()] = 0
        }
        self.classNum[res.class.label.toString()] =
          self.classNum[res.class.label.toString()] + 1
        if (res.folder === 'test') {
          self.testNum = self.testNum + 1
        } else if (res.folder === 'train') {
          self.trainNum = self.trainNum + 1
        } else if (res.folder === 'val') {
          self.valNum = self.valNum + 1
        }
      })
    },
    drawClassPieChart () {
      const data = [
        {
          values: [this.trainNum, this.testNum, this.valNum],
          labels: ['Train', 'Test', 'Validation'],
          type: 'pie'
        }
      ]
      const layout = {
        height: 400,
        width: 500
      }

      Plotly.newPlot('classPieChart', data, layout)
    }
  }
}
</script>

<style>
</style>
