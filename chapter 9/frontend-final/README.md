# Vue Frontend

This was created with the following commands to create a Vite powered, Vue3 application.

```bash
$ npm create vite@latest
✔ Project name: … vue-frontend
✔ Select a framework: › vue
✔ Select a variant: › vue

Scaffolding project in /Users/nickglynn/Projects/Becoming-a-Full-Stack-Go-Developer/chapter 9/frontend/vue-frontend...

Done. Now run:

  cd vue-frontend
  npm install
  npm run dev
$ npm install
```

That last command is how you can run the dev server and yields the following:

```bash
$ npm run dev

> vue-frontend@0.0.0 dev
> vite


  vite v2.9.12 dev server running at:

  > Local: http://localhost:3000/
  > Network: use `--host` to expose

  ready in 332ms.
```

Using Vite really speeds things up!

We're now going to install tailwind which is what our earlier server side app used.

```bash
$ npm install -D tailwindcss postcss autoprefixer
$ npx tailwindcss init -p


Created Tailwind CSS config file: tailwind.config.js
Created PostCSS config file: postcss.config.js
$ cat << EOF > tailwind.config.js
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

EOF
$ cat << EOF > ./src/tailwind.css
@tailwind base;
@tailwind components;
@tailwind utilities;

EOF
$ cat << EOF > ./src/main.js
import { createApp } from 'vue'
import App from './App.vue'
import './tailwind.css'

createApp(App).mount('#app')

EOF
$ 
```
