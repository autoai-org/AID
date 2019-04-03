<template>
  <v-container
    fluid
    grid-list-xl
  >
    <v-layout
      row
      wrap
    >
      <v-flex xs12>
        <v-btn
          id="cvpm-tour-import-git"
          outline
          color="indigo"
          @click="trigger_git_import"
        >
          <v-icon
            left
            dark
          >
            fab fa-github
          </v-icon>Import from Git
        </v-btn>
        <v-btn
          outline
          color="indigo"
          @click="trigger_hub_import"
        >
          <v-icon
            left
            dark
          >
            device_hub
          </v-icon>Import fron Hub
        </v-btn>
      </v-flex>
      <v-flex xs12>
        <cvpm-table
          :items="packages"
          :loading="loading"
          :headers="headers"
          class="cvpm-package-table"
        />
      </v-flex>
      <v-dialog
        v-model="gitImport"
        persistent
        width="60%"
      >
        <cvpm-git-import @closeDialog="trigger_git_import()" />
      </v-dialog>
    </v-layout>
  </v-container>
</template>

<script>
import cvpmTable from '@/components/CVPM-Table'
import cvpmGitImport from '@/components/CVPM-Git-Import'
import { systemService } from '@/services/system'
export default {
  components: {
    'cvpm-git-import': cvpmGitImport,
    'cvpm-table': cvpmTable
  },
  data () {
    return {
      gitImport: false,
      packages: [],
      loading: true,
      headers: [
        {
          text: this.$t('Packages.vendor'),
          align: 'left',
          sortable: true,
          value: 'Vendor'
        },
        {
          text: this.$t('Packages.name'),
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
  created () {
    this.fetch_packages()
  },
  methods: {
    trigger_git_import () {
      this.gitImport = !this.gitImport
    },
    trigger_hub_import () {
      alert('coming soon')
    },
    fetch_packages () {
      let self = this
      systemService.getPackages().then(function (res) {
        if (res.data === null) {
          self.packages = []
          self.loading = false
        } else {
          self.packages = res.data
          self.loading = false
        }
      }).catch(function (err) {
        console.log(err)
        self.packages = []
        self.loading = false
      })
    }
  }
}
</script>

<style scoped>
</style>
