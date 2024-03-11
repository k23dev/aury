<!-- src/App.svelte -->
<script>


  import { Input, Label, Helper } from 'flowbite-svelte';

  let name = '';
  let greeting = '';

  const handleInput = async () => {
    try {
      greeting = await Greet(name);
      console.log(greeting);
    } catch (error) {
      console.error('Error:', error);
    }
  };
</script>

<main>
  <form>
    <div class="grid gap-6 mb-6 md:grid-cols-2">
      <Label for="first_name" class="mb-2">First name</Label>
      <Input type="text" id="first_name" placeholder="John" bind:value={name} on:input={handleInput} required />
    </div>
  </form>
  <!-- <label for="name">Name:</label>
  <input type="text" id="name" bind:value={name} on:input={handleInput} /> -->

  {#if greeting}
    <p>{greeting}</p>
  {/if}
</main>
