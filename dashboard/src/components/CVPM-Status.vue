<template>
  <v-card class="cvpm-status-card">
    <v-card-title primary-title>
      <h2>Status</h2>
    </v-card-title>
    <v-card-text >
        <p class="cvpm-status-content">Installed: {{status.installed}}</p>
        <p class="cvpm-status-content">Running: {{status.running}} </p>
        <p class="cvpm-status-content">System Status: {{status.status}}</p>
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
      this.status = systemService.getStatus()
      console.log(this.status)
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
