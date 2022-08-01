<template>
    <div class="container mt-5">
        <div class="row">
            <div class="col-md-3"></div>
            <div class="col-md-6">
                <div class="card border-0 rounded shadow">
                    <div class="card-body">
                        <h4 class="text-center">Login</h4>
                        <hr>

                        <form @submit.prevent="login">
                            <div class="form-group">
                                <label for="email" class="font-weight-bold">Email</label>
                                <input type="text" class="form-control" rows="4" v-model="user.email" placeholder="Masukkan Email">
                            </div>
                            <div class="form-group">
                                <label for="password" class="font-weight-bold">Password</label>
                                <input type="password" class="form-control" v-model="user.password" placeholder="Masukkan Password">
                            </div>
                            <button type="submit" class="btn btn-primary my-3">SIMPAN</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

export default {

    setup() {
        const user = reactive({
            email: '',
            password: ''
        })

        //state validation
        const validation = ref([])
        
        //vue router
        const router = useRouter()

        //method login
        function login() {
            let email   = user.email
            let password = user.password

            axios.post('http://localhost:3000/login', {
                email: email,
                password: password
            }).then(response => {
                sessionStorage.setItem("jwt", response.data.jwt);
                console.log(response.data)
            }).then(() => {

                //redirect ke blog index
                router.push({
                    name: 'blog.index'
                })

            }).catch(error => {
                //assign state validation with error 
                console.log(error)
                validation.value = error.response.data
            })

        }

        //return
        return {
            user,
            validation,
            router,
            login
        }

    }
}
</script>

<style>
    body{
        background: lightgray;
    }
</style>