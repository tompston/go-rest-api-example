<script setup lang="ts">
import Container from '../components/global/Container.vue';
import Header from "../components/global/Header.vue"
</script>

<template>
    <Header />
    <Container>
            <div class="block mb-[70px]">
                <slot />
            </div>
        </Container>
</template>
