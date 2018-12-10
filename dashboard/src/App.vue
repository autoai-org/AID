<template>
  <div>
    <transition mode="out-in">
      <router-view></router-view>
    </transition>
    <v-tour name="cvpm-tour" :steps="steps" :callbacks="tourCallbacks"></v-tour>
  </div>
</template>

<script>
import { ConfigService } from '@/services/config'
export default {
  data () {
    return {
      steps: [
        {
          target: '#cvpm-tour-news',
          content: 'Discover latest news'
        },
        {
          target: '#cvpm-tour-status',
          content: 'Checkout the system status'
        },
        {
          target: '#cvpm-tour-packages',
          content: 'Click this to check installed packages'
        },
        {
          target: '#cvpm-tour-settings',
          content: 'Click this to check system settings'
        },
        {
          target: '#cvpm-tour-settings',
          content: 'To enable the hint, check it there'
        }
      ],
      tourCallbacks: {
        onNextStep: this.cvpmOnNextStep,
        onStop: this.cvpmStopHintMode
      }
    }
  },
  methods: {
    cvpmOnNextStep (currentStep) {
    },
    cvpmStopHintMode () {
      console.log('stop hint')
      let config = new ConfigService()
      config.hintMode = false
      config.write()
    }
  },
  mounted: function () {
    let config = new ConfigService()
    if (config.hintMode) {
      this.$tours['cvpm-tour'].start()
    }
  }
}
</script>
<style>
.v-step {
  z-index: 999;
}
</style>

