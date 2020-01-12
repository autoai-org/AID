<template>
  <v-container>
    <v-btn class="ma-2" outlined color="indigo" @click="showNewPackageDialog=true">New Package</v-btn>
    <new-package-dialog :show="showNewPackageDialog" @closed="showNewPackageDialog=false" />
    <build-dialog :show="showBuildDialog" @closed="showBuildDialog=false" :vendor="selectedSolver.Vendor" :packageName="selectedSolver.Package" :solver="selectedSolver.Name"></build-dialog>
    <v-card outlined min-width="100%">
      <v-card-title>Packages</v-card-title>
      <EnhancedTable :headers="headers" :select="true" :items="packages" :actions="actions" />
    </v-card>
    <v-spacer />
    <v-card outlined min-width="100%" class="solvers">
      <v-card-title>Solvers</v-card-title>
      <EnhancedTable
        :headers="solverHeaders"
        :select="false"
        :items="solvers"
        :actions="solverActions"
        @actionClicked="handleActions"
      />
    </v-card>
  </v-container>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";
import NewPackageDialog from "@/components/dialogs/NewPackageDialog.vue";
import EnhancedTable from "@/components/EnhancedTable.vue";
import BuildDialog from "@/components/dialogs/BuildDialog.vue"
import { fetchAllObjects } from "@/middlewares/api.mdw";

export default Vue.extend({
  components: {
    EnhancedTable,
    NewPackageDialog,
    BuildDialog
  },
  data() {
    return {
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
      solverActions: [{ text: "Build" }],
      showNewPackageDialog: false,
      showBuildDialog: false,
      selectedSolver: {}
    };
  },
  methods: {
    handleActions(action: any) {
      if (action.action==="Build") {
        this.selectedSolver = action.item
        this.showBuildDialog = true
      }
    }
  },
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