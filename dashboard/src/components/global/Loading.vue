<template>
  <v-dialog
    v-model="loading"
    hide-overlay
    persistent
    width="300"
  >
    <v-card
      color="primary"
      dark
    >
      <v-card-text>
        Please stand by
        <v-progress-linear
          indeterminate
          color="white"
          class="mb-0"
        />
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
import store from '@/store'
export default {
  computed: {
    loading () {
      return store.state.isLoading
    }
  }
}
</script>

<style>
</style>
