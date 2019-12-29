<template>
  <v-card outlined min-width="100%">
    <EnhancedTable :headers="headers" :items="packages"/>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";

import EnhancedTable from "@/components/EnhancedTable.vue";
import { fetchAllPackages } from "@/middlewares/api.mdw";

export default Vue.extend({
  components: {
    EnhancedTable
  },
  data: () => ({
      headers: [
          { text: 'ID', value: 'ID' },
          { text: 'Vendor', value: 'Vendor' },
          { text: 'Status', value: 'Status' },
        ],
  }),
  mounted() {
    fetchAllPackages();
  },
  computed: {
    ...mapState({
      packages: "packages"
    })
  }
});
</script>

<style scoped>
</style>