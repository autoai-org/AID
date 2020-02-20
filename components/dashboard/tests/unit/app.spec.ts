import { shallowMount } from '@vue/test-utils'
import App from '@/App.vue';
import router from '@/router';
import store from '@/store';

describe('App.vue', () => {
  it('renders title when created', () => {
    const wrapper = shallowMount(App, {
      router,
      store,
    })
    expect(wrapper.find('.title').text()).toMatch("AID Studio")
  })
})
