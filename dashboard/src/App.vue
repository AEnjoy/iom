<script setup>
</script>

<template>
  <Login v-if="loadFlag === 0" @login-flag="getFlag" />
  <Dashboard v-if="loadFlag === 1" />
</template>
<script>
import axios from "@/api/axiosInstance";
import Login from './views/Login.vue'
import Dashboard from './views/dashboard/index.vue'
export default {
  data() {
    return {
      loadFlag: 0
    }
  },
  async beforeMount() {
    await axios.get('/api/auth/signin').then(response => {
      if (response.status === 200) {
        this.loadFlag = 1
      }
    }, error => {
      console.log('E', error.message)
    })
    console.log(this.loadFlag)
    //loadFlag = 1
  },
  components: {
    Login,
    Dashboard
  },
  methods: {
    getFlag(msg) {
      this.loadFlag = msg
    }
  }
}
</script>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
