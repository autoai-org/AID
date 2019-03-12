<template>
  <v-card>
    <v-card-title>
      <h2>Meta Info</h2>
    </v-card-title>
    <v-expansion-panel class="cvpm-meta-detail">
      <v-expansion-panel-content v-if="config">
        <div slot="header">
          cvpm.toml
        </div>
        <v-card class="cvpm-config-text">
          <pre>{{ config }}</pre>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content v-if="readme">
        <div slot="header">
          Readme
        </div>
        <v-card>
          <vue-markdown class="cvpm-repo-readme">
            {{ readme }}
          </vue-markdown>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content v-if="dependency">
        <div slot="header">
          Dependency
        </div>
        <v-card class="cvpm-config-text">
          <pre>{{ dependency }}</pre>
        </v-card>
      </v-expansion-panel-content>
      <v-expansion-panel-content v-if="diskSize">
        <div slot="header">
          Disk Usage
        </div>
        <v-card class="cvpm-config-text">
          <pre>{{ diskSize }} MB</pre>
        </v-card>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-card>
</template>

<script>
import VueMarkdown from 'vue-markdown'

export default {
  components: {
    'vue-markdown': VueMarkdown
  },
  props: {
    config: {
      type: String,
      default: ''
    },
    dependency: {
      type: String,
      default: ''
    },
    diskSize: {
      type: String,
      default: ''
    },
    readme: {
      type: String,
      default: ''
    }
  }
}
</script>

<style scoped>
.cvpm-config-text {
  padding: 2em;
}
pre {
  word-wrap: break-word;
  white-space: pre-wrap;
}
.cvpm-repo-readme {
  padding: 2em;
}
.cvpm-meta-detail {
  margin-left: auto;
  margin-right: auto;
}
</style>
