<template>
  <cvpm-log :title="'System Log'" :messageList="messageList"></cvpm-log>
</template>

<script>
import Log from '@/components/CVPM-Log'
export default {
  data () {
    return {
      messageList: []
    }
  },
  components: {
    'cvpm-log': Log
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
