<template>
<v-container>
  <v-btn class="ma-2" outlined color="indigo" @click="showNewPackageDialog=true">New Package</v-btn>
  <new-package-dialog :show="showNewPackageDialog" @closed="showNewPackageDialog=false"/>
  <v-card outlined min-width="100%">
    <EnhancedTable :headers="headers" :items="packages" :actions="actions"/>
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
    actions: [{ text: "Build" }, { text: "Meta" }],
    showNewPackageDialog: false
  }),
  mounted() {
    fetchAllObjects("packages");
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