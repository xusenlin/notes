<template>
  <div id="app">
    <v-app>
      <v-toolbar dark color="primary">
        <v-toolbar-side-icon @click="$refs.navigation.toggle()"></v-toolbar-side-icon>
        <v-toolbar-title class="white--text">{{ $route.name }}</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon>
          <v-icon>search</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon>apps</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon>refresh</v-icon>
        </v-btn>
        <v-btn icon>
          <v-icon>more_vert</v-icon>
        </v-btn>
      </v-toolbar>
      <Navigation ref="navigation" />
      <router-view/>
    </v-app>
  </div>
</template>
<script>

    import Navigation from '@/components/Navigation.vue'
    export default {
        name: 'App',
        components: {
            Navigation
        },
        data:function () {
            return {
                navigation:true
            }
        }
    }
</script>
<style lang="stylus">
#app
  font-family 'Avenir', Helvetica, Arial, sans-serif
  -webkit-font-smoothing antialiased
  -moz-osx-font-smoothing grayscale

</style>
