<template>
    <div class="container mt-5">
        <div class="row">
            <div class="col-md-12">
                <div class="card border-0 rounded shadow">
                    <div class="card-body">
                        <h4>Blogs On-Boarding</h4>
                        <div style="text-align: right">
                            <button @click.prevent="logout" class="btn btn-sm btn-danger ml-1 hide-guest">Logout</button>
                        </div>
                        <hr>
                        <router-link :to="{name: 'blog.create'}" class="btn btn-sm btn-success hide-guest">Buat Blog Baru</router-link>
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
                                        <router-link :to="{name: 'blog.edit', params:{id: blog.ID }}" class="btn btn-sm btn-primary mx-3 hide-guest">Edit Blog</router-link>
                                        <button @click.prevent="blogDelete(blog.ID)" class="btn btn-sm btn-danger ml-1 hide-guest">Delete Blog</button>
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
import { useRouter } from 'vue-router'


export default {

    setup() {

        //reactive state
        let blogs = ref([])

        //vue router
        const router = useRouter()

        //mounted
        onMounted(() => {
            //get API from Laravel Backend
            axios.get('http://localhost:3000/blogs')
            .then(response => {
                blogs.value = response.data
            }).catch(error => {
                console.log(error.response.data)
            })
            if (sessionStorage.getItem("jwt") == null) {
                let elements = document.getElementsByClassName("hide-guest");
                for(let element of elements) {
                    element.style.display = "none";
                }
            }
        })

        //method delete
        function blogDelete(id) {
            
            //delete data blog by ID
            axios.delete(`http://localhost:3000/blogs/${id}`, {
                headers: {
                    'jwt': sessionStorage.getItem("jwt")
                }
            })
            .then(() => {
              
              //splice blogs 
              blogs.value.splice(blogs.value.indexOf(id), 1);
            }).catch(error => {
                console.log(error.response.data)
            })
        }

        function logout() {
            sessionStorage.removeItem("jwt");
            //delete data blog by ID
            axios.post(`http://localhost:3000/logout/`)
            .then(response => {
                console.log(response)
            }).then(() => {

                //redirect ke blog index
                router.push({
                    name: 'blog.index'
                })

            }).catch(error => {
                console.log(error.response.data)
            })
            let elements = document.getElementsByClassName("hide-guest");
            for(let element of elements) {
                element.style.display = "none";
            }
            location.reload();
        }
        //return
        return {
            blogs,
            blogDelete,
            logout,
            router
        }
    }
}

</script>

<style>
    body{
        background: lightgray;
    }
</style>