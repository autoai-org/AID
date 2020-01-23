<template>
  <v-card class="mx-auto" max-width="344">
    <v-card-title>What's New?</v-card-title>
    <v-card-actions>
        <v-btn color="primary" class="ak-trigger">Explore</v-btn>
        <v-btn color="secondary" class="ak-trigger" href="https://twitter.com/aid_aiops"><v-icon>mdi-twitter</v-icon></v-btn>
    </v-card-actions>
    <announce-kit catchClick=".ak-trigger" widget="https://announcekit.app/widget/1m2DKg" v-show="false"/>
  </v-card>
</template>
<script lang="ts">
import AnnounceKit from '@/components/external/AnnounceKit.vue'
import Vue from "vue";
export default Vue.extend({
    components:{
        'announce-kit': AnnounceKit
    }
});
</script>

<style scoped>
</style>