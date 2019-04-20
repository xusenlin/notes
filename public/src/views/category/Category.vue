<template>
    <div class="category">
        <v-container fluid grid-list-md>
            <v-data-iterator
                    :items="items"
                    :rows-per-page-items="rowsPerPageItems"
                    :pagination.sync="pagination"
                    content-tag="v-layout"
                    row
                    wrap
            >
                <template v-slot:item="props">
                    <v-flex
                            xs12
                            sm6
                            md4
                            lg3
                    >
                        <v-card>
                            <v-card-title><h4>{{ props.item.Name }}</h4></v-card-title>
                            <v-divider></v-divider>
                            <v-list dense>
                                <v-list-tile>
                                    <v-list-tile-content>ID:</v-list-tile-content>
                                    <v-list-tile-content class="align-end">{{ props.item.ID }}
                                    </v-list-tile-content>
                                </v-list-tile>
                                <v-list-tile>
                                    <v-list-tile-content>ParentId:</v-list-tile-content>
                                    <v-list-tile-content class="align-end">{{ props.item.ParentId }}</v-list-tile-content>
                                </v-list-tile>
                                <v-list-tile>
                                    <v-list-tile-content>Order:</v-list-tile-content>
                                    <v-list-tile-content class="align-end">{{ props.item.Order }}</v-list-tile-content>
                                </v-list-tile>
                                <v-list-tile>
                                    <v-list-tile-content>Slug:</v-list-tile-content>
                                    <v-list-tile-content class="align-end">{{ props.item.Slug }}
                                    </v-list-tile-content>
                                </v-list-tile>
                                <v-list-tile>
                                    <v-list-tile-content>CreatedAt:</v-list-tile-content>
                                    <v-list-tile-content class="align-end">{{ props.item.CreatedAt }}
                                    </v-list-tile-content>
                                </v-list-tile>
                                <v-list-tile>
                                    <v-list-tile-content>UpdatedAt:</v-list-tile-content>
                                    <v-list-tile-content class="align-end">{{ props.item.UpdatedAt }}</v-list-tile-content>
                                </v-list-tile>
                            </v-list>
                        </v-card>
                    </v-flex>
                </template>
            </v-data-iterator>
        </v-container>
        <div  class="fab-btn">
            <v-fab-transition >
                <v-btn @click="editCategory" color="primary" absolute bottom right fab>
                    <v-icon>add</v-icon>
                </v-btn>
            </v-fab-transition>
        </div>

        <v-dialog v-model="edit.show" persistent max-width="600px">
            <v-card>
                <v-card-title>
                    <span class="headline">User Profile</span>
                </v-card-title>
                <v-card-text>
                    <v-container grid-list-md>
                        <v-layout wrap>
                            <v-flex xs12 sm6 md4>
                                <v-text-field label="Legal first name*" required></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6 md4>
                                <v-text-field label="Legal middle name" hint="example of helper text only on focus"></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6 md4>
                                <v-text-field
                                        label="Legal last name*"
                                        hint="example of persistent helper text"
                                        persistent-hint
                                        required
                                ></v-text-field>
                            </v-flex>
                            <v-flex xs12>
                                <v-text-field label="Email*" required></v-text-field>
                            </v-flex>
                            <v-flex xs12>
                                <v-text-field label="Password*" type="password" required></v-text-field>
                            </v-flex>
                            <v-flex xs12 sm6>
                                <v-select
                                        :items="['0-17', '18-29', '30-54', '54+']"
                                        label="Age*"
                                        required
                                ></v-select>
                            </v-flex>
                            <v-flex xs12 sm6>
                                <v-autocomplete
                                        :items="['Skiing', 'Ice hockey', 'Soccer', 'Basketball', 'Hockey', 'Reading', 'Writing', 'Coding', 'Basejump']"
                                        label="Interests"
                                        multiple
                                ></v-autocomplete>
                            </v-flex>
                        </v-layout>
                    </v-container>
                    <small>*indicates required field</small>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" flat @click="edit.show = false">Close</v-btn>
                    <v-btn color="blue darken-1" flat @click="edit.show = false">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>

    </div>
</template>
<script>
    import {getAllList} from '../../api/category'
    export default {
        data: () => ({
            rowsPerPageItems: [4, 8, 12],
            pagination: {
                rowsPerPage: 8
            },
            items: [],
            edit:{
                show:false
            }
        }),
        methods: {
            editCategory(id=null){
                this.edit.show = true
            }
        },
        created() {
            getAllList().then(r => {
                this.items = r.data
            }).catch(e => {
                console.log(e)
            })
        },
    }
</script>
