<template>
  <v-card min-width="100%" color="transparent" flat>
    <v-card-title class="text-center justify-center py-6">
      <h1
        class="font-weight-bold display-1 basil--text"
      >{{$route.params.vendor}}/{{$route.params.package}}</h1>
      <p>
        <v-img width="32px" v-if="frameworks.includes('keras')" src="@/assets/framework-icons/keras.png"/>
        <v-img width="32px" v-if="frameworks.includes('opencv')" src="@/assets/framework-icons/opencv.png"/>
        <v-img width="32px" v-if="frameworks.includes('pytorch')" src="@/assets/framework-icons/pytorch.png"/>
        <v-img width="32px" v-if="frameworks.includes('tensorflow')" src="@/assets/framework-icons/tensorflow.png"/>
      </p>
    </v-card-title>
    <v-tabs v-model="tab" background-color="transparent" color="basil" grow>
      <v-tab v-for="item in items" :key="item">{{ item }}</v-tab>
    </v-tabs>

    <v-tabs-items v-model="tab">
      <v-tab-item :key="'Readme'">
        <v-card flat color="transparent">
          <v-card-text class="markdown-text">
            <vue-markdown :source="meta.readme" />
          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item :key="'Solvers'">
        <solvers-card v-if="isReady" :solvers="meta.solvers.Solvers" :vendor=$route.params.vendor :packageName=$route.params.package />
      </v-tab-item>
      <v-tab-item :key="'Insight'"></v-tab-item>
      <v-tab-item :key="'Settings'">
        <v-card flat color="transparent">
          <v-card-text class="markdown-text">
            <package-settings/>
          </v-card-text>
        </v-card>
      </v-tab-item>
      <v-tab-item :key="'Plugin'">
        <plugin-page></plugin-page>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script lang="ts">
import PackageSettings from "@/components/packageComponents/Settings.vue";
import SolversCard from "@/components/packageComponents/Solvers.vue"
import PluginPage from "@/components/packageComponents/PluginPage.vue"
import VueMarkdown from "vue-markdown";
import { fetchMeta } from "@/middlewares/api.mdw";
import Vue from "vue";
export default Vue.extend({
  data() {
    return {
      isReady: false,
      tab: null,
      items: ["Readme", "Solvers", "Insight", "Settings", "Plugin"],
      meta: {},
      frameworks:""
    };
  },
  components: {
    "package-settings": PackageSettings,
    "vue-markdown": VueMarkdown,
    "solvers-card": SolversCard,
    "plugin-page": PluginPage,
  },
  mounted() {
    let self = this;
    fetchMeta(this.$route.params.vendor, this.$route.params.package).then(
      function(res: any) {
        self.meta = res;
        self.frameworks = res.requirements
        self.isReady = true
      }
    );
  }
});
</script>

<style scoped>
.markdown-text {
  color: black !important;
}
</style>