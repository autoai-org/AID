<template>
  <v-card class="cvpm-news-card">
    <v-card-title primary-title>
      <h2>{{ $t(`Home.latest_news`) }}</h2>
    </v-card-title>
    <v-list
      two-line
      dense
    >
      <template v-for="(item, index) in news">
        <v-card-title
          v-if="item.title"
          :key="index"
          primary-title
        >
          <h4>
            <a
              :href="item.url"
              target="_blank"
            >{{ item.title }} </a>
          </h4>
        </v-card-title>
        <v-list-tile-content
          :key="item.id"
          class="cvpm-news-content"
        >
          <vue-markdown>{{ item.body }}</vue-markdown>
        </v-list-tile-content>
      </template>
    </v-list>
  </v-card>
</template>

<script>
import { discovery } from '@/services/discovery'
import VueMarkdown from 'vue-markdown'
export default {
  components: {
    'vue-markdown': VueMarkdown
  },
  data: () => ({
    news: []
  }),
  created () {
    this.getNews()
  },
  methods: {
    getNews () {
      const self = this
      discovery.fetchNews().then(function (res) {
        self.news = res.data.map(function (each) {
          each.title = each.body.split('\n\n')[0]
          each.url = 'https://write.as/autoai/' + each.slug
          if (each.body.length > 300) {
            each.body = each.body.substring(0, 300) + '...'
          }
          return each
        })
      })
    }
  }
}
</script>

<style scoped>
.cvpm-news-content {
  padding-left: 2em;
  padding-right: 1em;
}
.cvpm-news-card {
  width: 90%;
}
</style>
