<script setup>
import { ref } from 'vue';
import * as demoAPI from '@/api/demo';

// Sample to show how we can inspect mode
// and import env variables
const deploymentMode = import.meta.env.MODE;
const myBaseURL = import.meta.env.VITE_BASE_API_URL;

const username = ref('');
const password = ref('');

async function doLogin() {
  const { data } = await demoAPI.doLogin(username.value, password.value);
  console.log(data);
  warnIncorrect();
}

const disabled = ref(false);
function warnIncorrect() {
  disabled.value = true;
  setTimeout(() => {
    disabled.value = false;
  }, 1500);
}

</script>

<template>
  <div class="flex min-h-screen bg-gray-800 justify-center">
    <div class="bg-gray-800 flex flex-col items-center justify-center">
      <div class="flex flex-col space-x-2 items-start justify-start">
        <p class="text-2xl font-semibold text-white">FullyStacked 💪</p>
        <p class="text-xs text-gray-50">Login to your account</p>
      </div>
      <form class="mt-12" @submit.prevent>
        <!-- email -->
        <div class="pt-4 w-full flex space-x-2 items-center justify-end">
          <p class="text-xs font-bold text-white">Email Address</p>
          <input
            v-model="username"
            type="text"
            class="border focus-within:border-red-600 focus:outline-2 border-red-500 rounded-md p-2 text-sm"
            placeholder="you@somewhere.com"
          />
        </div>

        <!-- password -->
        <div class="pt-4 w-full flex space-x-2 items-center justify-end">
          <p class="text-xs font-bold text-white">Password</p>
          <input
            v-model="password"
            type="password"
            class="border focus-within:border-red-600 focus:outline-2 border-red-500 rounded-md p-2 text-sm"
            placeholder="🔐"
          />
        </div>

        <div class="flex items-center justify-center flex-1 mt-12">
          <button
            @click="doLogin()"
            :class="{ shakeyboi: disabled }"
            class="px-4 pt-2 pb-2.5 w-full rounded-lg bg-red-500 hover:bg-red-600"
          >
            <span class="flex-1 text-sm font-bold text-center text-white">Login</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
.shakeyboi {
  -webkit-animation: kf_shake 0.4s 1 linear;
  -moz-animation: kf_shake 0.4s 1 linear;
  -o-animation: kf_shake 0.4s 1 linear;
}
@-webkit-keyframes kf_shake {
  0% {
    -webkit-transform: translate(30px);
  }
  20% {
    -webkit-transform: translate(-30px);
  }
  40% {
    -webkit-transform: translate(15px);
  }
  60% {
    -webkit-transform: translate(-15px);
  }
  80% {
    -webkit-transform: translate(8px);
  }
  100% {
    -webkit-transform: translate(0px);
  }
}
@-moz-keyframes kf_shake {
  0% {
    -moz-transform: translate(30px);
  }
  20% {
    -moz-transform: translate(-30px);
  }
  40% {
    -moz-transform: translate(15px);
  }
  60% {
    -moz-transform: translate(-15px);
  }
  80% {
    -moz-transform: translate(8px);
  }
  100% {
    -moz-transform: translate(0px);
  }
}
@-o-keyframes kf_shake {
  0% {
    -o-transform: translate(30px);
  }
  20% {
    -o-transform: translate(-30px);
  }
  40% {
    -o-transform: translate(15px);
  }
  60% {
    -o-transform: translate(-15px);
  }
  80% {
    -o-transform: translate(8px);
  }
  100% {
    -o-origin-transform: translate(0px);
  }
}
</style>
