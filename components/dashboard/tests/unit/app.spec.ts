import { shallowMount } from '@vue/test-utils'
import App from '@/App.vue';
import router from '@/router';
import store from '@/store';
import vuetify from '@/plugins/vuetify';

describe('App.vue', () => {
  it('renders title when created', () => {
    const wrapper = shallowMount(App, {
      router,
      store,
      vuetify,
    })
    expect(wrapper.find('.title').text()).toMatch("AI Flow")
  })
})
