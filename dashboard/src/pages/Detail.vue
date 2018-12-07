<template>
  <v-container>
    <v-layout row wrap>
      <v-flex xs6 class="cvpm-package-detail-card">
        <cvpm-repo-meta
          :config="config"
          :dependency="dependency"
          :diskSize="diskSize"
          :readme="readme"
        ></cvpm-repo-meta>
      </v-flex>
      <v-flex xs5 class="cvpm-package-detail-card">
        <cvpm-log :title="'Package Log'" :messageList="messageList"></cvpm-log>
      </v-flex>
      <v-flex xs6 class="cvpm-package-detail-card">
          <cvpm-repo-solver
            :config="parsedConfig"
            :vendor="toSelectVendor"
            :packageName="toSelectPackage"
          ></cvpm-repo-solver>
      </v-flex>
      <v-flex xs6 class="cvpm-package-detail-card">
          <cvpm-actions 
            :config="parsedConfig"
            :vendor="toSelectVendor"
            :packageName="toSelectPackage"
          ></cvpm-actions>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import cvpmRepoMeta from '@/components/CVPM-Repo-Meta'
import Log from '@/components/CVPM-Log'
import cvpmRepoSolver from '@/components/CVPM-Repo-Solver'
import cvpmActions from '@/components/CVPM-Actions'
import { systemService } from '@/services/system'
import toml from 'toml'
export default {
  data () {
    return {
      readme: '',
      config: '',
      parsedConfig: '',
      diskSize: '',
      dependency: '',
      messageList: []
    }
  },
  components: {
    'cvpm-repo-meta': cvpmRepoMeta,
    'cvpm-log': Log,
    'cvpm-repo-solver': cvpmRepoSolver,
    'cvpmActions': cvpmActions
  },
  methods: {
    fetchMeta () {
      let self = this
      systemService
        .getRepoMeta(this.$route.params.vendor, this.$route.params.name)
        .then(function (res) {
          self.readme = res.data.Readme
          self.config = res.data.Config
          self.parsedConfig = toml.parse(res.data.Config)
          self.diskSize = res.data.DiskSize.toFixed(2)
          self.dependency = res.data.Dependency
        })
    }
  },
  created () {
    this.fetchMeta()
  },
  sockets: {
    package: function (data) {
      this.messageList.push(data)
    }
  },
  computed: {
    toSelectVendor () {
      return [this.$route.params.vendor]
    },
    toSelectPackage () {
      return [this.$route.params.name]
    }
  }
}
</script>

<style scoped>
.cvpm-package-detail-card {
    margin: 1em;
}
</style>
