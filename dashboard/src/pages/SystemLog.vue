<template>
  <cvpm-log
    :title="'System Log'"
    :message-list="messageList"
  />
</template>

<script>
import Log from '@/components/CVPM-Log'
export default {
  components: {
    'cvpm-log': Log
  },
  data () {
    return {
      messageList: []
    }
  },
  sockets: {
    system: function (data) {
      this.messageList.push(data)
    }
  }
}
</script>

<style>
</style>
