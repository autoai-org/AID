<template>
  <section id="signup" class="signUpPage">
    <VCard class="v-card--auth pa-5" color="white">
      <VContainer pa-0 grid-list-xl>
        <VLayout>
          <VFlex xs8>
            <h2 class="headline font-weight-black primary--text">AutoAI</h2>
            <h1 class="headline font-weight-regular mb-5">Create your AutoAI Account</h1>
            <VForm ref="form" v-model="form" class="mb-5">
              <VContainer pa-0>
                <VLayout wrap>
                  <VFlex xs6>
                    <VTextField
                      v-model="firstName"
                      :rules="[rules.required('Enter first name')]"
                      label="First name"
                    />
                  </VFlex>
                  <VFlex xs6>
                    <VTextField
                      v-model="lastName"
                      :rules="[rules.required('Enter last name')]"
                      label="Last name"
                    />
                  </VFlex>
                  <VFlex xs12>
                    <VTextField
                      v-model="email"
                      :rules="[rules.required('Enter your email address')]"
                      label="Email"
                      hint="You can use letters, numbers & periods"
                      name="email"
                      persistent-hint
                    />
                  </VFlex>
                  <VFlex xs6>
                    <VTextField
                      v-model="password"
                      :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                      :rules="[rules.required('Enter a password')]"
                      :type="show ? 'text' : 'password'"
                      label="Password"
                      @click:append="show = !show"
                    />
                  </VFlex>
                  <VFlex xs6>
                    <VTextField
                      v-model="confirmPassword"
                      :append-icon="showConfirm ? 'mdi-eye' : 'mdi-eye-off'"
                      :rules="[rules.required('Confirm you password'), rules.confirm]"
                      :type="showConfirm ? 'text' : 'password'"
                      label="Confirm password"
                      @click:append="showConfirm = !showConfirm"
                    />
                  </VFlex>
                  <VFlex xs12>
                    <VMessages
                      :value="['Use 8 or more characters with a mix of letters, numbers & symbols']"
                    />
                  </VFlex>
                </VLayout>
                <VLayout align-center justify-space-between pt-3>
                  <VFlex grow>
                    <BaseText to="/auth/signin">Sign in instead</BaseText>
                  </VFlex>
                  <VFlex shrink>
                    <BaseBtn :disabled="!form" :loading="isLoading" prominent @click="submit">Next</BaseBtn>
                  </VFlex>
                </VLayout>
              </VContainer>
            </VForm>
          </VFlex>
          <VFlex align-center column layout justify-center xs4>
            <VImg
              :src="require('@/assets/logo.png')"
              class="mb-3"
              height="100"
              max-height="100"
              style="width: 100%;"
              contain
            />
            <div
              class="flex-grow subheading font-weight-regular text-xs-center"
            >One account. All of AutoAI
              <br>working for you.
            </div>
          </VFlex>
        </VLayout>
      </VContainer>
    </VCard>
  </section>
</template>

<script>
// Utilities
import { mapActions, mapMutations } from 'vuex'
export default {
  data () {
    const data = {
      isLoading: false,
      form: false,
      firstName: undefined,
      lastName: undefined,
      email: undefined,
      password: undefined,
      confirmPassword: undefined,
      rules: {
        required: msg => v => !!v || msg,
        confirm: v => (v ? v === this.password : 'Passwords do not match')
      },
      show: false,
      showConfirm: false
    }
    return data
  },
  methods: {
    ...mapActions('cognito', ['registerUser']),
    ...mapMutations(['setSnackbar']),
    submit () {
      if (!this.$refs.form.validate()) return
      this.isLoading = true
      this.registerUser({
        username: this.email,
        password: this.password,
        attributes: {
          name: `${this.firstName} ${this.lastName}`,
          email: this.email
        }
      })
        .then(() => {
          this.setSnackbar({
            type: 'success',
            msg: 'Account created. Check your email for verification'
          })
          this.$router.push('/authenticated')
        })
        .catch(err => {
          this.setSnackbar({
            type: 'error',
            msg: err.message,
            timeout: 10000
          })
        })
        .finally(() => (this.isLoading = false))
    }
  }
}
</script>

<style scoped>
.signUpPage {
  width: 60%;
  height: 60%;
  margin: auto;
  margin-top: 5%;
}
</style>
