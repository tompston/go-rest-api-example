<script setup lang="ts">

defineProps<{
    values: DebugRowInterface[]
}>();

interface DebugRowInterface {
    label: any
    value: any
}

</script>

<template>
    <div class="absolute bottom-0 right-0 m-7 bg-slate-900 text-white p-4 max-w-[380px] w-full border-rad-3">
        <div class="fs-9 fw-700 opacity-80">DEBUG GRID</div>
        <hr class="opacity-40 my-2" />
        <div class="grid gap-2 text-[12px] font-mono">
            <div v-for="value in values" v-bind:key="value.label" class="grid grid-cols-2 gap-3">
                <div class="overflow-auto max-h-[30px]">
                    {{ value.label }}
                </div>
                <div class="overflow-auto max-h-[30px]">
                    {{ value.value }}
                </div>
            </div>
        </div>
    </div>
</template>