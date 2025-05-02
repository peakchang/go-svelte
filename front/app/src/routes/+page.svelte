<script lang="ts">
    import type { UserInfo } from "$lib/interface/testdb";

    let str = $state("gogogogo");

    let user: { name: string; age: number } = $state({
        name: "민수",
        age: 28,
    });

    console.log(str);
    console.log(user);

    const userGo: UserInfo = {
        name: "피크 애드",
        age: 98,
    };

    console.log(userGo);
</script>

<h1>Welcome to SvelteKit</h1>
<p>
    Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the
    documentation
</p>
