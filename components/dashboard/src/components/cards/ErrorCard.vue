<template>
  <v-card width="300px" class="aid-errors">
    <v-card-title>Errors</v-card-title>
    <v-card-text>
        <div v-if="errors.length===0">
            <v-icon color="green" size="60px">mdi-check-decagram</v-icon> <br/>
            Congratulations, No Errors has occurred.
        </div>
        <div v-else>
            <v-icon color="red" size="60px">mdi-alert-decagram</v-icon> <br/>
            <ul class="error_list">
                <li v-for="(error, index) in errors" :key="index">[{{error.OccuredAt}}] {{error.Msg}}</li>
            </ul>
        </div>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
    props: ["errors"]
});
</script>
<style scoped>
.aid-errors {
  justify-content: center;
  text-align: center;
}
.error_list {
    list-style-type:initial;
    text-align: left;
}
</style>