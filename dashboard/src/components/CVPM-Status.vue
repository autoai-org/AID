<template>
  <v-card class="cvpm-status-card">
    <v-card-title primary-title>
      <h2>{{$t('Home.status')}}</h2>
    </v-card-title>
    <v-card-text>
      <p
        v-if="status.cpu"
        class="cvpm-status-content"
      >
        <B>{{$t(`Home.cpu`)}}:</B> {{ status.cpu }}
      </p>
      <p
        v-if="status.memory"
        class="cvpm-status-content"
      >
        <B>{{$t(`Home.memory`)}}:</B> {{ (status.memory / 1024 / 1024).toFixed(2) }} MB
      </p>
      <p
        v-if="status.platformVersion"
        class="cvpm-status-content"
      >
        <B>{{$t(`Home.operating_system`)}}:</B> {{ status.platform }} {{ status.platformVersion }} on {{ status.os }}
      </p>
      <p
        v-if="status.installed"
        class="cvpm-status-content"
      >
        Installed: {{ status.installed }}
      </p>
      <p
        v-if="status.running"
        class="cvpm-status-content"
      >
        Running: {{ status.running }}
      </p>
      <p
        v-if="status.status"
        class="cvpm-status-content"
      >
        System Status: {{ status.status }}
      </p>
    </v-card-text>
  </v-card>
</template>

<script>
import { systemService } from '@/services/system'
export default {
  components: {
  },
  data: () => ({
    status: {}
  }),
  created () {
    this.getStatus()
  },
  methods: {
    getStatus () {
      let self = this
      systemService.getStatus().then(function (res) {
        self.status = res.data
      })
    }
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
  text-transform: capitalize;
}
</style>
