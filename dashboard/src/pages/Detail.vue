<template>
  <v-container>
    <v-layout
      row
      wrap
    >
      <v-flex
        xs5
        class="cvpm-package-detail-card"
      >
        <cvpm-repo-meta
          :config="config"
          :dependency="dependency"
          :disk-size="diskSize"
          :readme="readme"
        />
      </v-flex>
      <v-flex
        xs5
        class="cvpm-package-detail-card"
      >
        <cvpm-env-management
          :title="'environment variables'"
          :message-list="messageList"
        />
      </v-flex>
      <v-flex
        xs5
        class="cvpm-package-detail-card"
      >
        <cvpm-repo-solver
          :config="parsedConfig"
          :vendor="toSelectVendor"
          :package-name="toSelectPackage"
        />
      </v-flex>
      <v-flex
        xs5
        class="cvpm-package-detail-card"
      >
        <cvpm-actions
          :config="parsedConfig"
          :vendor="toSelectVendor"
          :package-name="toSelectPackage"
        />
      </v-flex>
      <v-flex
        xs5
        class="cvpm-package-detail-card"
      >
        <cvpm-log
          :title="$t('Packages_detail.package_log')"
          :message-list="messageList"
        />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import cvpmRepoMeta from '@/components/CVPM-Repo-Meta'
import Log from '@/components/CVPM-Log'
import cvpmRepoSolver from '@/components/CVPM-Repo-Solver'
import cvpmEnvManagement from '@/components/plugin/envs/CVPM-Env-Management'
import cvpmActions from '@/components/CVPM-Actions'
import { systemService } from '@/services/system'
import toml from 'toml'
export default {
  components: {
    'cvpm-repo-meta': cvpmRepoMeta,
    'cvpm-log': Log,
    'cvpm-repo-solver': cvpmRepoSolver,
    'cvpm-actions': cvpmActions,
    'cvpm-env-management': cvpmEnvManagement
  },
  data () {
    return {
      readme: '',
      config: '',
      parsedConfig: {},
      diskSize: '',
      dependency: '',
      messageList: []
    }
  },
  computed: {
    toSelectVendor () {
      return [this.$route.params.vendor]
    },
    toSelectPackage () {
      return [this.$route.params.name]
    }
  },
  created () {
    this.fetchMeta()
  },
  methods: {
    fetchMeta () {
      const self = this
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
  sockets: {
    package: function (data) {
      this.messageList.push(data)
    }
  }
}
</script>

<style scoped>
.cvpm-package-detail-card {
  margin: 1em;
}

</style>
