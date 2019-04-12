<template>
    <div class="notes">
        <v-container grid-list-md text-xs-center fluid>
            <v-layout row wrap>
                <v-flex xs12 sm6 md4 lg3 v-for="i in 10">
                    <v-card
                        class="mx-auto"
                        color="#26c6da"
                        dark
                        max-width="400"
                    >
                        <v-card-title>
                            <v-icon
                                large
                                left
                            >
                                card_membership
                            </v-icon>
                            <span class="title font-weight-light">Twitter</span>
                        </v-card-title>

                        <v-card-text class="headline font-weight-bold">
                            "Turns out semicolon-less style is easier and safer in TS because most gotcha edge cases are type invalid as well."
                        </v-card-text>

                        <v-card-actions>
                            <v-list-tile class="grow">
                                <v-list-tile-avatar color="grey darken-3">
                                    <v-img
                                        class="elevation-6"
                                        src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                                    ></v-img>
                                </v-list-tile-avatar>

                                <v-list-tile-content>
                                    <v-list-tile-title>Evan You</v-list-tile-title>
                                </v-list-tile-content>

                                <v-layout
                                    align-center
                                    justify-end
                                >
                                    <v-icon class="mr-1">mdi-heart</v-icon>
                                    <span class="subheading mr-2">256</span>
                                    <span class="mr-1">·</span>
                                    <v-icon class="mr-1">mdi-share-variant</v-icon>
                                    <span class="subheading">45</span>
                                </v-layout>
                            </v-list-tile>
                        </v-card-actions>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>

        <div class="pagination">
            <v-pagination
                v-model="page"
                :length="6"
                circle
            ></v-pagination>
        </div>
        <div class="edit-btn">
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
        .edit-btn
            position fixed
            right 20px
            bottom 100px

</style>
