<template>
  <v-card class="cvpm-status-card">
    <v-card-title primary-title>
      <h2>Status</h2>
    </v-card-title>
    <v-card-text >
        <p class="cvpm-status-content" v-if="status.cpu">CPU: {{status.cpu}}</p>
        <p class="cvpm-status-content" v-if="status.memory">Memory: {{(status.memory / 1024 / 1024).toFixed(2) }} MB</p>
        <p class="cvpm-status-content" v-if="status.platformVersion">Operating System: {{status.platform}} {{status.platformVersion}} on {{status.os}}</p>
        <p class="cvpm-status-content" v-if="status.installed">Installed: {{status.installed}}</p>
        <p class="cvpm-status-content" v-if="status.running">Running: {{status.running}} </p>
        <p class="cvpm-status-content" v-if="status.status">System Status: {{status.status}}</p>
    </v-card-text>
  </v-card>
</template>

<script>
import { systemService } from '@/services/system'
export default {
  data: () => ({
    status: {}
  }),
  methods: {
    getStatus () {
      let self = this
      systemService.getStatus().then(function (res) {
        self.status = res.data
      })
    }
  },
  created () {
    this.getStatus()
  },
  components: {
  }
}
</script>

<style scoped>
.cvpm-status-card {
  width: 95%;
}
.cvpm-status-content {
  padding-left: 2em;
  padding-right: 1em;
}
</style>
