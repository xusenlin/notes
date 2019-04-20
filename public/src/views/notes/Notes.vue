<template>
    <div class="notes">
        <v-container grid-list-md text-xs-center fluid style="margin-bottom: 100px">
            <v-layout row wrap>
                <v-flex xs12 sm6 md4 lg3 v-for="i in 20" :key="i" style="padding: 14px">
                    <v-card
                        class="mx-auto"
                        color="#26c6da"
                        dark
                    >
                        <v-card-title>
                            <v-icon large left>text_fields</v-icon>
                            <span class="title">这里是标题</span>
                        </v-card-title>

                        <v-card-text class="headline font-weight-bold">
                            "Turns out semicolon-less style is easier and safer in TS because most gotcha edge cases are type invalid as well."
                        </v-card-text>

                        <v-card-actions>
                            <v-list-tile class="grow">

                                <v-icon class="mr-1">folder_open</v-icon>
                                <v-list-tile-content>
                                    <v-list-tile-title>Evan You</v-list-tile-title>
                                </v-list-tile-content>

                                <v-icon class="mr-1">edit</v-icon>

                                <v-icon class="mr-1">delete</v-icon>
                            </v-list-tile>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>

        <div class="pagination" style="margin-bottom: 30px">
            <v-pagination
                v-model="page"
                :length="6"
                circle
            ></v-pagination>
        </div>
        <div class="fab-btn">
            <v-fab-transition>
                <v-btn color="primary" absolute bottom right fab>
                    <v-icon>edit</v-icon>
                </v-btn>
            </v-fab-transition>
        </div>
    </div>
</template>

<script>


    export default {
        name: 'notes',
        data:function () {
            return {
                page:1
            }
        }
    }
</script>
<style lang="stylus">
    .notes
        text-align center
        position relative
        width 100%
        min-height calc(100vh - 64px)
        .pagination
            width 100%
            bottom 10px
            position absolute


</style>
