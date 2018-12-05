<template>
  <cvpm-log :title="'System Log'" :messageList="messageList"></cvpm-log>
</template>

<script>
import Log from '@/components/CVPM-Log'
export default {
  data () {
    return {
      messageList: ['an error occured', 'everything works fine']
    }
  },
  components: {
    'cvpm-log': Log
  },
  sockets: {
    connect: function () {
      console.log('socket connected')
    },
    logevent: function (data) {
      this.messageList.push(data)
    }
  }
}
</script>

<style>
</style>
