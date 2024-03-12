<script lang="ts">

  import {handleHashChange} from './lib/aury/router';
  import { Input, Label, Helper } from 'flowbite-svelte';

  window.addEventListener('DOMContentLoaded', handleHashChange);
  window.addEventListener('hashchange', handleHashChange);
  let route=handleHashChange();
  console.log(`Loading... ${route}`);
  
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

  {#if greeting}
    <p>{greeting}</p>
  {/if}
</main>
