<template>
  <div id="app">
    <Quiz></Quiz>
  </div>
</template>

<script>

import Quiz from './components/Quiz'

export default {
  name: 'App',
  components: {
    Quiz
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, 'Arial Black', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e;
}
</style>
