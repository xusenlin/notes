<template>
    <div class="navigation">
        <v-navigation-drawer
            v-model="show"
            absolute
            temporary
        >
            <div class="nav-info">
                <h1>Notes</h1>
                <img class="wave" src="../assets/wave.gif" alt="">
            </div>
            <v-list dense>
                <v-list-tile
                    v-for="item in navigation"
                    :key="item.title"
                    @click="openRoute(item.path)"
                >
                    <v-list-tile-action>
                        <v-icon>{{ item.icon }}</v-icon>
                    </v-list-tile-action>

                    <v-list-tile-content>
                        <v-list-tile-title>{{ item.title }}</v-list-tile-title>
                    </v-list-tile-content>
                </v-list-tile>
            </v-list>
        </v-navigation-drawer>
    </div>
</template>

<script>
    import navigation from '../navigation/index'

    export default {
        name: 'Navigation',
        props: {},
        data: function () {
            return {
                show: false,
                navigation: navigation
            }
        },
        methods: {
            toggle() {
                this.show = !this.show;
            },
            openRoute(path){
                this.$router.push(path)

            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
    .nav-info
        width 100%
        height 70px
        position relative
        display flex
        color #fff
        justify-content center
        flex-direction column
        align-items center
        background: #3f51b5
        .wave
            width 100%
            position absolute
            bottom 0
            mix-blend-mode:screen;
            vertical-align: bottom;

</style>
