<template>
  <v-app
    :dark="dark"
    standalone="standalone"
  >
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant.sync="mini"
      persistent="persistent"
      enable-resize-watcher="enable-resize-watcher"
      :dark="dark"
      class="cvpm-main-drawer"
    >
      <div
        v-show="!mini"
        class="pa-3 text-xs-center"
      >
        <div class="display-2 py-4">
          CVPM
        </div>
        <!-- if enable logo
            <img src="https://i.loli.net/2018/10/20/5bcb455c17616.png" class="cvpm-logo"/>
            -->
        <p>{{$t(`Dashboard`)}}</p>
        <div style="padding-left:5em;">
          <v-switch
            v-model="dark"
            :label="(!dark ? $t(`Theme.lighttheme`): $t(`Theme.darktheme`)) + $t(`Theme.theme`)"
            :dark="dark"
            hide-details="hide-details"
          />
        </div>
        <div>
          <v-btn
            dark="dark"
            tag="a"
            href="https://github.com/unarxiv/cvpm"
            primary="primary"
          >
            <v-icon
              left="left"
              dark="dark"
            >
              fab fa-github
            </v-icon><span>Star Us! </span>
          </v-btn>
        </div>
      </div>
      <div
        v-show="mini"
        class="pa-3 text-xs-center"
      >
        <div class="display-2">
          A
        </div>
      </div>
      <v-divider />

      <v-list dense="dense">
        <template v-for="item in menu">
          <v-list-group
            v-if="item.items"
            :key="item.header"
            :group="item.group"
          >
            <v-list-tile
              slot="item"
              :to="item.href"
              :title="item.title"
            >
              <v-list-tile-action><v-icon>{{ item.icon }}</v-icon></v-list-tile-action>
              <v-list-tile-content><v-list-tile-title>{{ item.title }}</v-list-tile-title></v-list-tile-content>
              <v-list-tile-action><v-icon>keyboard_arrow_down</v-icon></v-list-tile-action>
            </v-list-tile>
            <v-list-tile
              v-for="subItem in item.items"
              :key="subItem.href"
              :to="subItem.href"
              :router="!subItem.target"
              ripple="ripple"
              :disabled="subItem.disabled"
              :target="subItem.target"
            >
              <v-list-tile-action v-if="subItem.icon">
                <v-icon class="success--text">
                  {{ subItem.icon }}
                </v-icon>
              </v-list-tile-action>
              <v-list-tile-content><v-list-tile-title>{{ subItem.title }}</v-list-tile-title></v-list-tile-content>
            </v-list-tile>
          </v-list-group>
          <v-subheader
            v-else-if="item.header"
            :key="item.header"
          >
            {{ $t(item.header) }}
          </v-subheader>
          <v-divider
            v-else-if="item.header"
            :key="item.header"
          />
          <v-list-tile
            v-else-if="item.openType==='nav'"
            :key="item.header"
            :to="item.href"
            router="router"
            ripple="ripple"
            :disabled="item.disabled"
            :title="item.title"
          >
            <v-list-tile-action><v-icon>{{ item.icon }}</v-icon></v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title :id="item.id">
                {{ $t(item.title + ".title") }}
              </v-list-tile-title>
            </v-list-tile-content>
            <v-list-tile-action v-if="item.subAction">
              <v-icon class="success--text">
                {{ item.subAction }}
              </v-icon>
            </v-list-tile-action>
          </v-list-tile>
          <v-list-tile
            v-else-if="item.openType==='_blank'"
            :key="item.header"
            :href="item.href"
            :target="item.openType"
            ripple="ripple"
            :disabled="item.disabled"
            :title="item.title"
          >
            <v-list-tile-action><v-icon>{{ item.icon }}</v-icon></v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title :id="item.id">
                {{ $t(item.title) }}
              </v-list-tile-title>
            </v-list-tile-content>
            <v-list-tile-action v-if="item.subAction">
              <v-icon class="success--text">
                {{ item.subAction }}
              </v-icon>
            </v-list-tile-action>
          </v-list-tile>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar
      class="darken-1"
      fixed="fixed"
      dark="dark"
      :class="theme"
    >
      <v-toolbar-side-icon
        dark="dark"
        @click.native.stop="drawer = !drawer"
      />
      <v-toolbar-title>CVPM Dashboard</v-toolbar-title>
      <v-spacer />
      <v-menu offset-y="offset-y">
        <v-btn
          slot="activator"
          icon="icon"
          dark="dark"
        >
          <v-icon dark="dark">
            language
          </v-icon>
        </v-btn>
        <v-list>
          <v-list-tile
            v-for="lang in locales"
            :key="lang"
            @mouseover.native="changeLocale(lang)"
          >
            <v-list-tile-title>{{ lang }}</v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-menu>
      <v-menu offset-y="offset-y">
        <v-btn
          slot="activator"
          icon="icon"
          dark="dark"
        >
          <v-icon dark="dark">
            format_paint
          </v-icon>
        </v-btn>
        <v-list>
          <v-list-tile
            v-for="n in colors"
            :key="n"
            :class="n"
            @mouseover.native="theme = n"
          />
        </v-list>
      </v-menu>
    </v-toolbar>
    <main>
      <v-container
        class="pa-4"
        fluid
      >
        <v-alert
          v-if="message.display"
          v-model="message.body"
          v-bind="message"
          dismissible="dismissible"
        >
          {{ message.body }}
        </v-alert>
        <div
          class="py-2"
          :class="{'cvpm-main-content-with-drawer': drawer, 'cvpm-main-content-without-drawer':!drawer}"
        >
          <v-slide-y-transition mode="out-in">
            <router-view />
          </v-slide-y-transition>
        </div>
      </v-container>
    </main>
  </v-app>
</template>

<script>
import { getMenus } from '@/services/menu'
export default {
  data () {
    return {
      dark: false,
      theme: 'primary',
      mini: false,
      drawer: true,
      locales: ['en-US', 'zh-CN'],
      colors: ['blue', 'green', 'purple', 'red'],
      message: {
        display: false,
        body: 'Hello'
      },
      menu: []
    }
  },
  created () {
    this.fetchMenu()
  },
  methods: {
    changeLocale (to) {
      global.helper.ls.set('locale', to)
    },
    fetchMenu () {
      this.menu = getMenus()
    }
  }
}
</script>

<style scoped>
.cvpm-logo {
  width: 20em;
}
.cvpm-main-drawer {
    position:fixed;
}
.cvpm-main-content-with-drawer {
    margin-top: 5em;
    margin-left:  320px;
}
.cvpm-main-content-without-drawer {
    margin-top: 5em;
    margin-left:  20px;
}
</style>
