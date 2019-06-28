import Vue from 'vue'
import Vuetify from 'vuetify'

import ImageUploadComponent from '@/components/basic/CVPM-Image-Upload'

Vue.use(Vuetify)

describe('CVPM-Image-Upload.vue', () => {
  it('has the empty data function', () => {
    expect(typeof ImageUploadComponent.data).toBe('function')
  })
  it('has the right methods', () => {
    expect(typeof ImageUploadComponent.methods).toBe('object')
    expect(typeof ImageUploadComponent.methods.selectFile).toBe('function')
    expect(typeof ImageUploadComponent.methods.uploadFile).toBe('function')
  })
})
