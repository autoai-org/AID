<template>
  <v-app id="inspire">
    <div class="parent">
      <div class="left">
        <v-btn outlined x-large color="#174e55" class="new_button">New host</v-btn>
        <v-list class="hosts_list">
          <v-subheader>HOSTS</v-subheader>
          <v-list-item-group v-model="item" color="primary">
            <v-list-item v-for="(item, i) in items" :key="i">
              <v-list-item-icon>
                <v-icon>mdi-server</v-icon>
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title v-text="item.text"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>
      </div>

      <div class="right">
        <v-btn
          outlined
          x-large
          color="#fff"
          @click="connect('http://localhost:10590/')"
        >Connect To Localhost</v-btn>
        <div class="version_text">Client Version: 1.0</div>
        <v-footer color="#174e55" padless class="landing_footer">
          <v-row justify="center" no-gutters>
            <v-btn
              v-for="link in links"
              :key="link"
              color="white"
              text
              rounded
              class="my-2"
            >{{ link }}</v-btn>
            <v-col class="#174e55 py-4 text-center white--text" cols="12">
              {{ new Date().getFullYear() }} —
              <strong>AutoAI - AID Project</strong>
            </v-col>
          </v-row>
        </v-footer>
      </div>
    </div>
  </v-app>
</template>
<script>
import store from "@/store";
import { setEndpoint } from "@/middlewares/api.mdw";
export default {
  data: () => ({
    item: 1,
    items: [
      { text: "127.0.0.1" },
      { text: "demo.aid.autoai.org" },
      { text: "111.231.241.85" },
    ],
    links: ["Docs", "License", "Source"],
  }),
  methods: {
    connect(target) {
      setEndpoint(target);
      store.commit("setConnected", true);
    },
  },
};
</script>
<style scoped>
.left {
  display: flex;
  width: 35%;
  background-color: white;
  flex-direction: column;
  align-items: center;
}

.new_button {
  margin-top: 15%;
  width: 61.8%;
}

.hosts_list {
  width: 100%;
}

.right {
  display: flex;
  flex: auto;
  background-color: #174e55;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
}

.version_text {
  color: white;
  margin-top: 10px;
  margin-bottom: 20%;
}

.landing_footer {
  justify-self: flex-end;
}

.parent {
  display: flex;
  height: 100vh;
}
#inspire {
  overflow-y: hidden;
}
</style>