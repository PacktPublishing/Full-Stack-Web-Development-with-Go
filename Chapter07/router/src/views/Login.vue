<template>
  <body class="bg-gray-900">

  <div class="flex min-h-screen justify-center">
    <div class="flex flex-col items-center justify-center">
      <div class="flex flex-col space-x-2 items-start justify-start">
        <p class="text-2xl font-semibold text-white">FullyStacked 💪</p>
        <p class="text-xs text-gray-50">{{ loginText }}</p>
      </div>
      <!-- email -->
      <div class="pt-4 w-full flex space-x-2 items-center justify-end">
        <p class="text-xs text-gray-50">{{ emailText }}</p>
        <input v-model="username"
               class="border focus-within:border-red-600 focus:outline-2 border-red-500 rounded-md p-2 text-sm"/>
      </div>

      <!-- password -->
      <div class="pt-4 w-full flex space-x-2 items-center justify-end">
        <p class="text-xs font-bold text-white">{{ passwordText }}</p>
        <input v-model="password"
               class="border focus-within:border-red-600 focus:outline-2 border-red-500 rounded-md p-2 text-sm"/>
      </div>

      <div class="flex items-center justify-center mt-12">
        <button @click="handleSubmit"
                class="px-4 pt-2 pb-2.5 w-full rounded-lg bg-red-500 hover:bg-red-600">
          <span class="flex-1 text-sm font-bold text-center text-white">Login</span>
        </button>
      </div>
    </div>
  </div>
  </body>
</template>
<script type="module">
export default {
  data() {
    return {
      loginText: 'Login to your account',
      emailText: 'Email Address',
      passwordText: 'Password',
      username: 'enter username',
      password: 'enter password',
    };
  },
  methods: {
    handleSubmit: function () {
      console.log(this.$data.username)
    }
  }
};
</script>
