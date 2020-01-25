<template>
  <v-card>
    <v-card-text>
      <v-row class="mb-12">
        <v-col cols="2" v-for="(item,index) in images" :key="index" class="d-flex flex-column justify-center solver_image_box">
          <v-icon size="64pt">mdi-package-variant-closed</v-icon>
            {{item.number}}-{{item.solverName}} <br/>
            {{item.CreatedAt}}
            <v-btn color="primary">Create Container</v-btn> <br/>
            <v-btn color="secondary" outlined>Delete Image</v-btn>
        </v-col>
        <v-col cols="2" class="d-flex flex-column justify-center solver_image_box solver_image_create_box">
          <v-icon size="64pt">mdi-plus</v-icon>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  props: ["images"]
});
</script>
<style scoped>
.solver_image_box {
  text-align: center;
  color: black!important;
}
.solver_image_create_box {
  border: 1px;
  border-style: dotted;
  border-color: gray;
}
</style>