<template>
  <v-card>
    <v-card-title>
      <span class="headline">{{ $t('Packages.import_from_git') }}</span>
    </v-card-title>
    <v-card-text>
      <v-text-field
        v-model="repo"
        label="Git URL*"
        required
        hint="e.g: https://github.com/cvmodel/Face_Utility"
      />
    </v-card-text>
    <v-alert
      class="cvpm-git-alert"
      :value="error"
      type="error"
    >
      {{ error }}
    </v-alert>
    <v-alert
      outline
      class="cvpm-git-alert"
      :value="info"
      type="info"
    >
      {{ info }}
    </v-alert>
    <v-expansion-panel class="cvpm-git-import-detail">
      <v-expansion-panel-content v-if="cvpmConfig">
        <div slot="header">
          cvpm.toml
        </div>
        <v-card class="cvpm-config-text">
          <pre>{{ cvpmConfig }}</pre>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content v-if="readme">
        <div slot="header">
          {{ $t('Packages_detail.readme') }}
        </div>
        <v-card>
          <vue-markdown class="cvpm-repo-readme">
            {{ readme }}
          </vue-markdown>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content v-if="dependency">
        <div slot="header">
          {{ $t('Packages_detail.dependency') }}
        </div>
        <v-card class="cvpm-config-text">
          <pre>{{ dependency }}</pre>
        </v-card>
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-card-actions>
      <v-spacer />
      <v-btn
        color="indigo darken-1"
        outline
        @click="closeDialog()"
      >
        Close
      </v-btn>
      <v-btn
        color="indigo darken-1"
        outline
        @click="fetchMeta()"
      >
        Fetch Meta
      </v-btn>
      <v-btn
        color="indigo darken-1"
        :loading="installing"
        outline
        @click="save()"
      >
        Install
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { GithubService } from '@/services/github'
import { systemService } from '@/services/system'
import VueMarkdown from 'vue-markdown'
const Base64 = require('js-base64').Base64
export default {
  components: {
    'vue-markdown': VueMarkdown
  },
  data () {
    return {
      repo: '',
      cvpmConfig: '',
      dependency: '',
      readme: '',
      error: '',
      info: '',
      installing: false
    }
  },
  inject: ['reload'],
  methods: {
    closeDialog () {
      this.$emit('closeDialog', true)
    },
    fetchMeta () {
      // re-init error/info to avoid #282
      // For more info: https://github.com/unarxiv/CVPM/issues/282
      this.error = ''
      this.info = ''
      const self = this
      const pureRepo = this.repo.split('/')[3] + '/' + this.repo.split('/')[4]
      const githubService = new GithubService(pureRepo)
      githubService
        .fetchCVPMConfig()
        .then(function (res) {
          self.cvpmConfig = Base64.decode(res.data.content)
        })
        .catch(function (err) {
          self.error = err.response.data.message
        })
      githubService
        .fetchDependency()
        .then(function (res) {
          self.dependency = Base64.decode(res.data.content)
        })
        .catch(function (err) {
          self.error = err.response.data.message
        })
      githubService
        .fetchReadme()
        .then(function (res) {
          self.readme = Base64.decode(res.data.content)
        })
        .catch(function (err) {
          self.error = err.response.data.message
        })
    },
    save () {
      const self = this
      self.installing = true
      systemService
        .installRepo('git', this.repo)
        .then(function (res) {
          self.installing = false
          self.reload()
        })
        .catch(function (err) {
          self.error = err.response.data
        })
    }
  }
}
</script>

<style scoped>
.cvpm-config-text {
  padding: 2em;
}
.cvpm-git-import-detail {
  width: 95%;
  margin-left: auto;
  margin-right: auto;
}
pre {
  word-wrap: break-word;
  white-space: pre-wrap;
}
.cvpm-repo-readme {
  padding: 2em;
}
.cvpm-git-alert {
  width: 95%;
  margin-left: auto;
  margin-right: auto;
}
</style>
