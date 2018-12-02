<template></template>

<script>
import { discovery } from '@/services/discovery'
export default {
  data: () => ({
    news: []
  }),
  methods: {
    getNews () {
      discovery.fetchNews().then(function (res) {
        console.log(res)
      })
    }
  },
  created () {
    this.getNews()
  }
}
</script>

<style>
</style>
