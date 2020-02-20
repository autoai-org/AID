<template>
  <div>
    <transition mode="out-in">
      <router-view v-if="isRouterAlive" />
    </transition>
    <v-tour
      name="cvpm-tour"
      :steps="steps"
      :callbacks="tourCallbacks"
    />
    <cvpm-loading />
  </div>
</template>

<script>
import { ConfigService } from '@/services/config'
import Loading from '@/components/global/Loading'
export default {
  components: {
    'cvpm-loading': Loading
  },
  data () {
    return {
      isRouterAlive: true,
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
  mounted: function () {
    const config = new ConfigService()
    if (config.hintMode) {
      this.$tours['cvpm-tour'].start()
    }
  },
  methods: {
    cvpmOnNextStep (currentStep) {
    },
    cvpmStopHintMode () {
      console.log('stop hint')
      const config = new ConfigService()
      config.hintMode = false
      config.write()
    },
    reload () {
      this.isRouterAlive = false
      this.$nextTick(function () {
        this.isRouterAlive = true
      })
    }
  },
  provide () {
    return {
      reload: this.reload
    }
  }
}
</script>
<style>
.v-step {
  z-index: 999;
}
</style>
