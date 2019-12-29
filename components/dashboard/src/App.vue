<template>
  <v-app id="inspire">
    <v-navigation-drawer :dark="dark" v-model="drawer" app clipped>
      <v-list subheader>
        <v-subheader>Overview</v-subheader>
        <v-list-item v-for="item in menus[0]" :key="item.text" link>
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider />
        <v-subheader>System</v-subheader>
        <v-list-item v-for="item in menus[1]" :key="item.text" link>
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>

        <v-divider />
        <v-subheader>About</v-subheader>
        <v-list-item v-for="item in menus[2]" :key="item.text" link>
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>{{ item.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar :dark="dark" app clipped-left color="primary">
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title class="mr-12 align-center">
        <span class="title">AI Flow</span>
      </v-toolbar-title>
      <v-spacer />
      <v-row align="center" style="max-width: 650px">
        <v-text-field
          :append-icon-cb="() => {}"
          placeholder="Search..."
          single-line
          append-icon="mdi-magnify"
          color="white"
          hide-details
        />
      </v-row>
    </v-app-bar>
    <v-dialog v-model="isLoading" hide-overlay persistent width="300">
      <v-card color="primary" dark>
        <v-card-text>
          Please stand by...
          <v-progress-linear indeterminate color="white" class="mb-0"></v-progress-linear>
        </v-card-text>
      </v-card>
    </v-dialog>
    <v-content>
      <v-container class="fill-height">
        <router-view />
      </v-container>
    </v-content>
  </v-app>
</template>

<script lang="ts">
import { mapState } from "vuex";
import { overview_menu, system_menu, about_menu } from "./router/menu";
export default {
  props: {
    source: String
  },
  data: () => ({
    dark: true,
    drawer: null,
    menus: [overview_menu, system_menu, about_menu]
  }),
  computed: {
    ...mapState({
      isLoading: "isLoading"
    })
  }
};
</script>

<style scoped>
.fill-height {
  align-items: baseline!important;
}
</style>