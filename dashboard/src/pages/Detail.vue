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
            :config="config"
            :vendor="toSelectVendor"
            :packageName="toSelectPackage"
          ></cvpm-repo-solver>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import cvpmRepoMeta from '@/components/CVPM-Repo-Meta'
import Log from '@/components/CVPM-Log'
import cvpmRepoSolver from '@/components/CVPM-Repo-Solver'
import { systemService } from '@/services/system'

export default {
  data () {
    return {
      readme: '',
      config: '',
      diskSize: '',
      dependency: '',
      messageList: []
    }
  },
  components: {
    'cvpm-repo-meta': cvpmRepoMeta,
    'cvpm-log': Log,
    'cvpm-repo-solver': cvpmRepoSolver
  },
  methods: {
    fetchMeta () {
      let self = this
      systemService
        .getRepoMeta(this.$route.params.vendor, this.$route.params.name)
        .then(function (res) {
          self.readme = res.data.Readme
          self.config = res.data.Config
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
