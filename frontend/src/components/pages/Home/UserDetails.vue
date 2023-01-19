<script setup lang="ts">
import { ref } from 'vue'

defineProps<{
    API_user: any
}>()

function secondsToMinutes(s: number) {
    let min = s / 60
    return min.toFixed(2)
}

</script>

<template>

    <div class="py-14 px-14 bg-black text-white mb-4 fw-700">
        <h3 class="fs-3 fw-800">{{ API_user.username }} </h3>
        <h6 class="f">Balance :: {{ API_user.balance }}</h6>
        <hr class="my-3">
        <div class="fw-400">
            <div>Created on :: {{ new Date(API_user.created_at).toDateString() }}</div>
            <div class="flex gap-2">
                <div>Minutes untill Next Bonus :: </div>
                <div class="fw-700">{{ secondsToMinutes(API_user.time_untill_bonus) }} min</div>
            </div>
            <div class="flex gap-2">
                ID for user :: <div class="font-mono opacity-60">{{ API_user.user_id }}</div>
            </div>
        </div>
    </div>


</template>
