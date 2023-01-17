<script setup lang="ts">
import Container from '../global/Container.vue';
import { useRouter } from 'vue-router';
import { Auth, StorageKey } from "../../assets/ts"

function Logout(): void {
    StorageKey.Delete(Auth.ACCESS_TOKEN_KEY())
    useRouter().push({ path: '/auth' })
}

</script>

<template>
    <Container>
        <div class="mb-3">
            <header class="py-4 fw-500 fs-9">
                <div class="grid grid-cols-[1fr_auto]">
                    <div>
                        <router-link to="/" class="remove-decorations">Home</router-link>
                    </div>
                    <div class="flex gap-3">
                        <router-link to="/auth">Login</router-link>
                        <div>/</div>
                        <router-link to="/auth" @click="Logout()">Logout</router-link>
                    </div>
                </div>
            </header>
            <hr class="hr-1-2">
        </div>
    </Container>
</template>
