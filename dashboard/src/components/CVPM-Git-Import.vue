<template>
  <v-card>
    <v-card-title>
      <span class="headline">Import From Git</span>
    </v-card-title>
    <v-card-text>
      <v-text-field label="Git URL*" required v-model="repo" hint="e.g: https://github.com/cvmodel/Face_Utility"></v-text-field>
    </v-card-text>
    <v-alert
      class="cvpm-git-alert"
      :value="error"
      type="error"
    >
      {{ error }}
    </v-alert>
    <v-alert
      class="cvpm-git-alert"
      :value="info"
      type="info"
    >
      {{ info }}
    </v-alert>
    <v-expansion-panel class="cvpm-git-import-detail">
      <v-expansion-panel-content v-if="cvpmConfig">
        <div slot="header">cvpm.toml</div>
        <v-card class="cvpm-config-text">
          <pre>{{ cvpmConfig }}</pre>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content v-if="readme">
        <div slot="header">Readme</div>
        <v-card>
          <vue-markdown class="cvpm-repo-readme">{{ readme }}</vue-markdown>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content v-if="dependency">
        <div slot="header">Dependency</div>
        <v-card class="cvpm-config-text">
          <pre>{{ dependency }}</pre>
        </v-card>
      </v-expansion-panel-content>
    </v-expansion-panel>
    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="indigo darken-1" outline @click="closeDialog()">Close</v-btn>
      <v-btn color="indigo darken-1" outline @click="fetchMeta()">Fetch Meta</v-btn>
      <v-btn color="indigo darken-1" outline @click="save()">Install</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
import { GithubService } from '@/services/github'
import VueMarkdown from 'vue-markdown'
let Base64 = require('js-base64').Base64
export default {
  data () {
    return {
      repo: '',
      cvpmConfig: '',
      dependency: '',
      readme: '',
      error: '',
      info: ''
    }
  },
  components: {
    'vue-markdown': VueMarkdown
  },
  methods: {
    closeDialog () {
      this.$emit('closeDialog', true)
    },
    fetchMeta () {
      let self = this
      const pureRepo = this.repo.split('/')[3] + '/' + this.repo.split('/')[4]
      let githubService = new GithubService(pureRepo)
      githubService.fetchCVPMConfig().then(function (res) {
        self.cvpmConfig = Base64.decode(res.data.content)
      }).catch(function (err) {
        self.error = err.response.data.message
      })
      githubService.fetchDependency().then(function (res) {
        self.dependency = Base64.decode(res.data.content)
      }).catch(function (err) {
        self.error = err.response.data.message
      })
      githubService.fetchReadme().then(function (res) {
        self.readme = Base64.decode(res.data.content)
      }).catch(function (err) {
        self.error = err.response.data.message
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
