<template>
  <v-container>
    <v-layout
      row
      wrap
    >
      <v-flex xs6>
        <cvpm-news id="cvpm-tour-news" />
      </v-flex>
      <v-flex xs6>
        <cvpm-status id="cvpm-tour-status" />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import News from '@/components/CVPM-News'
import Status from '@/components/CVPM-Status'
export default {
  components: {
    'cvpm-news': News,
    'cvpm-status': Status
  }

}
</script>

<style>
</style>
