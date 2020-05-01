<template>
  <div>
    <v-expansion-panels flat v-for="(item, index) in solvers" :key="index">
      <v-expansion-panel flat>
        <v-expansion-panel-header>{{item.Name}}</v-expansion-panel-header>
        <v-expansion-panel-content>
          <v-stepper v-model="e1">
            <v-stepper-header>
              <v-stepper-step editable :complete="e1 > 1" step="1">Build</v-stepper-step>
              <v-divider></v-divider>
              <v-stepper-step editable :complete="e1 > 2" step="2">Start</v-stepper-step>
              <v-divider></v-divider>
              <v-stepper-step editable :complete="e1 > 3" step="3">Test</v-stepper-step>
            </v-stepper-header>
            <v-stepper-items>
                <v-stepper-content step="1">
                    <build-step :images="images" :solverName="item.Name" :vendorName="vendor" :packageName="packageName" />
                </v-stepper-content>
                <v-stepper-content step="2">
                  <start-step :containers="containers" :solverName="item.Name"></start-step>
                </v-stepper-content>
                <v-stepper-content step="3">
                  <test-step></test-step>
                </v-stepper-content>
            </v-stepper-items>
          </v-stepper>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </div>
</template>

<script lang="ts">
import BuildStep from "@/components/packageComponents/BuildStep.vue"
import TestStep from "@/components/packageComponents/TestStep.vue"
import StartStep from "@/components/packageComponents/StartStep.vue"
import { fetchImages, fetchContainers } from '@/middlewares/api.mdw'
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      e1: 1,
      images:[],
      containers:[],
    };
  },
  props: ["solvers", "vendor", "packageName"],
  mounted () {
    let self = this
    fetchImages().then(function(res:any){
        console.log(res)
        self.images = res.map(function(each:any){
            each.solverName = each.Name.split("-").slice(-1)[0]
            each.number = each.Name.split("-")[1]
            each.vendor = each.Name.split("-")[2]
            each.package = each.Name.split("-")[3]
            return each
        }).filter(function(each:any){
            if (each.package == self.packageName) {
                return true
            } else {
                return false
            }
        })
        console.log(self.images)
    })
    fetchContainers().then(function(res:any){
      self.containers = res
    })
  },
  components:{
      'build-step': BuildStep,
      'start-step': StartStep,
      'test-step': TestStep
  }
});
</script>

<style scoped>
.solvers-text {
  color: black !important;
}
</style>