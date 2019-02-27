import { shallowMount } from '@vue/test-utils'
import Main from '@/pages/Main'

describe('Main.vue', () => {
  it('renders props.msg when passed', () => {
    const name = 'tester'

    const wrapper = shallowMount(Main, {
      propsData: { name }
    })

    expect(wrapper.text()).toBe('CVPM')
  })
})
