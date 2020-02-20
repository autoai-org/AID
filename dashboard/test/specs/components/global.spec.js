import Vue from 'vue'
import Vuetify from 'vuetify'

import LoadingComponent from '@/components/global/Loading'

Vue.use(Vuetify)

describe('LoadingComponent.vue', () => {
  it('sets the correct default data', () => {
    expect(typeof LoadingComponent.data).toBe('function')
    const defaultData = LoadingComponent.data()
    expect(defaultData.loadingHint).toBe('Please stand by')
  })
})
