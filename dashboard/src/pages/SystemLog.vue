<template>
  <cvpm-log
    :title="$t(`Log.system_log`)"
    :message="message"
  />
</template>

<script>
import { systemService } from '@/services/system'
import Log from '@/components/CVPM-Log'

let conn = new WebSocket(systemService.getWsEndpoint('system'))

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
    let self = this
    conn.onmessage = function (evt) {
      self.message = evt.data
    }
  }
}
</script>

<style>
</style>
