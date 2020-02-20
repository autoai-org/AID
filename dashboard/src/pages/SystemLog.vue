<template>
  <cvpm-log
    :title="$t(`Log.system_log`)"
    :message="message"
  />
</template>

<script>
import { systemService } from '@/services/system'
import Log from '@/components/CVPM-Log'

const conn = new WebSocket(systemService.getWsEndpoint('system'))

export default {
  components: {
    'cvpm-log': Log
  },
  data () {
    return {
      message: ''
    }
  },
  created () {
    const self = this
    conn.onmessage = function (evt) {
      self.message = evt.data
    }
  }
}
</script>

<style>
</style>
