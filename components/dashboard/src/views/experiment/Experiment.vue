<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-card-title>Experiment</v-card-title>
    <v-card-text>
      <v-btn outlined color="primary">
        <v-icon>mdi-history</v-icon>History
      </v-btn>
      <v-stepper v-model="e1">
        <v-stepper-header>
          <v-stepper-step editable :complete="e1 > 1" step="1">Choose Dataset</v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step editable :complete="e1 > 2" step="2">Choose Model</v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step editable :complete="e1 > 3" step="3">Confirmation</v-stepper-step>
          <v-divider></v-divider>
          <v-stepper-step editable :complete="e1 > 4" step="4">Progress and Results</v-stepper-step>
        </v-stepper-header>
        <v-stepper-items>
          <v-stepper-content step="1">
            <v-select :items="datasets" :item-text="'Name'" :item-value="'ID'" label="Dataset" v-model="selected_dataset_id"></v-select>
            <p v-if="selected_dataset_id">
                You have selected <B>{{selected_dataset.Name}}</B>, with the ID <B>{{selected_dataset.ID}}</B>.<br/>
                The size of your selected dataset is <B>{{selected_dataset.Size.toFixed(2)}} MB</B>.<br/>
                DeepAnalytics and Augmentor is disabled for this dataset.
            </p>
            <v-btn color="primary" @click="nextStep()">Confirm</v-btn>
          </v-stepper-content>
          <v-stepper-content step="2">
              
          </v-stepper-content>
          <v-stepper-content step="3"></v-stepper-content>
          <v-stepper-content step="4"></v-stepper-content>
        </v-stepper-items>
      </v-stepper>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { getDatasets } from "@/middlewares/api.mdw";
export default Vue.extend({
  data() {
    return {
      e1: 1,
      selected_dataset_id:'',
      selected_dataset: {},
      datasets: [] as Array<any>,
    };
  },
  mounted() {
    let self = this;
    getDatasets().then(function(res: any) {
      self.datasets = res;
    });
  },
  methods: {
      nextStep() {
          this.e1 = this.e1+1
      }
  },
  watch: {
      selected_dataset_id(val) {
          for(let item of this.datasets){
              if (item.ID===val) {
                  this.selected_dataset = item
              }
          }
      }
  }
});
</script>
<style scoped>
</style>