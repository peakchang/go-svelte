<script lang="ts">
    import type { UserInfo } from "$lib/interface/testdb";
    import axios, { AxiosError } from "axios";
    import { back_api } from "$lib/const";
    import { axiosPost } from "$lib/lib";

    let str = $state("gogogogo");

    let name: string = $state("");
    let age: number = $state(0);
    let idx: number = $state(0);

    async function insertDb() {
        const user: { name: string; age: number } = {
            name: name,
            age: Number(age),
        };
        const res = await axiosPost(`${back_api}/test/upload_test`, user);
        console.log(res);
    }

    async function updateDb() {
        const user: { name: string; age: number } = {
            name: name,
            age: Number(age),
        };
        const res = await axiosPost(`${back_api}/test/update_test`, user);
        console.log(res);
    }
</script>

<div class="px-3 max-w-7xl">
    <h1>Welcome to SvelteKit</h1>
    <p>
        Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read
        the documentation
    </p>

    <div class="mb-5">
        <input
            type="text"
            class="py-1 px-2 border border-gray-400 rounded-sm focus:outline-none"
            bind:value={name}
        />
        <input
            type="text"
            class="py-1 px-2 border border-gray-400 rounded-sm focus:outline-none"
            bind:value={age}
        />
    </div>

    <!-- svelte-ignore event_directive_deprecated -->
    <button
        class="bg-blue-400 py-1 px-3 rounded-md text-white"
        on:click={insertDb}
    >
        INSERT!!
    </button>

    <!-- svelte-ignore event_directive_deprecated -->
    <button
        class="bg-blue-400 py-1 px-3 rounded-md text-white"
        on:click={updateDb}
    >
        UPDATE!!
    </button>
</div>
