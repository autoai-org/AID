<template>
<v-container>
    <cvpm-news></cvpm-news>
    <cvpm-status></cvpm-status>
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
