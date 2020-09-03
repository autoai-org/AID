<template>
  <v-card outlined min-width="100%" min-height="100%">
    <v-card-title>Packages</v-card-title>
    <v-container width="100%">
      <v-btn class="ma-2" outlined color="indigo" @click="showNewPackageDialog=true">New Package</v-btn>
      <new-package-dialog :show="showNewPackageDialog" @closed="showNewPackageDialog=false" />
      <build-dialog
        :show="showBuildDialog"
        @closed="showBuildDialog=false"
        :vendor="selectedSolver.Vendor"
        :packageName="selectedSolver.Package"
        :solver="selectedSolver.Name"
      ></build-dialog>
      <v-card outlined min-width="100%">
        <v-card-title>Packages</v-card-title>
        <EnhancedTable
          :headers="headers"
          :select="true"
          :items="packages"
          :actions="actions"
          @actionClicked="handleActions"
        />
      </v-card>
      <v-spacer />
    </v-container>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from "vuex";
import NewPackageDialog from "@/components/dialogs/NewPackageDialog.vue";
import EnhancedTable from "@/components/EnhancedTable.vue";
import BuildDialog from "@/components/dialogs/BuildDialog.vue";
import {
  fetchAllObjects,
  fetchPackages
} from "@/middlewares/api.mdw";

export default Vue.extend({
  components: {
    EnhancedTable,
    NewPackageDialog,
    BuildDialog
  },
  data() {
    return {
      packages: [],
      headers: [
        { text: "ID", value: "ID" },
        { text: "Vendor", value: "Vendor" },
        { text: "Name", value: "Name" },
        { text: "Frameworks", value: "Frameworks" },
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
      if (action.action === "Build") {
        this.selectedSolver = action.item;
        this.showBuildDialog = true;
      } else if (action.action === "View") {
        console.log(action.item);
        this.$router.replace(
          "/package/" + action.item.Vendor + "/" + action.item.Name
        );
      }
    }
  },
  mounted() {
    let self = this;
    fetchPackages().then(function(res: any) {
      self.packages = res;
    });
    fetchAllObjects("solvers");
  },
  computed: {
    ...mapState({
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