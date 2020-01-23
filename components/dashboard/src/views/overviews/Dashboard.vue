<template>
  <v-container>
    <v-row>
      <v-col>
        <stat-card :title="'Installed Packages'" :content="'1'" />
      </v-col>
      <v-col>
        <stat-card :title="'Installed Packages'" :content="'1'" />
      </v-col>
      <v-col>
        <stat-card :title="'Installed Packages'" :content="'1'" />
      </v-col>
      <v-col>
        <stat-card :title="'Installed Packages'" :content="'1'" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <news-card/>
      </v-col>
      <v-col></v-col>
      <v-col></v-col>
      <v-col></v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import StatCard from "@/components/cards/StatCard.vue";
import NewsCard from "@/components/cards/NewsCard.vue";
import Vue from "vue";
export default Vue.extend({
  components: {
    "stat-card": StatCard,
    "news-card": NewsCard
  },
  mounted() {}
});
</script>

<style scoped>
</style>