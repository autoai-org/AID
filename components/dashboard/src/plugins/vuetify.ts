import Vue from 'vue';
import Vuetify from 'vuetify/lib';
import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
      themes: {
        dark: {
          primary: colors.lightBlue.darken4, // #E53935
          secondary: colors.shades.white, // #FFCDD2
          accent: colors.indigo.base, // #3F51B5
        },
      },
    },
  })
  