<template>
<v-container>
  <v-btn class="ma-2" outlined color="indigo" @click="showNewPackageDialog=true">New Package</v-btn>
  <new-package-dialog :show="showNewPackageDialog" @closed="showNewPackageDialog=false"/>
  <v-card outlined min-width="100%">
    <v-card-title>Packages</v-card-title>
    <EnhancedTable :headers="headers" :select="true" :items="packages" :actions="actions"/>
  </v-card>
  <v-spacer/>
  <v-card outlined min-width="100%" class="solvers">
    <v-card-title>Solvers</v-card-title>
    <EnhancedTable :headers="solverHeaders" :select="false" :items="solvers" :actions="solverActions"/>
  </v-card>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";
import NewPackageDialog from '@/components/dialogs/NewPackageDialog.vue'
import EnhancedTable from "@/components/EnhancedTable.vue";
import { fetchAllObjects } from "@/middlewares/api.mdw";

export default Vue.extend({
  components: {
    EnhancedTable,
    NewPackageDialog
  },
  data: () => ({
    headers: [
      { text: "ID", value: "ID" },
      { text: "Vendor", value: "Vendor" },
      { text: "Name", value: "Name" },
      { text: "Status", value: "Status" },
      { text: "Created At", value: "CreatedAt" },
      { text: "Actions", value: "action", sortable: false }
    ],
    solverHeaders: [
      { text: "Vendor", value: "Vendor" },
      { text: "Package", value: "Package" },
      { text: "Name", value: "Name" },
      { text: "Status", value: "Status" },
      { text: "Created At", value: "CreatedAt" },
      { text: "Actions", value: "action", sortable: false }
    ],
    actions: [{ text: "View" }],
    solverActions:[{text:"Build"}],
    showNewPackageDialog: false
  }),
  mounted() {
    fetchAllObjects("packages");
    fetchAllObjects("solvers");
  },
  computed: {
    ...mapState({
      packages: "packages",
      solvers: "solvers"
    })
  }
});
</script>

<style scoped>
.solvers {
  margin-top: 5%;
}
</style>