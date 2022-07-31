<template>
    <div class="container mt-5">
        <div class="row">
            <div class="col-md-12">
                <div class="card border-0 rounded shadow">
                    <div class="card-body">
                        <h4>EDIT BLOG</h4>
                        <hr>
                        <form @submit.prevent="update">
                            <div class="form-group">
                                <label for="title" class="font-weight-bold">TITLE</label>
                                <input type="text" class="form-control" v-model="blog.title" placeholder="Masukkan Nama blog">
                                <!-- validation -->
                                <div v-if="validation.title" class="mt-2 alert alert-danger">
                                    {{ validation.title[0] }}
                                </div>
                            </div>
                            <div class="form-group">
                                <label for="content" class="font-weight-bold">CONTENT</label>
                                <textarea class="form-control" rows="10" v-model="blog.content" placeholder="Masukkan content blog"></textarea>
                                <!-- validation -->
                                <div v-if="validation.content" class="mt-2 alert alert-danger">
                                    {{ validation.content[0] }}
                                </div>
                            </div>
                            <button type="submit" class="btn btn-primary">SIMPAN</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { reactive, ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import axios from 'axios'

export default {

    setup() {

        //state blogs
        const blog = reactive({
            title: '',
            content: ''
        })

        //state validation
        const validation = ref([])

        //vue router
        const router = useRouter()

        //vue router
        const route = useRoute()

        //mounted
        onMounted(() => {

            //get API from Laravel Backend
            axios.get(`http://localhost:3000/blogs/${route.params.id}`)
            .then(response => {
              
              //assign state blogs with response data
              blog.title    = response.data.title  
              blog.content  = response.data.content  

            }).catch(error => {
                console.log(error.response.data)
            })

        })

        //method update
        function update() {

            let title   = blog.title
            let content = blog.content

            axios.put(`http://localhost:3000/blogs/${route.params.id}`, {
                title: title,
                content: content
            }).then(() => {

                //redirect ke blog index
                router.push({
                    name: 'blog.index'
                })

            }).catch(error => {
                //assign state validation with error 
                validation.value = error.response.data
            })

        }

        //return
        return {
            blog,
            validation,
            router,
            update
        }

    }

}
</script>

<style>
    body{
        background: lightgray;
    }
</style>