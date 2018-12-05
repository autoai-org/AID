<template>
  <v-container fluid grid-list-xl>
    <v-layout row wrap>
      <v-flex xs12>
      <v-btn outline color="indigo" @click="trigger_git_import">
        <v-icon left dark>add</v-icon>Import from Git
      </v-btn>
      <v-btn outline color="indigo">
        <v-icon left dark>device_hub</v-icon>Import fron Hub
      </v-btn>
      </v-flex>
      <v-flex xs12>
      <cvpm-table :items="packages" :loading="loading" :headers="headers" class="cvpm-package-table"></cvpm-table>
      </v-flex>
      <v-dialog persistent v-model="gitImport" width="60%">
        <cvpm-git-import v-on:closeDialog="trigger_git_import()"></cvpm-git-import>
      </v-dialog>
    </v-layout>
  </v-container>
</template>

<script>
import cvpmTable from '@/components/CVPM-Table'
import cvpmGitImport from '@/components/CVPM-Git-Import'
import { systemService } from '@/services/system'
export default {
  data () {
    return {
      gitImport: false,
      packages: [],
      loading: true,
      headers: [
        {
          text: 'Vendor',
          align: 'left',
          sortable: true,
          value: 'Vendor'
        },
        {
          text: 'Name',
          align: 'left',
          sortable: true,
          value: 'Name'
        },
        {
          text: 'LocalFolder',
          align: 'left',
          sortable: true,
          value: 'LocalFolder'
        },
        { text: 'Actions',
          align: 'left',
          value: 'name',
          sortable: false
        }
      ]
    }
  },
  methods: {
    trigger_git_import () {
      this.gitImport = !this.gitImport
    },
    fetch_packages () {
      let self = this
      systemService.getPackages().then(function (res) {
        self.packages = res.data
        self.loading = false
      })
    }
  },
  components: {
    'cvpm-git-import': cvpmGitImport,
    'cvpm-table': cvpmTable
  },
  created () {
    this.fetch_packages()
  }
}
</script>

<style scoped>
</style>
