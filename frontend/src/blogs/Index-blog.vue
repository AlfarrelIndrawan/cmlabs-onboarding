<template>
    <div class="container mt-5">
        <div class="row">
            <div class="col-md-12">
                <div class="card border-0 rounded shadow">
                    <div class="card-body">
                        <h4>Blogs On-Boarding</h4>
                        <hr>
                        <router-link :to="{name: 'blog.create'}" class="btn btn-md btn-success">Buat Blog Baru</router-link>
                        <br>
                        <div v-for="(blog, index) in blogs" :key="index" class="mt-4">
                            <router-link :to="{name: 'blog.detail', params:{id: blog.ID}}" class="text-reset text-decoration-none">
                                <div class="card">
                                    <div class="card-header">
                                        {{new Date(blog.CreatedAt).toDateString()}}
                                    </div>
                                    <div class="card-body">
                                        <h5 class="card-title">{{blog.title}}</h5>
                                        <p class="card-text">{{blog.content.substring(0,500)}}...</p>
                                        <router-link :to="{name: 'blog.edit', params:{id: blog.ID }}" class="btn btn-sm btn-primary mx-3">Edit Blog</router-link>
                                        <button @click.prevent="blogDelete(blog.ID)" class="btn btn-sm btn-danger ml-1">Delete Blog</button>
                                    </div>
                                </div>
                            </router-link>
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

export default {

    setup() {

        //reactive state
        let blogs = ref([])

        //mounted
        onMounted(() => {

            //get API from Laravel Backend
            axios.get('http://localhost:3000/blogs')
            .then(response => {
              
              //assign state blogs with response data
              blogs.value = response.data

            }).catch(error => {
                console.log(error.response.data)
            })

        })

        //method delete
        function blogDelete(id) {
            
            //delete data blog by ID
            axios.delete(`http://localhost:3000/blogs/${id}`)
            .then(() => {
              
              //splice blogs 
              blogs.value.splice(blogs.value.indexOf(id), 1);
            }).catch(error => {
                console.log(error.response.data)
            })
        }

        //return
        return {
            blogs,
            blogDelete
        }

    }

}
</script>

<style>
    body{
        background: lightgray;
    }
</style>