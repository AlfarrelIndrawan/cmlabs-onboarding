<template>
    <div class="container mt-5">
        <div class="row">
            <div class="col-md-12">
                <div class="card border-0 rounded shadow">
                    <div class="card-body">
                        <h4>Blogs On-Boarding</h4>
                        <hr>
                        <router-link :to="{name: 'blog.index'}" class="btn btn-md btn-primary">Back</router-link>
                        <div class="card mt-4">
                            <div class="card-header">
                                {{new Date(blog.CreatedAt).toDateString()}}
                            </div>
                            <div class="card-body">
                                <h5 class="card-title">{{blog.title}}</h5>
                                <div v-for="(content, index) in contents" :key="index" class="card-text">
                                    <p>{{content}}</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

export default {

    setup() {

        //state blogs
        let blog = ref([])
        let contents = ref([])

        const route = useRoute()

        //mounted
        onMounted(() => {

            //get API from Laravel Backend
            axios.get(`http://localhost:3000/blogs/${route.params.id}`)
            .then(response => {
              
              //assign state blogs with response data
              blog.value = response.data
              contents.value = new String(response.data.content).split("\n")

            }).catch(error => {
                console.log(error.response.data)
            })

        })

        //return
        return {
            blog,
            contents
        }
    }
}
</script>

<style>
    body{
        background: lightgray;
    }
</style>