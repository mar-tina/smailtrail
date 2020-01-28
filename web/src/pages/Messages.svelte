<script>
  import { onMount } from "svelte";
  import { push, pop, replace } from "svelte-spa-router";

  function setSender(from) {
    sender.set(from);
  }

  let nextpagetoken = "";
  let from = "";

  let msgs = [];

  onMount(async () => {
    const res = await fetch(
      `http://localhost:8000/allmessages?nextpagetoken=${nextpagetoken}`
    );
    let result = await res.json();
    msgs = result.msgs;

    console.log(result.msgs);
  });

  function handleClick() {
    let rout = "/thread/:";
    let res = rout.concat(sender);
    console.log("FROM", sender);
    console.log("THE ROUTE", res);
    push(rout);
  }
</script>

<style>
  .from {
    padding: 20px;
  }

  .subject {
    padding: 20px;
  }

  .container {
    border: 3px solid lightcoral;
    margin: 25px;
    padding: 20px;
    border-radius: 2%;
  }

  .container:hover {
    cursor: pointer;
  }
</style>

<div>

  {#each msgs as msg, i}
    <div class="container">
      {#each msg.Headers as hr, i}
        <div
          on:click={() => {
            if (hr.name == 'From') {
              from = hr.value;
            }
            console.log('WTF', from);
            push('/thread/:' + hr.value);
          }}>
          {#if hr.name == 'From'}
            <div class="from">{(from = hr.value)}</div>
          {/if}

          {#if hr.name == 'Subject'}
            <div class="subject">SUBJECT : {hr.value}</div>
          {/if}
        </div>
      {/each}
    </div>
  {/each}

</div>
